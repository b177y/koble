package podman

import (
	"context"
	"fmt"

	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/util/validator"
	"github.com/spf13/cobra"
)

type PodmanDriver struct {
	Conn       context.Context
	Name       string // friendly name
	DriverName string
	Config     Config
}

func (pd *PodmanDriver) Machine(name, namespace string) (m driver.Machine,
	err error) {
	if !validator.IsValidName(name) {
		return m, fmt.Errorf("machine name '%s' must be alphanumeric and no more than 32 chars", name)
	}
	m = &Machine{
		name:      name,
		namespace: namespace,
		pd:        pd,
	}
	return m, nil
}

func (pd *PodmanDriver) Network(name, namespace string) (n driver.Network,
	err error) {
	if !validator.IsValidName(name) {
		return n, fmt.Errorf("network name '%s' must be alphanumeric and no more than 32 chars", name)
	}
	n = &Network{
		name:      name,
		Namespace: namespace,
		pd:        pd,
	}
	return n, nil
}

func (pd *PodmanDriver) GetCLICommand() (command *cobra.Command, err error) {
	return new(cobra.Command), nil
}
