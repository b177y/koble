package vecnet

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
func NewBridge(name, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		la := netlink.NewLinkAttrs()
		la.Name = name
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

// Remove a bridge from a specified net namespace
func DelBridge(name, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		br, err := netlink.LinkByName(name)
		if err != nil {
			return fmt.Errorf("Could not find bridge %s to delete: %w", name, err)
		}
		return netlink.LinkDel(br)
	})
}

func AddHost(tapname, bridge, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		br, err := netlink.LinkByName(bridge)
		if err != nil {
			return fmt.Errorf("Error finding bridge %s: %w", bridge, err)
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

func DelHost(tapname, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		tap, err := netlink.LinkByName(tapname)
		if err != nil {
			return fmt.Errorf("Error finding %s: %w", tapname, err)
		}
		return netlink.LinkDel(tap)
	})
}

func SetupExternal(iface, network, namespace string) error {
	nsPath := filepath.Join("/run/user", os.Getenv("UML_ORIG_UID"), "uml/ns", namespace, "netns.bind")
	cmd := exec.Cmd{
		Path: "/usr/bin/slirp4netns",
		Args: []string{"--configure", "--mtu=65520", "--disable-host-loopback",
			"--disable-dns", "--netns-type=path", nsPath, iface},
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	return WithNetNS(namespace, func(ns.NetNS) error {
		var tap netlink.Link
		var err error
		// timeout ~5s to wait for slirp4netns to create new tap
		for i := 0; i < 10; i++ {
			tap, err = netlink.LinkByName(iface)
			if err == nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		if err != nil {
			return fmt.Errorf("Error finding tap %s: %w", network, err)
		}
		br, err := netlink.LinkByName(network)
		if err != nil {
			return fmt.Errorf("Error finding bridge %s: %w", network, err)
		}
		err = netlink.LinkSetUp(tap)
		if err != nil {
			return err
		}
		return netlink.LinkSetMaster(tap, br)
	})
}

func DelExternal(iface, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		tap, err := netlink.LinkByName(iface)
		if err != nil {
			return fmt.Errorf("Error finding tapout %s: %w", iface, err)
		}
		// KILL SLIRP PROCESS
		return netlink.LinkDel(tap)
	})
}
