package podman

import (
	"context"
	"fmt"

	"github.com/b177y/koble/pkg/driver"
	"github.com/go-playground/validator/v10"
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
	if err := validator.New().Var(name, "alphanum,max=30"); err != nil {
		return m, fmt.Errorf("machine name '%s' must be alphanumeric and no more than 30 chars", name)
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
	if err := validator.New().Var(name, "alphanum,max=30"); err != nil {
		return n, fmt.Errorf("network name '%s' must be alphanumeric and no more than 30 chars", name)
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
