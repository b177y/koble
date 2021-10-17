package podman

import "github.com/b177y/netkit/driver"

func (pd *PodmanDriver) CreateNetwork(net driver.Network) (id string,
	err error) {
	return id, err
}

func (pd *PodmanDriver) StartNetwork(name, lab string) (err error) {
	return err
}

func (pd *PodmanDriver) RemoveNetwork(name, lab string) (err error) {
	return err
}

func (pd *PodmanDriver) StopNetwork(name, lab string) (err error) {
	return err
}

func (pd *PodmanDriver) GetNetworkState(name, lab string) (state string,
	err error) {
	return "", err
}

func (pd *PodmanDriver) ListNetworks(lab string, all bool) error {
	return nil
}

func (pd *PodmanDriver) NetworkExists(name, lab string) (bool, error) {
	return false, nil
}
