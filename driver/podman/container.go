package podman

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/b177y/netkit/driver"
	"github.com/containers/podman/v3/pkg/api/handlers"
	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/specgen"
)

type PodmanDriver struct {
	conn context.Context
}

func (pd *PodmanDriver) SetupDriver() (err error) {
	pd.conn, err = bindings.NewConnection(context.Background(), "unix://run/user/1000/podman/podman.sock")
	return err
}

func (pd *PodmanDriver) StartMachine(m driver.Machine) (id string, err error) {
	fmt.Println("Checking if image exists")
	exists, err := images.Exists(pd.conn, m.Image, nil)
	if err != nil {
		return "", err
	}
	if !exists {
		fmt.Println("Image", m.Image, "does not already exist, attempting to pull...")
		_, err = images.Pull(pd.conn, m.Image, nil)
		if err != nil {
			return "", err
		}
	}
	fmt.Println("new spec")
	s := specgen.NewSpecGenerator(m.Image, false)
	s.Name = m.Name
	s.Hostname = m.Name
	s.Command = []string{"/sbin/init"}
	s.CapAdd = []string{"NET_ADMIN", "SYS_ADMIN", "CAP_NET_BIND_SERVICE", "CAP_NET_RAW", "CAP_SYS_NICE", "CAP_IPC_LOCK", "CAP_CHOWN"}
	s.CNINetworks = m.Networks
	s.Terminal = true
	createResponse, err := containers.CreateWithSpec(pd.conn, s, nil)
	if err != nil {
		return "", err
	}
	err = containers.Start(pd.conn, createResponse.ID, nil)
	if err != nil {
		return createResponse.ID, err
	}
	return createResponse.ID, nil
}

func (pd *PodmanDriver) StopMachine(name string) error {
	err := containers.Stop(pd.conn, name, nil)
	return err
}

func (pd *PodmanDriver) CrashMachine(name string) error {
	err := containers.Kill(pd.conn, name, nil)
	return err
}

func (pd *PodmanDriver) GetMachineStatus(name string) (data interface{}, err error) {
	data, err = containers.Inspect(pd.conn, name, nil)
	return data, err
}

func (pd *PodmanDriver) AttachToMachine(name string) (err error) {
	opts := new(containers.AttachOptions)
	fmt.Printf("Attaching to %s, Use key sequence <ctrl><p>, <ctrl><q> to detach.\n", name)
	fmt.Printf("You might need to hit <enter> once attached to get a prompt.\n\n")
	err = containers.Attach(pd.conn, name, os.Stdin, os.Stdout, os.Stderr, nil, opts)
	return err
}

func (pd *PodmanDriver) MachineExecShell(name string) (err error) {
	ec := new(handlers.ExecCreateConfig)
	ec.Cmd = []string{"/bin/bash"}
	exId, err := containers.ExecCreate(pd.conn, name, ec)
	if err != nil {
		return err
	}
	// options := new(containers.ExecStartAndAttachOptions).WithOutputStream(io.WriteCloser(os.Stdout)).WithErrorStream(io.WriteCloser(os.Stderr)).WithInputStream(*bufio.NewReader(os.Stdin)).WithAttachError(true).WithAttachInput(true).WithAttachOutput(true)
	options := new(containers.ExecStartAndAttachOptions)
	options.WithOutputStream(io.WriteCloser(os.Stdout))
	options.WithAttachOutput(true)
	options.WithErrorStream(io.WriteCloser(os.Stderr))
	options.WithAttachError(true)
	options.WithInputStream(*bufio.NewReader(os.Stdin))
	options.WithAttachInput(true)
	err = containers.ExecStartAndAttach(pd.conn, exId, options)
	fmt.Println("Error in execstartandattach is", err)
	return err
}

func (pd *PodmanDriver) GetMachineLogs(name string, stdoutChan, stderrChan chan string) (err error) {
	opts := new(containers.LogOptions)
	opts.WithStdout(true)
	err = containers.Logs(pd.conn, name, nil, stdoutChan, stderrChan)
	return err
}
