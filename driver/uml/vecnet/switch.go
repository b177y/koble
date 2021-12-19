package vecnet

import (
	"fmt"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
)

// Create a new bridge within the uml net namespace
func NewBridge(name string) error {
	return WithNetNS(func(ns.NetNS) error {
		la := netlink.NewLinkAttrs()
		la.Name = name
		mybridge := &netlink.Bridge{LinkAttrs: la}
		return netlink.LinkAdd(mybridge)
	})
}

func AddHost(machine, network string) error {
	// ip tuntap add tap X
	// ip link set up
	// brctl addif to bridge
	return WithNetNS(func(ns.NetNS) error {
		br, err := netlink.LinkByName(network)
		if err != nil {
			return fmt.Errorf("Error finding %s: %w", network, err)
		}
		config := water.Config{
			DeviceType: water.TAP,
		}
		config.Name = machine
		config.Persist = true
		_, err = water.New(config)
		if err != nil {
			return err
		}
		fmt.Println("should have made tap, now tryna find it")
		tap, err := netlink.LinkByName(machine)
		if err != nil {
			return err
		}
		return netlink.LinkSetMaster(tap, br)
	})
}

func AddTapout(iface, network, namespace string) error {
	// exec slirp4netns with params
	// create bridge
	// add host to bridge
	// add slirp's iface to bridge
	return nil
}
