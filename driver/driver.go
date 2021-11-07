// defines driver interface for netkit compatible drivers
package driver

import (
	"errors"
	"fmt"

	"github.com/cri-o/ocicni/pkg/ocicni"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

var (
	ErrExists = errors.New("Already exists")
)

type Machine struct {
	Name     string
	Lab      string
	Hostlab  string
	Hosthome bool `default:"false"`
	Networks []string
	Volumes  []spec.Mount
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
	Lab      string
	External bool
	Gateway  string
	IpRange  string
	Subnet   string
	IPv6     string
}

type Driver interface {
	GetDefaultImage() string
	SetupDriver(conf map[string]interface{}) (err error)

	StartMachine(m Machine, lab string) (id string, err error)
	HaltMachine(m Machine, force bool) (err error)
	RemoveMachine(m Machine) (err error)

	ListMachines(lab string, all bool) ([]MachineInfo, error)
	MachineExists(name string, lab string) (exists bool, err error)
	GetMachineState(name string, lab string) (state string, err error)
	AttachToMachine(name string, lab string) (err error)
	MachineExecShell(name, lab, command, user string,
		detach bool, workdir string) (err error)
	GetMachineLogs(name, lab string, stdoutChan, stderrChan chan string,
		follow bool, tail int) (err error)

	ListNetworks(lab string, all bool) error
	NetworkExists(name, lab string) (exists bool, err error)
	CreateNetwork(net Network, lab string) (id string, err error)
	StartNetwork(name, lab string) (err error)
	StopNetwork(name, lab string) (err error)
	RemoveNetwork(name, lab string) (err error)
	GetNetworkState(name, lab string) (state string, err error)
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
