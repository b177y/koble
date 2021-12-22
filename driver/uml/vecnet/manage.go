package vecnet

import (
	"fmt"
	"io"
	"os"

	"github.com/containernetworking/plugins/pkg/ns"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func SetupMgmtIface(machine, namespace, sockpath string) (ifaceName string, err error) {
	bridgeAlias := fmt.Sprintf("mgmt_br_%s", machine)
	tapAlias := fmt.Sprintf("mgmt_tap_%s", machine)
	slirpAlias := fmt.Sprintf("mgmt_slirp_%s", machine)
	err = WithNetNS(namespace, func(ns.NetNS) error {
		err = NewBridge(bridgeAlias)
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
		return "", err
	}
	err = AddSlirpIface(slirpAlias, bridgeAlias, namespace, "10.22.2.0/24", sockpath)
	return ifaceName, err
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
		conn, err := ssh.Dial("tcp", "10.22.2.110:46222", sshConfig)
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
