package uml

import (
	"fmt"

	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/driver/podman"
	"github.com/b177y/koble/util/validator"
	"github.com/spf13/cobra"
)

type UMLDriver struct {
	Name   string
	Config Config
	Podman podman.PodmanDriver
}

func (ud *UMLDriver) Machine(name, namespace string) (m driver.Machine,
	err error) {
	if !validator.IsValidName(name) {
		return m, fmt.Errorf("machine name '%s' must be alphanumeric and no more than 32 chars", name)
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
	if !validator.IsValidName(name) {
		return n, fmt.Errorf("network name '%s' must be alphanumeric and no more than 32 chars", name)
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
