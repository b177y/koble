package driver

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ErrExists         = errors.New("Already exists")
	ErrNotExists      = errors.New("Doesn't exist")
	ErrNotImplemented = errors.New("Not implemented by current driver")
)

type DriverInitialiser func() Driver

var registeredDrivers = map[string]DriverInitialiser{}

func RegisterDriver(name string, d DriverInitialiser) {
	if _, ok := registeredDrivers[name]; ok {
		log.Warnf("not registering driver %s, already registered.\n", name)
	}
	registeredDrivers[name] = d
}

func RegisterDriverCmds(cmd *cobra.Command) error {
	for _, d := range registeredDrivers {
		if dCmd, err := d().GetCLICommand(); err == ErrNotImplemented {
			continue
		} else if err != nil {
			return err
		} else {
			cmd.AddCommand(dCmd)
		}
	}
	return nil
}

func GetDriver(name string, conf map[string]interface{}) (d Driver,
	err error) {
	if initialiser, ok := registeredDrivers[name]; ok {
		d := initialiser()
		return d, d.SetupDriver(conf)
	} else {
		return d, fmt.Errorf("Driver %s is not currently supported.", name)
	}
}

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
