package vecnet

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"

	"github.com/containernetworking/plugins/pkg/ns"
	"golang.org/x/sys/unix"
)

var NSPID = "/run/user/1000/uml/ns-pid"

// Creates and persists a new user and mount namespace by running a 'pause' process
// Runs bash with infinite sleep to keep new ns open, as namespaces cannot be bound without root privs
// The pid is then saved to NSPID which can then be used with /proc/PID/ns/user and /proc/PID/ns/mount to enter these namespaces
// For debugging: `nsenter --preserve-credentials -t PID --user --mount`
func NewUserNS() error {
	dir, _ := filepath.Split(NSPID)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	// Pause process
	// https://www.redhat.com/sysadmin/behind-scenes-podman
	cmd := exec.Cmd{
		Path: "/bin/bash",
		Args: []string{"-i", "-c", "sleep infinity"},
		SysProcAttr: &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER, // new mount and user ns
			UidMappings: []syscall.SysProcIDMap{{
				ContainerID: 0,
				HostID:      1000, // TODO use UID
				Size:        1,
			}}, // TODO add GID mapping ?
		},
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	// pidPath := filepath.Join("/run/user/1000/uml/ns-pid")
	return nil
}

// Reexec of current process with nsenter as intermediate
// before child process runs.
// Child process will be almost identical but will be running
// within the UML user + mount namespace opened by NewNS()
// Problems with golang's runtime mean that entering a user ns cannot be done
// natively with syscalls
// https://github.com/kata-containers/runtime/issues/148
// https://github.com/golang/go/issues/8676
// cgo could be used as an alternative to rexec but will affect portability
func EnterUserNS() error {
	pid, err := ioutil.ReadFile(NSPID)
	if err != nil {
		return err
	}
	cmdline, err := ioutil.ReadFile("/proc/self/cmdline")
	if err != nil {
		return err
	}
	fmt.Println("I was called with", string(cmdline))
	nsArgs := fmt.Sprintf("/usr/bin/nsenter --preserve-credentials -t %s --user --mount %s", string(pid), "/home/billy/repos/github.com/b177y/netkit/nettest")
	bashArgs := []string{"-i", "-c"}
	// nsArgs = append(nsArgs, fmt.Sprintf("/proc/%d/exe", os.Getpid()))
	bashArgs = append(bashArgs, nsArgs)
	cmd := exec.Cmd{
		Path: "/bin/bash",
		Args: bashArgs,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func NewNetNS(name string) error {
	// setup 'rundir'
	nsDir := "/run/user/1000/uml"
	err := os.MkdirAll(nsDir, 0755)
	if err != nil {
		return err
	}
	err = syscall.Mount("", nsDir, "none", syscall.MS_SHARED|syscall.MS_REC, "")
	if err == syscall.EINVAL {
		err = syscall.Mount(nsDir, nsDir, "none", syscall.MS_BIND|syscall.MS_REC, "")
		if err != nil {
			return fmt.Errorf("mount --rbind %s %s failed: %q", nsDir, nsDir, err)
		}
		err = syscall.Mount("", nsDir, "none", syscall.MS_SHARED|syscall.MS_REC, "")
	} else if err != nil {
		return fmt.Errorf("mount --make-rshared %s failed: %q", nsDir, err)
	}
	// TODO dont hardcode
	nsPath := filepath.Join(nsDir, "net-ns")
	f, err := os.OpenFile(nsPath, os.O_CREATE|os.O_EXCL, 0444)
	f.Close()
	// Ensure the mount point is cleaned up on errors; if the namespace
	// was successfully mounted this will have no effect because the file
	// is in-use - from containernetworking/plugins/pkg/testutils
	defer os.RemoveAll(nsPath)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		runtime.LockOSThread()
		// Don't unlock. By not unlocking, golang will kill the OS thread when the
		// goroutine is done (for go1.10+)
		// https://github.com/containernetworking/plugins/blob/2c46a726805bcf13e2f78580c57b21e9de107285/pkg/testutils/netns_linux.go
		defer wg.Done()
		origNS, err := ns.GetNS(fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid()))
		if err != nil {
			return
		}
		defer origNS.Close()

		err = syscall.Unshare(syscall.CLONE_NEWNET)
		if err != nil {
			return
		}

		defer origNS.Set()
		fmt.Println("mounting to", nsPath)
		err = syscall.Mount(fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid()), nsPath, "bind", syscall.MS_BIND, "")
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
	return err
}

// Wrapper arround containernetworking's netns helper
// fills in the pid for the uml namespace
// the function toRun will be run within the uml net namespace
// opened by NewNS()
func WithNetNS(toRun func(ns.NetNS) error) error {
	// TODO dont hardcode
	nsPath := "/run/user/1000/uml/net-ns"
	return ns.WithNetNSPath(nsPath, toRun)
}
