package koble

import (
	"strconv"

	"github.com/b177y/koble/driver"
)

func (n *Network) Start() error {
	dn, err := n.nk.Driver.Network(n.Name, n.nk.Namespace)
	if err != nil {
		return err
	}
	err = dn.Create(nil) // TODO add options
	if err != nil {
		return err
	}
	err = dn.Start()
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

func (n *Network) Info() error {
	dn, err := n.nk.Driver.Network(n.Name, n.nk.Namespace)
	if err != nil {
		return err
	}
	var infoTable [][]string
	infoTable = append(infoTable, []string{"Name", n.Name})
	// get machines connected
	info, err := dn.Info()
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
