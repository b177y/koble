package mconsole

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

// find pid file created by UML by looking for a file named 'pid'
// in the same directory as sockpath,
// and send SIGINT to the process
func InterruptUML(sockpath string) error {
	pidFile := filepath.Join(filepath.Dir(sockpath), "pid")
	pidBytes, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("error reading pidfile %s: %w", pidFile, err)
	}
	pid, err := strconv.Atoi(strings.Trim(string(pidBytes), "\n"))
	if err != nil {
		return fmt.Errorf("could not read pid file %s as integer: %w",
			pidFile, err)
	}
	// TODO possibly add choice of signal?
	// currently just doing SIGINT as with original code:
	// https://salsa.debian.org/uml-team/uml-utilities/-/blob/master/mconsole/uml_mconsole.c#L452
	return syscall.Kill(pid, syscall.SIGINT)
}
