package vecnet

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

const (
	CLONE_NEWNS   = 0x00020000
	CLONE_NEWUTS  = 0x04000000
	CLONE_NEWIPC  = 0x08000000
	CLONE_NEWNET  = 0x40000000
	CLONE_NEWUSER = 0x10000000
	CLONE_NEWPID  = 0x20000000
	SYS_SETNS     = 308
)

func NewNS(namespace string) error {
	if _, err := os.Stat("/run/user/1000/uml"); os.IsNotExist(err) {
		err = os.MkdirAll("/run/user/1000/uml", 0755)
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
	// pidPath := filepath.Join("/run/user/1000/uml/netkit-ns")
	return nil
}

func JoinNS(network string) error {
	runtime.LockOSThread()
	pid, err := ioutil.ReadFile("/run/user/1000/uml/netkit-ns")
	if err != nil {
		return err
	}
	cmdline, err := ioutil.ReadFile("/proc/self/cmdline")
	if err != nil {
		return err
	}
	fmt.Println("I was called with", string(cmdline))
	nsArgs := fmt.Sprintf("/usr/bin/nsenter --preserve-credentials -t %s --user --net %s", string(pid), "./nettest")
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
	fmt.Println("Running bash with", nsArgs)
	return cmd.Run()
}

func NewBridge(network, namespace string) error {
	// switch ns
	// add bridge
	// set bridge up
	// TODO set stp on or off, default off
	return nil
}

func AddHost(iface, network, namespace string) error {
	// switch ns
	// ip tuntap add tap X
	// ip link set up
	// brctl addif to bridge
	return nil
}

func AddTapout(iface, network, namespace string) error {
	// check ns exists
	// exec slirp4netns with params
	// create bridge
	// add host to bridge
	// add slirp's iface to bridge
	return nil
}
