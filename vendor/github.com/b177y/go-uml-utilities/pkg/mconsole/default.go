package mconsole

import (
	"fmt"

	"github.com/b177y/go-uml-utilities/pkg/mconsole/sysrq"
)

// Gets the UML version.
// Equivalent of `uname -srnmv` within uml instance.
func UMLKernelVersion() (command string) {
	return "version"
}

// Shut the machine down immediately
// with no syncing of disks and no clean shutdown of userspace.
func Halt() (command string) {
	return "halt"
}

// Reboot the machine immediately
// with no syncing of disks and no clean shutdown of userspace.
func Reboot() (command string) {
	return "reboot"
}

// Calls the generic kernel’s SysRq driver, which does whatever is called for by that argument.
// See https://www.kernel.org/doc/html/latest/admin-guide/sysrq.html for more info.
func SysRQCommand(cmd sysrq.SRQCommand) string {
	return fmt.Sprintf("sysrq %c", cmd)
}

// Invokes the Ctl-Alt-Del action in the running image.
// What exactly this ends up doing is up to init, systemd, etc.
// Normally, it reboots the machine.
func CtrlAltDel() (command string) {
	return "cad"

}

// Puts the UML in a loop reading mconsole requests until a ‘go’ mconsole command is received.
func Stop() (command string) {
	return "stop"
}

// Resumes a UML after being paused by a ‘stop’ command.
// Note that when the UML has resumed, TCP connections may have timed out
// and if the UML is paused for a long period of time,
// crond might go a little crazy, running all the jobs it didn’t do earlier.
func Go() (command string) {
	return "go"
}

// Gets the contents of a specified file from the `/proc` directory
func Proc(procfile string) (command string) {
	return fmt.Sprintf("proc %s", procfile)
}

// Gets the stack for a specified process
func Stack(pid int) (command string) {
	return fmt.Sprintf("stack %d", pid)
}

// Logs message to UML kernel log
func Log(message string) (command string) {
	return fmt.Sprintf("log %s", message)
}
