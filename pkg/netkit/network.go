package netkit

import (
	"net"

	"github.com/b177y/netkit/driver"
)

type Network struct {
	Name     string `yaml:"name" validate:"alphanum,max=30"`
	External bool   `yaml:"external,omitempty"`
	Gateway  net.IP `yaml:"gateway,omitempty" validate:"ip"`
	Subnet   string `yaml:"subnet,omitempty" validate:"cidr"`
	IPv6     bool   `yaml:"ipv6,omitempty" validate:"ipv6"`
}

func (nk *Netkit) StartNetwork(name string) error {
	n := driver.Network{
		Name: name,
	}
	err := nk.Driver.CreateNetwork(n, nk.Lab.Name)
	if err != nil {
		return err
	}
	err = nk.Driver.StartNetwork(name, nk.Lab.Name)
	return err
}
