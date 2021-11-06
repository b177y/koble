package podman

import (
	"fmt"
	"net"

	"github.com/b177y/netkit/driver"
	"github.com/containers/podman/v3/pkg/bindings/network"
)

func (pd *PodmanDriver) CreateNetwork(n driver.Network, lab string) (id string,
	err error) {
	exists, err := pd.NetworkExists(n.Name, lab)
	if err != nil {
		return "", err
	}
	if exists {
		return "", driver.ErrExists
	}
	opts := new(network.CreateOptions)
	opts.WithName(getName(n.Name, lab))
	opts.WithLabels(getLabels(n.Name, lab))
	if n.Subnet != "" && n.Gateway != "" {
		_, sn, err := net.ParseCIDR(n.Subnet)
		if err != nil {
			return "", err
		}
		gw := net.ParseIP(n.Gateway)
		if gw == nil {
			return "", fmt.Errorf("Could not parse IP %s as Gateway", n.Gateway)
		}
		opts.WithGateway(gw)
		opts.WithSubnet(*sn)
	}
	opts.WithInternal(!n.External)
	_, err = network.Create(pd.conn, opts)
	return id, err
}

func (pd *PodmanDriver) StartNetwork(name, lab string) (err error) {
	return nil
}

func (pd *PodmanDriver) RemoveNetwork(name, lab string) (err error) {
	nk_fullname := getName(name, lab)
	_, err = network.Remove(pd.conn, nk_fullname, nil)
	return err
}

func (pd *PodmanDriver) StopNetwork(name, lab string) (err error) {
	return nil
}

func (pd *PodmanDriver) GetNetworkState(name, lab string) (state string,
	err error) {
	nk_fullname := getName(name, lab)
	network, err := network.Inspect(pd.conn, nk_fullname, nil)
	fmt.Println("NETWORK INCLUDES", network)
	return "bruh", err
}

func (pd *PodmanDriver) ListNetworks(lab string, all bool) error {
	return nil
}

func (pd *PodmanDriver) NetworkExists(name, lab string) (bool, error) {
	nk_fullname := getName(name, lab)
	return network.Exists(pd.conn, nk_fullname, nil)
}
