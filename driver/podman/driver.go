package podman

import (
	"context"

	"github.com/b177y/koble/driver"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

type PodmanDriver struct {
	conn   context.Context
	Name   string
	Config Config
}

func (pd *PodmanDriver) Machine(name, namespace string) (m driver.Machine,
	err error) {
	if err := validator.New().Var(name, "alphanum,max=30"); err != nil {
		return m, err
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
		return n, err
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
