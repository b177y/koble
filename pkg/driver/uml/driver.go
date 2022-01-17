package uml

import (
	"github.com/b177y/koble/pkg/driver"
	"github.com/go-playground/validator/v10"
)

type UMLDriver struct {
	Name   string
	Config Config
}

func (ud *UMLDriver) Machine(name, namespace string) (m driver.Machine,
	err error) {
	if err := validator.New().Var(name, "alphanum,max=30"); err != nil {
		return m, err
	}
	m = &Machine{
		name:      name,
		namespace: namespace,
		ud:        ud,
	}
	return m, nil
}

func (ud *UMLDriver) Network(name, namespace string) (n driver.Network,
	err error) {
	if err := validator.New().Var(name, "alphanum,max=30"); err != nil {
		return n, err
	}
	n = &Network{
		name:      name,
		namespace: namespace,
		ud:        ud,
	}
	return n, nil
}
