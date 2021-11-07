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

type NetworkState struct {
	Running bool
}

type NetInfo struct {
	Name      string
	Lab       string
	Interface string
	External  bool
	Gateway   string
	Subnet    string
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

func (n *Network) Fullname() string {
	name := "netkit_" + n.Name
	if n.Lab != "" {
		name += "_" + n.Lab
	}
	return name
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

	ListNetworks(lab string, all bool) (networks []NetInfo, err error)
	NetworkExists(net Network) (exists bool, err error)
	CreateNetwork(net Network) (err error)
	StartNetwork(net Network) (err error)
	StopNetwork(net Network) (err error)
	RemoveNetwork(net Network) (err error)
	GetNetworkState(net Network) (state NetworkState, err error)
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
