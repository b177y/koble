package podman

import (
	"fmt"
	"net"

	"github.com/b177y/netkit/driver"
	"github.com/containers/podman/v3/pkg/bindings/network"
	"github.com/containers/podman/v3/pkg/domain/entities"
)

func getNetLabels(n driver.Network) map[string]string {
	labels := make(map[string]string)
	labels["netkit"] = "true"
	labels["netkit:name"] = n.Name
	if n.Lab != "" {
		labels["netkit:lab"] = n.Lab
	} else {
		labels["netkit:nolab"] = "true"
	}
	labels["netkit:namespace"] = n.Namespace
	return labels
}

func (pd *PodmanDriver) CreateNetwork(n driver.Network) (err error) {
	exists, err := pd.NetworkExists(n)
	if err != nil {
		return err
	}
	if exists {
		return driver.ErrExists
	}
	opts := new(network.CreateOptions)
	opts.WithName(n.Fullname())
	opts.WithLabels(getNetLabels(n))
	if n.Subnet != "" && n.Gateway != "" {
		_, sn, err := net.ParseCIDR(n.Subnet)
		if err != nil {
			return err
		}
		gw := net.ParseIP(n.Gateway)
		if gw == nil {
			return fmt.Errorf("Could not parse IP %s as Gateway", n.Gateway)
		}
		opts.WithGateway(gw)
		opts.WithSubnet(*sn)
	}
	opts.WithInternal(!n.External)
	_, err = network.Create(pd.conn, opts)
	return err
}

func (pd *PodmanDriver) StartNetwork(net driver.Network) (err error) {
	// podman network doesn't need manual starting
	return nil
}

func (pd *PodmanDriver) RemoveNetwork(net driver.Network) (err error) {
	_, err = network.Remove(pd.conn, net.Fullname(), nil)
	return err
}

func (pd *PodmanDriver) StopNetwork(net driver.Network) (err error) {
	// podman network doesn't need manual stopping
	return nil
}

func (pd *PodmanDriver) GetNetworkState(net driver.Network) (state driver.NetworkState,
	err error) {
	state.Running, err = pd.NetworkExists(net)
	return state, err
}

func (pd *PodmanDriver) ListNetworks(lab string, all bool) (networks []driver.NetInfo, err error) {
	opts := new(network.ListOptions)
	filters := getFilters("", lab, "GLOBAL", all)
	opts.WithFilters(filters)
	nets, err := network.List(pd.conn, opts)
	if err != nil {
		return networks, err
	}
	for _, n := range nets {
		name, lab := getInfoFromLabels(n.Labels)
		n := driver.Network{
			Name: name,
			Lab:  lab,
		}
		info, err := network.Inspect(pd.conn, n.Fullname(), nil)
		if err != nil {
			return networks, err
		}
		nw := driver.NetInfo{
			Name: name,
			Lab:  lab,
		}
		// this is currently very cursed due to podman bindings at v3.4
		// returning map[string]interface{}
		// future bindings will return
		// https://github.com/containers/podman/blob/abbd6c167e8163a711680db80137a0731e06e564/libpod/network/types/network.go#L34
		// update this code to make it cleaner when this is released :)
		if v, ok := info[0]["plugins"]; ok {
			parsed := v.([]interface{})
			basicInfo := parsed[0].(map[string]interface{})
			if v, ok := basicInfo["bridge"]; ok {
				nw.Interface = v.(string)
			}
			if v, ok := basicInfo["ipam"]; ok {
				ipamParsed := v.(map[string]interface{})
				if v, ok := ipamParsed["isGateway"]; ok {
					nw.External = v.(bool)
				}
				if v, ok := ipamParsed["ranges"]; ok {
					rangesMap := v.([]interface{})[0].([]interface{})[0].(map[string]interface{})
					if v, ok := rangesMap["gateway"]; ok {
						nw.Gateway = v.(string)
					}
					if v, ok := rangesMap["subnet"]; ok {
						nw.Subnet = v.(string)
					}
				}
			}
		}
		networks = append(networks, nw)
	}
	return networks, nil
}

func (pd *PodmanDriver) NetworkExists(net driver.Network) (bool, error) {
	return network.Exists(pd.conn, net.Fullname(), nil)
}

func (pd *PodmanDriver) NetInfo(net driver.Network) (nInfo driver.NetInfo, err error) {
	exists, err := pd.NetworkExists(net)
	if err != nil {
		return nInfo, err
	}
	if !exists {
		return nInfo, driver.ErrNotExists
	}
	info, err := network.Inspect(pd.conn, net.Fullname(), nil)
	if err != nil {
		return nInfo, err
	}
	nInfo, err = netInfoFromInspect(net, info)
	return nInfo, err
}

func netInfoFromInspect(nw driver.Network, insp []entities.NetworkInspectReport) (netInfo driver.NetInfo, err error) {
	// this is currently very cursed due to podman bindings at v3.4
	// returning map[string]interface{}
	// future bindings will return
	// https://github.com/containers/podman/blob/abbd6c167e8163a711680db80137a0731e06e564/libpod/network/types/network.go#L34
	// update this code to make it cleaner when this is released :)
	netInfo = driver.NetInfo{
		Name: nw.Name,
		Lab:  nw.Lab,
	}
	if v, ok := insp[0]["plugins"]; ok {
		parsed := v.([]interface{})
		basicInfo := parsed[0].(map[string]interface{})
		if v, ok := basicInfo["bridge"]; ok {
			netInfo.Interface = v.(string)
		}
		if v, ok := basicInfo["ipam"]; ok {
			ipamParsed := v.(map[string]interface{})
			if v, ok := ipamParsed["isGateway"]; ok {
				netInfo.External = v.(bool)
			}
			if v, ok := ipamParsed["ranges"]; ok {
				rangesMap := v.([]interface{})[0].([]interface{})[0].(map[string]interface{})
				if v, ok := rangesMap["gateway"]; ok {
					netInfo.Gateway = v.(string)
				}
				if v, ok := rangesMap["subnet"]; ok {
					netInfo.Subnet = v.(string)
				}
			}
		}
	}
	return netInfo, err
}
