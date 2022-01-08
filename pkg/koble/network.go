package koble

import (
	"net"
	"strconv"

	"github.com/b177y/koble/driver"
)

type Network struct {
	Name     string `yaml:"name" validate:"alphanum,max=30"`
	External bool   `yaml:"external,omitempty"`
	Gateway  net.IP `yaml:"gateway,omitempty" validate:"ip"`
	Subnet   string `yaml:"subnet,omitempty" validate:"cidr"`
	IPv6     bool   `yaml:"ipv6,omitempty" validate:"ipv6"`
}

func (nk *Koble) StartNetwork(network Network) error {
	n, err := nk.Driver.Network(network.Name, nk.Namespace)
	if err != nil {
		return err
	}
	err = n.Create(nil) // TODO add options
	if err != nil {
		return err
	}
	err = n.Start()
	return err
}

func (nk *Koble) ListNetworks(all bool) error {
	networks, err := nk.Driver.ListNetworks(nk.Lab.Name, all)
	if err != nil {
		return err
	}
	nlist, headers := NetInfoToStringArr(networks, all)
	RenderTable(headers, nlist)
	return nil
}

func (nk *Koble) NetworkInfo(name string) error {
	n, err := nk.Driver.Network(name, nk.Namespace)
	if err != nil {
		return err
	}
	var infoTable [][]string
	infoTable = append(infoTable, []string{"Name", n.Name()})
	// get machines connected
	info, err := n.Info()
	if err != nil && err != driver.ErrNotExists {
		return err
	}
	if err != driver.ErrNotExists {
		// infoTable = append(infoTable, []string{"Interface", info.Interface})
		infoTable = append(infoTable, []string{"External", strconv.FormatBool(info.External)})
		infoTable = append(infoTable, []string{"Gateway", info.Gateway})
		infoTable = append(infoTable, []string{"Subnet", info.Subnet})
	}
	RenderTable([]string{}, infoTable)
	return nil
}
