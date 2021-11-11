package uml

import (
	"github.com/b177y/netkit/driver"
)

func (ud *UMLDriver) CreateNetwork(n driver.Network) (err error) {
	exists, err := ud.NetworkExists(n)
	if err != nil {
		return err
	}
	if exists {
		return driver.ErrExists
	}
	return err
}

func (ud *UMLDriver) StartNetwork(net driver.Network) (err error) {
	// podman network doesn't need manual starting
	return nil
}

func (ud *UMLDriver) RemoveNetwork(net driver.Network) (err error) {
	// _, err = network.Remove(ud.conn, net.Fullname(), nil)
	return err
}

func (ud *UMLDriver) StopNetwork(net driver.Network) (err error) {
	// podman network doesn't need manual stopping
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

func (ud *UMLDriver) NetworkExists(net driver.Network) (bool, error) {
	// return network.Exists(ud.conn, net.Fullname(), nil)
	return false, nil
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
