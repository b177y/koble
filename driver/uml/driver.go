package uml

import "github.com/b177y/koble/driver"

type UMLDriver struct {
	Name         string
	DefaultImage string
	Kernel       string
	RunDir       string
	StorageDir   string
	Testing      bool
}

func (ud *UMLDriver) Machine(name, namespace string) (m driver.Machine,
	err error) {
	m = &Machine{
		name:      name,
		namespace: namespace,
		ud:        ud,
	}
	return m, nil
}

func (ud *UMLDriver) Network(name, namespace string) (n driver.Network,
	err error) {
	n = &Network{
		name:      name,
		namespace: namespace,
		ud:        ud,
	}
	return n, nil
}
