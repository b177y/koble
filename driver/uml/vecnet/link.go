package vecnet

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

var IDLEN = 15
var IDCHARSET = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func contains(list []string, toMatch string) bool {
	for _, item := range list {
		if item == toMatch {
			return true
		}
	}
	return false
}

func genID() string {
	sRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, IDLEN)
	for i := range b {
		b[i] = IDCHARSET[sRand.Intn(len(IDCHARSET))]
	}
	return string(b)
}

func newLinkID() (id string, err error) {
	devs, err := netlink.LinkList()
	var takenNames []string
	for _, d := range devs {
		takenNames = append(takenNames, d.Attrs().Name)
	}
	var newId string
	newId = genID()
	counter := 0
	for contains(takenNames, newId) {
		newId = genID()
		counter++
		if counter > 10 {
			return "", errors.New("Couldn't create a unique interface name after 10 attempts")
		}
	}
	return newId, nil
}

func IfaceExistsByAlias(alias string) (exists bool, err error) {
	devs, err := netlink.LinkList()
	if err != nil {
		return false, err
	}
	for _, d := range devs {
		if d.Attrs().Alias == alias {
			return true, nil
		}
	}
	return false, nil
}

// Create a new bridge within specified net namespace
func NewBridge(name string) error {
	// check if exists
	exists, err := IfaceExistsByAlias(name)
	if err != nil {
		return err
	} else if exists {
		// bridge already exists
		return nil
	}
	id, err := newLinkID()
	if err != nil {
		return err
	}
	la := netlink.NewLinkAttrs()
	la.Name = id
	la.MTU = 1500
	la.TxQLen = -1
	newBr := &netlink.Bridge{
		LinkAttrs: la,
	}
	err = netlink.LinkAdd(newBr)
	if err != nil {
		return err
	}
	fmt.Println("created bridge")
	err = netlink.LinkSetAlias(newBr, name)
	if err != nil {
		return err
	}
	fmt.Println("set bridge alias", newBr.LinkAttrs.Alias)
	return netlink.LinkSetUp(newBr)
}

// Remove a bridge from a specified net namespace
func DelBridge(alias string) error {
	br, err := netlink.LinkByAlias(alias)
	if err != nil {
		return fmt.Errorf("Could not find bridge %s to delete: %w", alias, err)
	}
	return netlink.LinkDel(br)
}

func NewTap(alias string) (ifaceName string, err error) {
	exists, err := IfaceExistsByAlias(alias)
	if err != nil {
		return "", err
	} else if exists {
		// tap already exists
		return "", nil
	}
	ifaceName, err = newLinkID()
	if err != nil {
		return "", err
	}
	tap := &netlink.Tuntap{
		LinkAttrs: netlink.LinkAttrs{
			Name: ifaceName,
		},
		Mode:  netlink.TUNTAP_MODE_TAP,
		Owner: 0,
		Group: 0,
	}
	err = netlink.LinkAdd(tap)
	if err != nil {
		return "", err
	}
	err = netlink.LinkSetAlias(tap, alias)
	if err != nil {
		return "", err
	}
	err = netlink.LinkSetUp(tap)
	return ifaceName, err
}

func AddTapToBridge(tapAlias, bridgeAlias string) (err error) {
	br, err := netlink.LinkByAlias(bridgeAlias)
	if err != nil {
		return fmt.Errorf("Error finding bridge %s: %w", bridgeAlias, err)
	}
	tap, err := netlink.LinkByAlias(tapAlias)
	if err != nil {
		return fmt.Errorf("Error finding bridge %s: %w", tapAlias, err)
	}
	return netlink.LinkSetMaster(tap, br)
}

func DelTap(alias string) error {
	tap, err := netlink.LinkByAlias(alias)
	if err != nil {
		return fmt.Errorf("Error finding tap %s to delete: %w", alias, err)
	}
	return netlink.LinkDel(tap)
}

func getSlirpArgs(nsPath, iface, subnet, sockpath string) (args []string) {
	args = []string{"--configure", "--mtu=65520", "--disable-host-loopback",
		"--disable-dns", "--netns-type=path"}
	if subnet != "" {
		args = append(args, "--cidr", subnet)
	}
	if sockpath != "" {
		args = append(args, "--api-socket", sockpath)
	}
	args = append(args, nsPath, iface)
	return args
}

func AddSlirpIface(name, bridge, namespace, subnet, sockpath string) error {
	nsPath := filepath.Join("/run/user", os.Getenv("UML_ORIG_UID"), "uml/ns", namespace, "netns.bind")
	ifaceName, err := newLinkID()
	if err != nil {
		return err
	}
	cmd := exec.Cmd{
		Path: "/usr/bin/slirp4netns",
		Args: getSlirpArgs(nsPath, ifaceName, subnet, sockpath),
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	return WithNetNS(namespace, func(ns.NetNS) error {
		var tap netlink.Link
		var err error
		exists, err := IfaceExistsByAlias(name)
		if err != nil {
			return err
		} else if exists {
			// TODO change this bc iface might exist but slirp might have stopped
			// tapout already exists
			return nil
		}
		// timeout ~5s to wait for slirp4netns to create new tap
		for i := 0; i < 10; i++ {
			tap, err = netlink.LinkByName(ifaceName)
			if err == nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		if err != nil {
			return fmt.Errorf("Error finding slirp tap %s: %w", ifaceName, err)
		}
		br, err := netlink.LinkByAlias(bridge)
		if err != nil {
			return fmt.Errorf("Error finding bridge %s: %w", bridge, err)
		}
		err = netlink.LinkSetUp(tap)
		if err != nil {
			return err
		}
		err = netlink.LinkSetAlias(tap, name)
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
		// TODO KILL SLIRP PROCESS
		return netlink.LinkDel(tap)
	})
}
