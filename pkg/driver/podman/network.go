package podman

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"text/template"

	"github.com/b177y/koble/pkg/driver"
	"github.com/containers/podman/v3/pkg/bindings/network"
	"github.com/containers/podman/v3/pkg/domain/entities"
)

func (n *Network) getNetLabels() map[string]string {
	labels := make(map[string]string)
	labels["koble"] = "true"
	labels["koble:name"] = n.Name()
	// if n.Lab != "" {
	// 	labels["koble:lab"] = n.Lab
	// } else {
	// 	labels["koble:nolab"] = "true"
	// }
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
	return "koble." + n.Namespace + "." + n.name
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
		cOpts := new(network.CreateOptions)
		if opts.Subnet != "" && opts.Gateway != "" {
			_, sn, err := net.ParseCIDR(opts.Subnet)
			if err != nil {
				return err
			}
			gw := net.ParseIP(opts.Gateway)
			if gw == nil {
				return fmt.Errorf("Could not parse IP %s as Gateway", opts.Gateway)
			}
			cOpts.WithGateway(gw)
			cOpts.WithSubnet(*sn)
		}
		cOpts.WithName(n.Id())
		cOpts.WithLabels(n.getNetLabels())
		cOpts.WithInternal(false)
		_, err = network.Create(n.pd.Conn, cOpts)
		return err
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		// TODO check ~/.config/cni/net.d/cni.lock ??
		f, err := os.Create(filepath.Join(home, ".config", "cni", "net.d", n.Id()+".conflist"))
		defer f.Close()
		if err != nil {
			return err
		}
		tmpl, err := template.New("netconf").Parse(NET)
		if err != nil {
			return err
		}
		err = tmpl.Execute(f, n)
		return err
	}
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
	// opts := new(network.ListOptions)
	// filters := getFilters("", lab, "GLOBAL", all)
	// opts.WithFilters(filters)
	// nets, err := network.List(pd.conn, opts)
	// if err != nil {
	// 	return networks, err
	// }
	// for _, n := range nets {
	// 	name, namespace, lab := getInfoFromLabels(n.Labels)
	// 	n := driver.Network{
	// 		Name:      name,
	// 		Namespace: namespace,
	// 		Lab:       lab,
	// 	}
	// 	info, err := network.Inspect(pd.conn, n.Fullname(), nil)
	// 	if err != nil {
	// 		return networks, err
	// 	}
	// 	nw := driver.NetInfo{
	// 		Name: name,
	// 		Lab:  lab,
	// 	}
	// this is currently very cursed due to podman bindings at v3.4
	// returning map[string]interface{}
	// future bindings will return
	// https://github.com/containers/podman/blob/abbd6c167e8163a711680db80137a0731e06e564/libpod/network/types/network.go#L34
	// update this code to make it cleaner when this is released :)
	// if v, ok := info[0]["plugins"]; ok {
	// 	parsed := v.([]interface{})
	// 	basicInfo := parsed[0].(map[string]interface{})
	// if v, ok := basicInfo["bridge"]; ok {
	// 	nw.Interface = v.(string)
	// }
	// if v, ok := basicInfo["ipam"]; ok {
	// 	ipamParsed := v.(map[string]interface{})
	// 	if v, ok := ipamParsed["isGateway"]; ok {
	// 		nw.External = v.(bool)
	// 	}
	// 	if v, ok := ipamParsed["ranges"]; ok {
	// 		rangesMap := v.([]interface{})[0].([]interface{})[0].(map[string]interface{})
	// 		if v, ok := rangesMap["gateway"]; ok {
	// 			nw.Gateway = v.(string)
	// 		}
	// 		if v, ok := rangesMap["subnet"]; ok {
	// 			nw.Subnet = v.(string)
	// 		}
	// 	}
	// }
	// }
	// networks = append(networks, nw)
	// }
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

func netInfoFromInspect(nw Network, insp []entities.NetworkInspectReport) (netInfo driver.NetInfo, err error) {
	// this is currently very cursed due to podman bindings at v3.4
	// returning map[string]interface{}
	// future bindings will return
	// https://github.com/containers/podman/blob/abbd6c167e8163a711680db80137a0731e06e564/libpod/network/types/network.go#L34
	// update this code to make it cleaner when this is released :)
	netInfo = driver.NetInfo{
		Name:      nw.Name(),
		Namespace: nw.Namespace,
	}
	if v, ok := insp[0]["plugins"]; ok {
		parsed := v.([]interface{})
		basicInfo := parsed[0].(map[string]interface{})
		// if v, ok := basicInfo["bridge"]; ok {
		// 	netInfo.Interface = v.(string)
		// }
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
