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
	log "github.com/sirupsen/logrus"
)

type PodmanDriver struct {
	conn context.Context
	Name string
}

func (pd *PodmanDriver) SetupDriver() (err error) {
	pd.Name = "Podman"
	log.Debug("Attempting to connect to podman socket.")
	pd.conn, err = bindings.NewConnection(context.Background(), "unix://run/user/1000/podman/podman.sock")
	if err != nil {
		return driver.NewDriverError(err, pd.Name, "SetupDriver")
	}
	return nil
}

func getName(name, lab string) string {
	name = "netkit_" + name
	if lab != "" {
		name += "_" + lab
	}
	return name
}

func getLabels(name, lab string) map[string]string {
	labels := make(map[string]string)
	labels["netkit"] = "true"
	labels["netkit:name"] = name
	if lab != "" {
		labels["netkit:lab"] = lab
	} else {
		labels["netkit:nolab"] = "true"
	}

	return labels
}

func getFilters(machine, lab string, all bool) map[string][]string {
	filters := make(map[string][]string)
	var labelFilters []string
	labelFilters = append(labelFilters, "netkit=true")
	if lab != "" && !all {
		labelFilters = append(labelFilters, "netkit:lab="+lab)
	} else if !all {
		labelFilters = append(labelFilters, "netkit:nolab=true")
	}
	if machine != "" && !all {
		labelFilters = append(labelFilters, "netkit:name="+machine)
	}
	filters["label"] = labelFilters
	return filters
}

func (pd *PodmanDriver) MachineExists(name string) (exists bool,
	err error) {
	exists, err = containers.Exists(pd.conn, name, nil)
	if err != nil {
		return exists, driver.NewDriverError(err, pd.Name, "MachineExists")
	}
	return exists, nil
}

func (pd *PodmanDriver) StartMachine(m driver.Machine, lab string) (id string, err error) {
	name := getName(m.Name, lab)
	exists, err := images.Exists(pd.conn, m.Image, nil)
	if err != nil {
		return "", driver.NewDriverError(err, pd.Name, "StartMachine")
	}
	if !exists {
		fmt.Println("Image", m.Image, "does not already exist, attempting to pull...")
		_, err = images.Pull(pd.conn, m.Image, nil)
		if err != nil {
			return "", err
		}
	}
	s := specgen.NewSpecGenerator(m.Image, false)
	s.Name = name
	s.Hostname = name
	s.Command = []string{"/sbin/init"}
	s.CapAdd = []string{"NET_ADMIN", "SYS_ADMIN", "CAP_NET_BIND_SERVICE", "CAP_NET_RAW", "CAP_SYS_NICE", "CAP_IPC_LOCK", "CAP_CHOWN"}
	s.CNINetworks = m.Networks
	s.Terminal = true
	s.Labels = getLabels(m.Name, lab)
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

func (pd *PodmanDriver) MachineExecShell(name, command, user string,
	detach bool, workdir string) (err error) {
	exists, err := pd.MachineExists(name)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", name)
	}
	ec := new(handlers.ExecCreateConfig)
	ec.Cmd = []string{command}
	exId, err := containers.ExecCreate(pd.conn, name, ec)
	if err != nil {
		return err
	}
	options := new(containers.ExecStartAndAttachOptions)
	options.WithOutputStream(io.WriteCloser(os.Stdout))
	options.WithAttachOutput(true)
	options.WithErrorStream(io.WriteCloser(os.Stderr))
	options.WithAttachError(true)
	options.WithInputStream(*bufio.NewReader(os.Stdin))
	options.WithAttachInput(true)
	err = containers.ExecStartAndAttach(pd.conn, exId, options)
	return err
}

func (pd *PodmanDriver) GetMachineLogs(name, lab string,
	stdoutChan, stderrChan chan string,
	follow bool, tail int) (err error) {
	name = getName(name, lab)
	exists, err := pd.MachineExists(name)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", name)
	}
	opts := new(containers.LogOptions)
	opts.WithStdout(true)
	opts.WithStderr(true)
	opts.WithTail(fmt.Sprint(tail))
	opts.WithFollow(follow)
	err = containers.Logs(pd.conn, name, opts, stdoutChan, stderrChan)
	return err
}

func (pd *PodmanDriver) ListMachines(lab string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	opts := new(containers.ListOptions)
	if lab == "" {
	}
	filters := getFilters("", lab, all)
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.conn, opts)
	if err != nil {
		return machines, err
	}
	for _, c := range ctrs {
		machines = append(machines, driver.MachineInfo{
			Name:     c.Names[0],
			Image:    c.Image,
			Networks: c.Networks,
			State:    c.State,
			Uptime:   c.Status,
			Exited:   c.Exited,
			ExitCode: c.ExitCode,
			ExitedAt: c.ExitedAt,
			Mounts:   c.Mounts,
			HostPid:  c.Pid,
			Ports:    c.Ports,
		})
	}
	return machines, nil
}
