package vecnet

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/containernetworking/plugins/pkg/ns"
)

var NSPID = "/run/user/1000/uml/ns-pid"

// Creates and persists a new user and netnamespace by running a 'pause' process
// Runs bash with infinite sleep to keep new ns open, as namespaces cannot be bound without root privs
// The pid is then saved to NSPID which can then be used with /proc/PID/ns/user and /proc/PID/ns/net to enter these namespaces
// For debugging: `nsenter --preserve-credentials -t PID --user --net`
func NewNS() error {
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
			Cloneflags: syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
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
// within the UML user namespace opened by NewNS()
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
	nsArgs := fmt.Sprintf("/usr/bin/nsenter --preserve-credentials -t %s --user %s", string(pid), "./nettest")
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

// Wrapper arround containernetworking's netns helper
// fills in the pid for the uml namespace
// the function toRun will be run within the uml net namespace
// opened by NewNS()
func WithNetNS(toRun func(ns.NetNS) error) error {
	pid, err := ioutil.ReadFile("/run/user/1000/uml/ns-pid")
	if err != nil {
		return err
	}
	nsPath := fmt.Sprintf("/proc/%s/ns/net", string(pid))
	return ns.WithNetNSPath(nsPath, toRun)
}
