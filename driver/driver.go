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
	// Setup the driver using config map
	SetupDriver(conf map[string]interface{}) (err error)

	// Get a Machine instance within a specified namespace
	Machine(name, namespace string) (Machine, error)
	// Get a list of all machines within a specified namespace,
	// or all machines if all is true.
	// If all is true, the namespace parameter will be ignored
	ListMachines(namespace string, all bool) ([]MachineInfo, error)

	// Get a Network instance within a specified namespace
	Network(name, namespace string) (Network, error)
	// Get a list of all networks within a specified namespace,
	// or all networks if all is true.
	// If all is true, the namespace parameter will be ignored
	ListNetworks(namespace string, all bool) (networks []NetInfo, err error)

	// Return a list of all namespaces
	ListAllNamespaces() ([]string, error)
	// Return a Cobra CLI command which can be added to an existing CLI
	// as a subcommand
	GetCLICommand() (command *cobra.Command, err error)
}
