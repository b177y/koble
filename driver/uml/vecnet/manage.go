package vecnet

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func ipFromSubnet(subnet net.IPNet, num int) string {
	ip := make(net.IP, len(subnet.IP))
	copy(ip, subnet.IP)
	ip4 := ip.To4()
	ip4[3] = byte(num)
	return ip.String() + "/24"
}

// set bridge ipv4 to .101 in the subnet given
func setBridgeIP(bridgeAlias string, ip string) error {
	br, err := netlink.LinkByAlias(bridgeAlias)
	if err != nil {
		return err
	}
	// delete any existing addresses
	addrs, err := netlink.AddrList(br, netlink.FAMILY_V4)
	for _, a := range addrs {
		netlink.AddrDel(br, &a)
	}
	addr, err := netlink.ParseAddr(ip)
	if err != nil {
		return err
	}
	return netlink.AddrAdd(br, addr)
}

// get the ipv4 of a bridge by alias, interface must have exactly one address
func getBridgeIP(bridgeAlias string) (ip net.IP, err error) {
	br, err := netlink.LinkByAlias(bridgeAlias)
	if err != nil {
		return ip, err
	}
	addrs, err := netlink.AddrList(br, netlink.FAMILY_V4)
	if len(addrs) != 1 {
		return ip, fmt.Errorf("bridge %s has %d addresses instead of 1",
			bridgeAlias, len(addrs))
	}
	return addrs[0].IP, nil
}

func getUsedIPs(namespace string, addrFamily int) (used []*net.IPNet, err error) {
	err = WithNetNS(namespace, func(ns.NetNS) error {
		devs, err := netlink.LinkList()
		if err != nil {
			return err
		}
		for _, d := range devs {
			addrs, err := netlink.AddrList(d, addrFamily)
			if err != nil {
				return err
			}
			for _, a := range addrs {
				used = append(used, a.IPNet)
			}
		}
		return nil
	})
	return used, err
}

func SetupMgmtIface(machine, namespace, sockpath string) (ifaceName string, mgmtIp string, err error) {
	bridgeAlias := fmt.Sprintf("mgmt_br_%s", machine)
	tapAlias := fmt.Sprintf("mgmt_tap_%s", machine)
	slirpAlias := fmt.Sprintf("mgmt_slirp_%s", machine)
	used, err := getUsedIPs(namespace, netlink.FAMILY_V4)
	if err != nil {
		return "", "", err
	}
	subnet, err := getFreeIPv4NetworkSubnet(used, net.IP{10, 22, 0, 0})
	if err != nil {
		return "", "", err
	}
	err = WithNetNS(namespace, func(ns.NetNS) error {
		err = NewBridge(bridgeAlias)
		if err != nil {
			return err
		}
		err = setBridgeIP(bridgeAlias, ipFromSubnet(*subnet, 101))
		if err != nil {
			return err
		}
		ifaceName, err = NewTap(tapAlias)
		if err != nil {
			return err
		}
		return AddTapToBridge(tapAlias, bridgeAlias)
	})
	if err != nil {
		return "", "", err
	}
	// TODO check if already running
	err = AddSlirpIface(slirpAlias, bridgeAlias, namespace, subnet.String(), sockpath)
	return ifaceName, ipFromSubnet(*subnet, 110), err
}

func WithSSHSession(machine, user, namespace string, toRun func(s *ssh.Session) error) (err error) {
	return WithNetNS(namespace, func(ns.NetNS) error {
		if user == "" {
			user = "root"
		}
		sshConfig := &ssh.ClientConfig{
			User:            user,
			Auth:            []ssh.AuthMethod{ssh.Password("netkit")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		addr, err := getBridgeIP(fmt.Sprintf("mgmt_br_%s", machine))
		if err != nil {
			return err
		}
		addr.To4()[3] = 110
		conn, err := ssh.Dial("tcp", addr.String()+":46222", sshConfig)
		if err != nil {
			return err
		}
		defer conn.Close()
		sess, err := conn.NewSession()
		if err != nil {
			return err
		}
		defer sess.Close()
		return toRun(sess)
	})
}

func copyStdChans(sess *ssh.Session,
	sOut, sErr, sIn bool) error {
	if sOut {
		sessStdOut, err := sess.StdoutPipe()
		if err != nil {
			return err
		}
		go io.Copy(os.Stdout, sessStdOut)
	}
	if sErr {
		sessStderr, err := sess.StderrPipe()
		if err != nil {
			return err
		}
		go io.Copy(os.Stderr, sessStderr)
	}
	if sIn {
		sessStdin, err := sess.StdinPipe()
		if err != nil {
			return err
		}
		go io.Copy(sessStdin, os.Stdin)
	}
	return nil
}

func ExecCommand(machine, user, command, namespace string) error {
	return WithSSHSession(machine, user,
		namespace, func(sess *ssh.Session) error {
			err := copyStdChans(sess, true, true, false)
			if err != nil {
				return err
			}
			return sess.Run(command)
		})
}

func RunShell(machine, user, namespace string) error {
	return WithSSHSession(machine, user,
		namespace, func(sess *ssh.Session) error {
			fd := int(os.Stdin.Fd())
			state, err := terminal.MakeRaw(fd)
			if err != nil {
				return err
			}
			defer terminal.Restore(fd, state)
			w, h, err := terminal.GetSize(fd)
			if err != nil {
				return fmt.Errorf("terminal get size: %s", err)
			}
			modes := ssh.TerminalModes{
				ssh.ECHO:          1,
				ssh.TTY_OP_ISPEED: 14400,
				ssh.TTY_OP_OSPEED: 14400,
			}
			term := os.Getenv("TERM")
			if term == "" {
				term = "xterm-256color"
			}
			if err := sess.RequestPty(term, h, w, modes); err != nil {
				return fmt.Errorf("session xterm: %s", err)
			}
			sess.Stdout = os.Stdout
			sess.Stderr = os.Stderr
			sess.Stdin = os.Stdin
			err = sess.Shell()
			if err != nil {
				return err
			}
			return sess.Wait()
		})
}
