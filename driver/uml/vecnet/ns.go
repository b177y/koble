package vecnet

/*
extern void switchNamespace(void);
extern char *nsMsg;
void __attribute__((constructor)) init(void) {
	switchNamespace();
}
*/
import "C"
import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var NSPID_DIR = "/run/user/1000/uml"

func NSStatus() error {
	if C.nsMsg == nil {
		return nil
	}
	return errors.New(C.GoString(C.nsMsg))
}

func UserNSExists(name string) (exists bool, err error) {
	// TODO
	return false, err
}

// Creates and persists a new user and mount namespace by running a 'pause' process
// Runs bash with infinite sleep to keep new ns open, as namespaces cannot be bound without root privs
// The pid is then saved to NSPID which can then be used with /proc/PID/ns/user and /proc/PID/ns/mount to enter these namespaces
// For debugging: `nsenter --preserve-credentials -t PID --user --mount`
func CreateUserNS(name string) error {
	if _, err := os.Stat(NSPID_DIR); os.IsNotExist(err) {
		err = os.MkdirAll(NSPID_DIR, 0755)
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
	// save pid of pause process to file
	return ioutil.WriteFile(filepath.Join(NSPID_DIR, name+"-ns.pid"),
		[]byte(fmt.Sprint(cmd.Process.Pid)), 0644)
}

// Reexec of current process with UML_NS_PID environment variable set
// for cgo code to pick up and enter user and mount namespace before
// go runtime begins executing code
// REVISIT: in future this may be fixed in the golang runtime
// https://groups.google.com/g/golang-dev/c/Q5NVkQbg6bs
// https://github.com/golang/go/issues/8676 [status wontfix :(]
func ExecUserNS(name string) error {
	pid, err := ioutil.ReadFile(filepath.Join(NSPID_DIR, name+"-ns.pid"))
	if err != nil {
		return err
	}
	cmd := exec.Cmd{
		Path: "/proc/self/exe",
		Args: os.Args[1:],
		Env:  append(os.Environ(), "UML_NS_PID="+string(pid)),
	}
	return cmd.Run()
}

func CreateAndEnterUserNS(name string) error {
	pidEnv := os.Getenv("UML_NS_PID")
	if pidEnv == "" {
		exists, err := UserNSExists(name)
		if err != nil {
			return err
		}
		if !exists {
			err = CreateUserNS(name)
			if err != nil {
				return err
			}
		}
		return ExecUserNS(name)
	} else {
		return NSStatus()
	}
}
