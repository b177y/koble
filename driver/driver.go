// defines driver interface for netkit compatible drivers
package driver

import (
	"errors"
	"time"

	"github.com/cri-o/ocicni/pkg/ocicni"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

var (
	ErrExists         = errors.New("Already exists")
	ErrNotExists      = errors.New("Doesn't exist")
	ErrNotImplemented = errors.New("Not implemented by current driver")
)

type Machine struct {
	Name         string `yaml:"name" validate:"alphanum,max=30"`
	Lab          string
	Namespace    string
	Hostlab      string
	Dependencies []string               `yaml:"depends_on,omitempty"`
	HostHome     bool                   `yaml:"hosthome,omitempty"`
	Networks     []string               `yaml:"networks,omitempty" validate:"alphanum,max=30"`
	Volumes      []spec.Mount           `yaml:"volumes,omitempty"`
	Image        string                 `yaml:"image,omitempty"`
	DriverExtra  map[string]interface{} `yaml:"driver_extra,omitempty"`
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
	return "netkit_" + m.Name + "_" + m.Namespace
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
	Name      string
	Lab       string
	Namespace string
	External  bool
	Gateway   string
	IpRange   string
	Subnet    string
	IPv6      string
}

func (n *Network) Fullname() string {
	return "netkit_" + n.Name + "_" + n.Namespace
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
	MachineInfo(m Machine) (info MachineInfo, err error)
	AttachToMachine(m Machine) (err error)
	MachineExecShell(m Machine, command, user string,
		detach bool, workdir string) (err error)
	GetMachineLogs(m Machine,
		follow bool, tail int) (err error)

	ListNetworks(lab string, all bool) (networks []NetInfo, err error)
	NetInfo(n Network) (info NetInfo, err error)
	NetworkExists(net Network) (exists bool, err error)
	CreateNetwork(net Network) (err error)
	StartNetwork(net Network) (err error)
	StopNetwork(net Network) (err error)
	RemoveNetwork(net Network) (err error)
	GetNetworkState(net Network) (state NetworkState, err error)
}
