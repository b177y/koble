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

type Driver interface {
	GetDefaultImage() string
	SetupDriver(conf map[string]interface{}) (err error)

	Machine(name, namespace string) (Machine, error)
	ListMachines(namespace string, all bool) ([]MachineInfo, error)

	Network(name, namespace string) (Network, error)
	ListNetworks(namespace string, all bool) (networks []NetInfo, err error)

	ListAllNamespaces() ([]string, error)
	GetCLICommand() (command *cobra.Command, err error)
}
