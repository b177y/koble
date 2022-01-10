package podman

import (
	"context"

	"github.com/b177y/koble/driver"
	"github.com/spf13/cobra"
)

type PodmanDriver struct {
	conn   context.Context
	Name   string
	Config Config
}

func (pd *PodmanDriver) Machine(name, namespace string) (m driver.Machine,
	err error) {
	m = &Machine{
		name:      name,
		namespace: namespace,
		pd:        pd,
	}
	return m, nil
}

func (pd *PodmanDriver) Network(name, namespace string) (n driver.Network,
	err error) {
	n = &Network{
		name:      name,
		namespace: namespace,
		pd:        pd,
	}
	return n, nil
}

func (pd *PodmanDriver) GetCLICommand() (command *cobra.Command, err error) {
	return new(cobra.Command), nil
}
