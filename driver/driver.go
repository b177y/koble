// defines driver interface for netkit compatible drivers
package driver

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	ErrExists         = errors.New("Already exists")
	ErrNotExists      = errors.New("Doesn't exist")
	ErrNotImplemented = errors.New("Not implemented by current driver")
)

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

	Machine() (Machine, error)
	ListMachines(namespace string, all bool) ([]MachineInfo, error)

	ListNetworks(lab string, all bool) (networks []NetInfo, err error)
	NetInfo(n Network) (info NetInfo, err error)
	NetworkExists(net Network) (exists bool, err error)
	CreateNetwork(net Network) (err error)
	StartNetwork(net Network) (err error)
	StopNetwork(net Network) (err error)
	RemoveNetwork(net Network) (err error)
	GetNetworkState(net Network) (state NetworkState, err error)

	ListAllNamespaces() ([]string, error)
	GetCLICommand() (command *cobra.Command, err error)
}
