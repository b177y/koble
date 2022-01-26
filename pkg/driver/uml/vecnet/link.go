package vecnet

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/containernetworking/plugins/pkg/ns"
	log "github.com/sirupsen/logrus"
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
	err = netlink.LinkSetAlias(newBr, name)
	if err != nil {
		return err
	}
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
	t, err := netlink.LinkByAlias(alias)
	if err == nil {
		return t.Attrs().Name, nil
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

type process struct {
	pid     int
	cmdline string
}

func getProcesses() (pList []process, err error) {
	dirs, err := ioutil.ReadDir("/proc")
	if err != nil {
		return pList, err
	}
	for _, entry := range dirs {
		if pid, err := strconv.Atoi(entry.Name()); err == nil {
			cmdline, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
			if err != nil {
				log.Tracef("Could not read /proc/%d/cmdline: %v\n", pid, err)
				continue
			}
			pList = append(pList, process{
				pid:     pid,
				cmdline: strings.TrimSuffix(string(cmdline), "\n"),
			})
		}
	}
	return pList, nil
}

func RemoveSlirp(match string) error {
	processes, err := getProcesses()
	if err != nil {
		return nil
	}
	for _, p := range processes {
		if strings.Contains(p.cmdline, match) {
			return syscall.Kill(p.pid, syscall.SIGKILL)
		}
	}
	return fmt.Errorf("Could not find slirp socket to remove")
}

func AddSlirpIface(name, bridge, namespace, subnet, sockpath string) error {
	// remove existing slirp for machine first
	RemoveSlirp(sockpath)
	nsPath := filepath.Join("/run/user", os.Getenv("UML_ORIG_UID"), "uml/ns", namespace, "netns.bind")
	ifaceName, err := newLinkID()
	if err != nil {
		return err
	}
	sp, err := exec.LookPath("slirp4netns")
	if err != nil {
		return err
	}
	cmd := exec.Cmd{
		Path: sp,
		Args: getSlirpArgs(nsPath, ifaceName, subnet, sockpath),
	}
	// TODO find a better way to get errors from start
	// if sockpath == "" {
	// 	fmt.Println("running slirp with:", cmd)
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Stderr = os.Stderr
	// 	err = cmd.Run()
	// } else {
	err = cmd.Start()
	// }
	if err != nil {
		return fmt.Errorf("could not start slirp4netns with args %v: %w", cmd.Args, err)
	}
	return WithNetNS(namespace, func(ns.NetNS) error {
		var tap netlink.Link
		var err error
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
