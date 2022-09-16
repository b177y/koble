package podman

import (
	"fmt"
	"net"

	"github.com/b177y/koble/pkg/driver"
	"github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v4/pkg/bindings/network"
)

func (n *Network) getNetLabels() map[string]string {
	labels := make(map[string]string)
	labels["koble"] = "true"
	labels["koble:name"] = n.Name()
	labels["koble:driver"] = "podman"
	labels["koble:namespace"] = n.Namespace
	return labels
}

type Network struct {
	name      string
	Namespace string
	pd        *PodmanDriver
}

func (n *Network) Name() string {
	return n.name
}

func (n *Network) Id() string {
	return "koble." + n.Namespace + "." + n.name + "." + n.pd.DriverName
}

func (n *Network) Create(opts *driver.NetConfig) (err error) {
	exists, err := n.Exists()
	if err != nil {
		return err
	}
	if exists {
		return driver.ErrExists
	}
	if opts.External {
		if opts.Subnet == "" || opts.Gateway == "" {
			return fmt.Errorf("gateway and subnet must be specified for network %s", n.name)
		}
		_, _, err := net.ParseCIDR(opts.Subnet)
		if err != nil {
			return err
		}
		gw := net.ParseIP(opts.Gateway)
		if gw == nil {
			return fmt.Errorf("Could not parse IP %s as gateway", opts.Gateway)
		}
	}
	newNet := &types.Network{
		Name: n.Id(),
		//Subnets: []types.Subnet{opts.Subnet}, // TODO
		//IPv6Enabled: opts.IPv6,
		Internal: !opts.External,
	}
	if !opts.External {
		ipamOpts := map[string]string{}
		ipamOpts["driver"] = "none"
		newNet.IPAMOptions = ipamOpts
	}
	newNet.Labels = map[string]string{}
	newNet.Labels["koble"] = "true"
	newNet.Labels["koble:driver"] = n.pd.DriverName
	newNet.Labels["koble:namespace"] = n.Namespace
	_, err = network.Create(n.pd.Conn, newNet)
	return err
}

func (n *Network) Start() (err error) {
	// podman network doesn't need manual starting
	return nil
}

func (n *Network) Remove() (err error) {
	_, err = network.Remove(n.pd.Conn, n.Id(), nil)
	return err
}

func (n *Network) Stop() (err error) {
	// podman network doesn't need manual stopping
	return nil
}

func (n *Network) Running() (running bool, err error) {
	return n.Exists()
}

func (pd *PodmanDriver) ListNetworks(namespace string, all bool) (networks []driver.NetInfo, err error) {
	opts := new(network.ListOptions)
	filters := getFilters("", namespace, pd.DriverName, all)
	opts.WithFilters(filters)
	nets, err := network.List(pd.Conn, opts)
	if err != nil {
		return networks, err
	}
	for _, n := range nets {
		networks = append(networks, driver.NetInfo{
			Name:      n.Name, // TODO cut 3rd element of name
			Namespace: n.Name, // TODO cut 2nd element of name
			External:  !n.Internal,
			Gateway:   "", // TODO
			IpRange:   "", // TODO
			Subnet:    "", // TODO
			IPv6:      "", // TODO
		})
	}
	return networks, nil
}

func (n *Network) Exists() (bool, error) {
	return network.Exists(n.pd.Conn, n.Id(), nil)
}

func (n *Network) Info() (nInfo driver.NetInfo, err error) {
	exists, err := n.Exists()
	if err != nil {
		return nInfo, err
	}
	if !exists {
		return nInfo, driver.ErrNotExists
	}
	info, err := network.Inspect(n.pd.Conn, n.Id(), nil)
	if err != nil {
		return nInfo, err
	}
	nInfo, err = netInfoFromInspect(*n, info)
	return nInfo, err
}

func netInfoFromInspect(nw Network, insp types.Network) (netInfo driver.NetInfo, err error) {
	netInfo = driver.NetInfo{
		Name:      nw.Name(),
		Namespace: nw.Namespace,
	}
	fmt.Printf("found network %s\n", insp)
	// if v, ok := insp[0]["plugins"]; ok {
	// 	parsed := v.([]interface{})
	// 	basicInfo := parsed[0].(map[string]interface{})
	// 	// if v, ok := basicInfo["bridge"]; ok {
	// 	// 	netInfo.Interface = v.(string)
	// 	// }
	// 	if v, ok := basicInfo["ipam"]; ok {
	// 		ipamParsed := v.(map[string]interface{})
	// 		if v, ok := ipamParsed["isGateway"]; ok {
	// 			netInfo.External = v.(bool)
	// 		}
	// 		if v, ok := ipamParsed["ranges"]; ok {
	// 			rangesMap := v.([]interface{})[0].([]interface{})[0].(map[string]interface{})
	// 			if v, ok := rangesMap["gateway"]; ok {
	// 				netInfo.Gateway = v.(string)
	// 			}
	// 			if v, ok := rangesMap["subnet"]; ok {
	// 				netInfo.Subnet = v.(string)
	// 			}
	// 		}
	// 	}
	// }
	return netInfo, err
}
