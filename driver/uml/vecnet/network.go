package vecnet

import (
	"fmt"
	"strings"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

// ip addresses
// https://github.com/containers/podman/blob/375ff223f430301edf25ef5a5f03a1ae1e029bef/libpod/network/internal/util/util.go
// https://github.com/containers/podman/blob/375ff223f430301edf25ef5a5f03a1ae1e029bef/libpod/network/internal/util/ip.go

func NetExists(name, namespace string) (exists bool, err error) {
	err = WithNetNS(namespace, func(ns.NetNS) error {
		exists, err = IfaceExistsByAlias("nbr_" + name)
		return err
	})
	return exists, err
}

// Create a new bridge within specified net namespace
func NewNet(name, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		return NewBridge("nbr_" + name)
	})
}

func RemoveNet(name, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		return DelBridge("nbr_" + name)
	})
}

func AddHostToNet(machine, network,
	namespace string) (ifaceName string, err error) {
	err = WithNetNS(namespace, func(ns.NetNS) error {
		tapAlias := fmt.Sprintf("mtap_%s_net_%s", machine, network)
		ifaceName, err = NewTap(tapAlias)
		if err != nil {
			return err
		}
		return AddTapToBridge(tapAlias, "nbr_"+network)
	})
	return ifaceName, err
}

func RemoveHostTap(machine, network, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		tapAlias := fmt.Sprintf("mtap_%s_net_%s", machine, network)
		return DelTap(tapAlias)
	})
}

func MakeNetExternal(network, namespace, subnet string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		bridge := "nbr_" + network
		return AddSlirpIface("ext_"+network, bridge, namespace, subnet, "")
	})
}

func bridgeUsed(bridge netlink.Link) (used bool, err error) {
	interfaces, err := netlink.LinkList()
	if err != nil {
		return false, err
	}
	for _, i := range interfaces {
		if i.Attrs().MasterIndex == bridge.Attrs().Index {
			return true, nil
		}
	}
	return false, nil
}

func RemoveMachineNets(machine, namespace string, rmNet bool) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		interfaces, err := netlink.LinkList()
		if err != nil {
			return err
		}
		for _, i := range interfaces {
			if !strings.Contains(i.Attrs().Alias, "mtap_"+machine) {
				continue
			}
			err = netlink.LinkDel(i)
			if err != nil {
				return fmt.Errorf("Could not delete tap %s: %w", i.Attrs().Alias, err)
			}
			if rmNet {
				br, err := netlink.LinkByIndex(i.Attrs().MasterIndex)
				if err != nil {
					return fmt.Errorf("Could not find tap %s's master by index %d: %w",
						i.Attrs().Alias, i.Attrs().MasterIndex, err)
				}
				fmt.Printf("Found bridge %s for tap %s\n", br.Attrs().Alias, i.Attrs().Alias)
				if used, err := bridgeUsed(br); err != nil {
					return err
				} else if !used {
					return netlink.LinkDel(br)
				}
			}
		}
		return nil
	})
}
