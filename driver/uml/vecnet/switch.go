package vecnet

import (
	"fmt"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/songgao/water"
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
	fmt.Printf("Creating bridge %s in namespace %s\n", brname, namespace)
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
		fmt.Println("Created bridge, now setting UP")
		return netlink.LinkSetUp(newBr)
	})
}

// TODO Remove bridge

func AddHost(namespace, tapname, bridge string) error {
	fmt.Printf("Creating tap %s (->%s) in namespace %s\n", tapname, bridge, namespace)
	return WithNetNS(namespace, func(ns.NetNS) error {
		br, err := netlink.LinkByName(bridge)
		if err != nil {
			return fmt.Errorf("Error finding %s: %w", bridge, err)
		}
		fmt.Println("found bridge, now creating tap")
		config := water.Config{
			DeviceType: water.TAP,
		}
		config.Name = tapname
		config.Persist = true
		_, err = water.New(config)
		if err != nil {
			return err
		}
		fmt.Println("should have made tap, now tryna find it")
		tap, err := netlink.LinkByName(tapname)
		if err != nil {
			return err
		}
		fmt.Println("found tap, now connecting to bridge")
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
