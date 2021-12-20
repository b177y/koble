package vecnet

import (
	"fmt"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

func IfaceExists(namespace, iface string) bool {
	return WithNetNS(namespace, func(ns.NetNS) error {
		_, err := netlink.LinkByName(iface)
		return err
	}) == nil
}

// Create a new bridge within specified net namespace
func NewBridge(namespace string, brname string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		la := netlink.NewLinkAttrs()
		la.Name = brname
		la.MTU = 1500
		la.TxQLen = -1
		newBr := &netlink.Bridge{
			LinkAttrs: la,
		}
		err := netlink.LinkAdd(newBr)
		if err != nil {
			return err
		}
		return netlink.LinkSetUp(newBr)
	})
}

// TODO Remove bridge

func AddHost(namespace, tapname, bridge string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		br, err := netlink.LinkByName(bridge)
		if err != nil {
			return fmt.Errorf("Error finding %s: %w", bridge, err)
		}
		tap := &netlink.Tuntap{
			LinkAttrs: netlink.LinkAttrs{
				Name: tapname,
			},
			Mode:  netlink.TUNTAP_MODE_TAP,
			Owner: 0,
			Group: 0,
		}
		err = netlink.LinkAdd(tap)
		if err != nil {
			return err
		}
		err = netlink.LinkSetUp(tap)
		if err != nil {
			return err
		}
		return netlink.LinkSetMaster(tap, br)
	})
}

// TODO remove host

func AddTapout(iface, network, namespace string) error {
	// exec slirp4netns with params
	// create bridge
	// add host to bridge
	// add slirp's iface to bridge
	return nil
}

// TODO remove tapout
