package netkit

import (
	"fmt"
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
		Lab:  nk.Lab.Name,
	}
	err := nk.Driver.CreateNetwork(n)
	if err != nil {
		return err
	}
	err = nk.Driver.StartNetwork(n)
	return err
}

func (nk *Netkit) ListNetworks(all bool) error {
	if !all {
		if nk.Lab.Name == "" {
			fmt.Println("Listing all networks which are not associated with a lab.")
			fmt.Printf("To see all machines use `netkit net list --all`\n\n")
		} else {
			fmt.Printf("Listing all networks within this lab (%s).\n", nk.Lab.Name)
			fmt.Printf("To see all machines use `netkit net list --all`\n\n")
		}
		networks, err := nk.Driver.ListNetworks(nk.Lab.Name, all)
		if err != nil {
			return err
		}
		nlist, headers := NetInfoToStringArr(networks, all)
		RenderTable(headers, nlist)
	}
	return nil
}
