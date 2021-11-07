// defines driver interface for netkit compatible drivers
package driver

import (
	"errors"
	"fmt"
	"time"

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

type MachineState struct {
	Pid       int
	Status    string
	Running   bool
	StartedAt time.Time
	ExitCode  int32
}

func (m *Machine) Fullname() string {
	name := "netkit_" + m.Name
	if m.Lab != "" {
		name += "_" + m.Lab
	}
	return name
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

	StartMachine(m Machine) (err error)
	HaltMachine(m Machine, force bool) (err error)
	RemoveMachine(m Machine) (err error)

	ListMachines(lab string, all bool) ([]MachineInfo, error)
	MachineExists(m Machine) (exists bool, err error)
	GetMachineState(m Machine) (state MachineState, err error)
	AttachToMachine(m Machine) (err error)
	MachineExecShell(m Machine, command, user string,
		detach bool, workdir string) (err error)
	GetMachineLogs(m Machine, stdoutChan, stderrChan chan string,
		follow bool, tail int) (err error)

	ListNetworks(lab string, all bool) error
	NetworkExists(name, lab string) (exists bool, err error)
	CreateNetwork(net Network, lab string) (err error)
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
