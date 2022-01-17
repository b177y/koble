package mock

import (
	"github.com/b177y/koble/driver"
)

func (md *MockDriver) CreateNetwork(n driver.Network) (err error) {
	return err
}

func (md *MockDriver) StartNetwork(net driver.Network) (err error) {
	return nil
}

func (md *MockDriver) RemoveNetwork(net driver.Network) (err error) {
	return err
}

func (md *MockDriver) StopNetwork(net driver.Network) (err error) {
	return nil
}

func (md *MockDriver) GetNetworkState(net driver.Network) (state driver.NetworkState,
	err error) {
	return state, err
}

func (md *MockDriver) ListNetworks(lab string, all bool) (networks []driver.NetInfo, err error) {
	return networks, nil
}

func (md *MockDriver) NetworkExists(net driver.Network) (bool, error) {
	return false, nil
}

func (md *MockDriver) NetInfo(net driver.Network) (nInfo driver.NetInfo, err error) {
	return nInfo, err
}
