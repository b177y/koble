// defines driver interface for netkit compatible drivers
package driver

import (
	"fmt"

	"github.com/cri-o/ocicni/pkg/ocicni"
)

type Machine struct {
	Name     string
	Hostlab  string
	Hosthome string
	Networks []string
	Image    string
}

type MachineInfo struct {
	Name     string
	Lab      string
	Networks []string
	Image    string
	State    string
	Uptime   string
	ExitCode int32
	Exited   bool
	ExitedAt int64
	Mounts   []string
	HostPid  int
	Ports    []ocicni.PortMapping
}

type Network struct {
	Name     string
	External bool
	Gateway  string
	IpRange  string
	Subnet   string
	IPv6     string
}

type Driver interface {
	SetupDriver() (err error)

	StartMachine(m Machine, lab string) (id string, err error)
	StopMachine(name string) (err error)
	CrashMachine(name string) (err error)

	ListMachines(lab string) ([]MachineInfo, error)
	MachineExists(name string, lab string) (exists bool, err error)
	GetMachineState(name string, lab string) (state struct{}, err error)
	AttachToMachine(name string, lab string) (err error)
	MachineExecShell(name, lab, command, user string,
		detach bool, workdir string) (err error)
	GetMachineLogs(name, lab string, stdoutChan, stderrChan chan string,
		follow bool, tail int) (err error)

	ListNetworks(lab string) error
	NetworkExists(name, lab string) (exists bool, err error)
	CreateNetwork(net Network) (id string, err error)
	StartNetwork(name string, lab string) (err error)
	StopNetwork(name string, lab string) (err error)
	RemoveNetwork(name string, lab string) (err error)
	GetNetworkState(name string, lab string) (state string, err error)
}

type DriverError struct {
	Function string
	Driver   string
	Err      error
}

func (e *DriverError) Error() string {
	return fmt.Sprintf("Driver Error [%s] In %s: %v", e.Driver, e.Function, e.Err)
}

func NewDriverError(err error, driver, function string) *DriverError {
	return &DriverError{
		Function: function,
		Driver:   driver,
		Err:      err,
	}
}
