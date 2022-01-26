package upod

import (
	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/driver/podman"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

type UMLDriver struct {
	Name   string
	Config Config
	Podman podman.PodmanDriver
}

func (ud *UMLDriver) Machine(name, namespace string) (m driver.Machine,
	err error) {
	if err := validator.New().Var(name, "alphanum,max=30"); err != nil {
		return m, err
	}
	pm, err := ud.Podman.Machine(name, namespace)
	if err != nil {
		return m, err
	}
	m = &Machine{
		name:      name,
		namespace: namespace,
		ud:        ud,
		p:         pm,
	}
	return m, nil
}

func (ud *UMLDriver) Network(name, namespace string) (n driver.Network,
	err error) {
	if err := validator.New().Var(name, "alphanum,max=30"); err != nil {
		return n, err
	}
	pn, err := ud.Podman.Network(name, namespace)
	if err != nil {
		return n, err
	}
	n = &Network{
		name:      name,
		namespace: namespace,
		ud:        ud,
		p:         pn,
	}
	return n, nil
}

func (ud *UMLDriver) GetCLICommand() (command *cobra.Command, err error) {
	return new(cobra.Command), nil
}
