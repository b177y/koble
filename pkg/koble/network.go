package koble

import (
	"strconv"

	"github.com/b177y/koble/pkg/driver"
)

func (nk *Koble) StartNetwork(name string, conf driver.NetConfig) error {
	n, err := nk.Driver.Network(name, nk.Config.Namespace)
	if err != nil {
		return err
	}
	if err := n.Create(&conf); err != nil {
		return err
	}
	return n.Start()
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

func (nk *Koble) NetworkInfo(name string, json bool) error {
	n, err := nk.Driver.Network(name, nk.Config.Namespace)
	if err != nil {
		return err
	}
	var infoTable [][]string
	infoTable = append(infoTable, []string{"Name", name})
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
