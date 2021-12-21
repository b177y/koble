package uml

import (
	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/uml/vecnet"
)

func (ud *UMLDriver) CreateNetwork(n driver.Network) (err error) {
	exists, err := ud.NetworkExists(n)
	if err != nil {
		return err
	}
	if exists {
		return driver.ErrExists
	}
	err = vecnet.NewBridge("br_"+n.Name, n.Namespace)
	if err != nil {
		return err
	}
	if n.External {
		return vecnet.SetupExternal("tap0", "br_"+n.Name, n.Namespace, "", "")
	}
	return nil
}

func (ud *UMLDriver) StartNetwork(net driver.Network) (err error) {
	return nil
}

func (ud *UMLDriver) RemoveNetwork(net driver.Network) (err error) {
	return err
}

func (ud *UMLDriver) StopNetwork(net driver.Network) (err error) {
	return nil
}

func (ud *UMLDriver) GetNetworkState(net driver.Network) (state driver.NetworkState,
	err error) {
	state.Running, err = ud.NetworkExists(net)
	return state, err
}

func (ud *UMLDriver) ListNetworks(lab string, all bool) (networks []driver.NetInfo, err error) {
	return networks, nil
}

func (ud *UMLDriver) NetworkExists(n driver.Network) (bool, error) {
	return vecnet.IfaceExistsWithNS(n.Namespace, "br_"+n.Name)
}

func (ud *UMLDriver) NetInfo(net driver.Network) (nInfo driver.NetInfo, err error) {
	exists, err := ud.NetworkExists(net)
	if err != nil {
		return nInfo, err
	}
	if !exists {
		return nInfo, driver.ErrNotExists
	}
	return nInfo, err
}
