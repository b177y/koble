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
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var NSPID_DIR = "/run/user/1000/uml"

func NSStatus() error {
	if C.nsMsg == nil {
		return nil
	}
	return errors.New(C.GoString(C.nsMsg))
}

func UserNSExists(name string) (exists bool, err error) {
	pidPath := filepath.Join(NSPID_DIR, name+"-ns.pid")
	pidBytes, err := ioutil.ReadFile(pidPath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("Error checking if pause pid file exists: %w", err)
	}
	pid, err := strconv.Atoi(string(pidBytes))
	if err != nil {
		return false, fmt.Errorf("uml ns pidfile doesn't contain integer: %s", string(pidBytes))
	}
	// if process is not running, remove reference and return ns does not exist
	err = syscall.Kill(pid, syscall.Signal(0))
	if err != nil {
		os.RemoveAll(pidPath)
		return false, nil
	}
	return true, nil
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
				HostID:      os.Getuid(),
				Size:        1,
			}},
			GidMappings: []syscall.SysProcIDMap{{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			}},
		},
	}
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("Could not run User NS pause process: %w", err)
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
	env := os.Environ()
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	env = append(env, "UML_NS_PID="+string(pid))
	env = append(env, "UML_ORIG_HOME="+home)
	env = append(env, "UML_ORIG_UID="+fmt.Sprint(os.Getuid()))
	env = append(env, "UML_ORIG_EUID="+fmt.Sprint(os.Geteuid()))
	env = append(env, "UML_ORIG_GID="+fmt.Sprint(os.Getgid()))
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	env = append(env, "UML_ORIG_WD="+wd)
	cmd := exec.Cmd{
		Path:   "/proc/self/exe",
		Args:   os.Args,
		Env:    env,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("Error running ns reexec: %w", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Debugf("Error from reexec child: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
	return nil
}

func CreateAndEnterUserNS(name string) error {
	pidEnv := os.Getenv("UML_NS_PID")
	if pidEnv == "" && os.Getuid() != 0 {
		fmt.Println("gonna enter ns")
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
		if os.Getenv("UML_ORIG_UID") == "" {
			return errors.New("environment variable UML_ORIG_UID has not been set")
		} else if os.Getenv("UML_ORIG_EUID") == "" {
			return errors.New("environment variable UML_ORIG_EUID has not been set")
		} else if os.Getenv("UML_ORIG_GID") == "" {
			return errors.New("environment variable UML_ORIG_GID has not been set")
		} else if os.Getenv("UML_ORIG_WD") == "" {
			return errors.New("environment variable UML_ORIG_WD has not been set")
		} else if os.Getenv("UML_ORIG_HOME") == "" {
			return errors.New("environment variable UML_ORIG_HOME has not been set")
		}
		if os.Getuid() != 0 {
			return NSStatus()
		} else {
			return nil
		}
	}
}
