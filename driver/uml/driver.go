package uml

import "github.com/b177y/netkit/driver"

type UMLDriver struct {
	Name         string
	DefaultImage string
	Kernel       string
	RunDir       string
	StorageDir   string
	Testing      bool
}

func (ud *UMLDriver) GetDefaultImage() string {
	return ud.DefaultImage
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
