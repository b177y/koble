package vecnet

import (
	"io"
	"os"

	"github.com/containernetworking/plugins/pkg/ns"
	"golang.org/x/crypto/ssh"
)

func SetupMgmtIface(machine, namespace, sockpath string) error {
	err := NewBridge(machine+"nkbr", namespace)
	if err != nil {
		return err
	}
	err = AddHost(machine+"nkmg", machine+"nkbr", namespace)
	if err != nil {
		return err
	}
	return SetupExternal(machine+"sl0", machine+"nkbr", namespace, "10.22.2.0/24", sockpath)
}

func ExecCommand(machine, user, command, namespace string) error {
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
		sessStdOut, err := sess.StdoutPipe()
		if err != nil {
			return err
		}
		go io.Copy(os.Stdout, sessStdOut)
		sessStderr, err := sess.StderrPipe()
		if err != nil {
			return err
		}
		go io.Copy(os.Stderr, sessStderr)
		err = sess.Run(command)
		if err != nil {
			return err
		}
		return nil
	})
}

// TODO - for now we only need to create portmappings when machines are started,
// and theres no feature for adding / removing mappings during runtime
// func RemoveForward(hostPort uint16, sockpath string) error {
// 	return nil
// }
