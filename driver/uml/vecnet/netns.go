package vecnet

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"

	"github.com/containernetworking/plugins/pkg/ns"
	"golang.org/x/sys/unix"
)

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