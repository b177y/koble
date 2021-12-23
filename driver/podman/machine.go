package podman

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	copier "github.com/containers/buildah/copier"
	"github.com/spf13/cobra"

	"github.com/b177y/netkit/driver"
	"github.com/containers/podman/v3/pkg/api/handlers"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/specgen"
	log "github.com/sirupsen/logrus"
)

func getLabels(m driver.Machine) map[string]string {
	labels := make(map[string]string)
	labels["netkit"] = "true"
	labels["netkit:name"] = m.Name
	if m.Lab != "" {
		labels["netkit:lab"] = m.Lab
	}
	labels["netkit:namespace"] = m.Namespace
	return labels
}

func getInfoFromLabels(labels map[string]string) (name, namespace, lab string) {
	if val, ok := labels["netkit:name"]; ok {
		name = val
	}
	if val, ok := labels["netkit:lab"]; ok {
		lab = val
	}
	if val, ok := labels["netkit:namespace"]; ok {
		namespace = val
	}
	return name, namespace, lab
}

func getFilters(machine, lab, namespace string, all bool) map[string][]string {
	filters := make(map[string][]string)
	var labelFilters []string
	labelFilters = append(labelFilters, "netkit=true")
	labelFilters = append(labelFilters, "netkit:namespace="+namespace)
	if lab != "" && !all {
		labelFilters = append(labelFilters, "netkit:lab="+lab)
	} // else if !all {
	//labelFilters = append(labelFilters, "netkit:nolab=true")
	//}
	if machine != "" && !all {
		labelFilters = append(labelFilters, "netkit:name="+machine)
	}
	filters["label"] = labelFilters
	return filters
}

func (pd *PodmanDriver) MachineExists(m driver.Machine) (exists bool,
	err error) {
	exists, err = containers.Exists(pd.conn, m.Fullname(), nil)
	if err != nil {
		return exists, err
	}
	return exists, nil
}

func (pd *PodmanDriver) StartMachine(m driver.Machine) (err error) {
	exists, err := pd.MachineExists(m)
	if err != nil {
		return err
	}
	if exists {
		state, err := pd.GetMachineState(m)
		if err != nil {
			return err
		}
		if state.Running {
			return nil
		} else {
			prev := log.GetLevel()
			log.SetLevel(log.ErrorLevel)
			err = containers.Start(pd.conn, m.Fullname(), nil)
			log.SetLevel(prev)
			return err
		}
	}
	if m.Image == "" {
		m.Image = pd.DefaultImage
	}
	imExists, err := images.Exists(pd.conn, m.Image, nil)
	if err != nil {
		return err
	}
	if !imExists {
		fmt.Println("Image", m.Image, "does not already exist, attempting to pull...")
		_, err = images.Pull(pd.conn, m.Image, nil)
		if err != nil {
			return err
		}
	}
	s := specgen.NewSpecGenerator(m.Image, false)
	s.Name = m.Fullname()
	s.Hostname = m.Name
	s.Command = []string{"/sbin/init"}
	s.CapAdd = []string{"NET_ADMIN", "SYS_ADMIN", "CAP_NET_BIND_SERVICE", "CAP_NET_RAW", "CAP_SYS_NICE", "CAP_IPC_LOCK", "CAP_CHOWN"}
	s.NetNS = specgen.Namespace{
		NSMode: specgen.Bridge,
	}
	s.UseImageHosts = true
	s.UseImageResolvConf = true
	for _, n := range m.Networks {
		net := driver.Network{
			Name:      n,
			Namespace: m.Namespace,
		}
		s.CNINetworks = append(s.CNINetworks, net.Fullname())
	}
	s.Terminal = true
	s.Labels = getLabels(m)
	for _, mnt := range m.Volumes {
		if mnt.Type == "" {
			mnt.Type = "bind"
		}
		s.Mounts = append(s.Mounts, mnt)
	}
	createResponse, err := containers.CreateWithSpec(pd.conn, s, nil)
	if err != nil {
		return err
	}
	err = pd.CopyInFiles(m, m.Hostlab)
	if err != nil {
		return err
	}
	// temporary fix to https://github.com/containers/podman/issues/12204
	prev := log.GetLevel()
	log.SetLevel(log.ErrorLevel)
	err = containers.Start(pd.conn, createResponse.ID, nil)
	log.SetLevel(prev)
	return err
}

func (pd *PodmanDriver) HaltMachine(m driver.Machine, force bool) error {
	exists, err := pd.MachineExists(m)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist", m.Name)
	}
	state, err := pd.GetMachineState(m)
	if err != nil {
		return err
	}
	if !state.Running {
		return fmt.Errorf("Can't stop %s as it isn't running", m.Name)
	}
	err = containers.Stop(pd.conn, m.Fullname(), nil)
	return err
}

func (pd *PodmanDriver) RemoveMachine(m driver.Machine) error {
	err := containers.Remove(pd.conn, m.Fullname(), nil)
	return err
}

func (pd *PodmanDriver) GetMachineState(m driver.Machine) (state driver.MachineState, err error) {
	s, err := containers.Inspect(pd.conn, m.Fullname(), nil)
	if err != nil {
		return state, err
	}
	state = driver.MachineState{
		Pid:       s.State.Pid,
		Status:    s.State.Status,
		Running:   s.State.Running,
		StartedAt: s.State.StartedAt,
		ExitCode:  s.State.ExitCode,
	}
	return state, nil
}

func (pd *PodmanDriver) AttachToMachine(m driver.Machine) (err error) {
	exists, err := pd.MachineExists(m)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name)
	}
	opts := new(containers.AttachOptions)
	fmt.Printf("Attaching to %s, Use key sequence <ctrl><p>, <ctrl><q> to detach.\n", m.Name)
	fmt.Printf("You might need to hit <enter> once attached to get a prompt.\n\n")
	err = containers.Attach(pd.conn, m.Fullname(), os.Stdin, os.Stdout, os.Stderr, nil, opts)
	return err
}

func (pd *PodmanDriver) Shell(m driver.Machine, user, workdir string) (err error) {
	exists, err := pd.MachineExists(m)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name)
	}
	ec := new(handlers.ExecCreateConfig)
	ec.Cmd = []string{"/bin/bash"}
	ec.User = user
	ec.WorkingDir = workdir
	ec.AttachStderr = true
	ec.AttachStdin = true
	ec.AttachStdout = true
	ec.Tty = true
	exId, err := containers.ExecCreate(pd.conn, m.Fullname(), ec)
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

func (pd *PodmanDriver) Exec(m driver.Machine, command,
	user string, detach bool, workdir string) (err error) {
	exists, err := pd.MachineExists(m)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name)
	}
	ec := new(handlers.ExecCreateConfig)
	ec.Cmd = strings.Fields(command)
	ec.User = user
	ec.Detach = detach
	ec.WorkingDir = workdir
	if !detach {
		ec.AttachStderr = true
		ec.AttachStdout = true
	}
	exId, err := containers.ExecCreate(pd.conn, m.Fullname(), ec)
	if err != nil {
		return err
	}
	options := new(containers.ExecStartAndAttachOptions)
	if !detach {
		options.WithOutputStream(io.WriteCloser(os.Stdout))
		options.WithAttachOutput(true)
		options.WithErrorStream(io.WriteCloser(os.Stderr))
		options.WithAttachError(true)
	}
	err = containers.ExecStartAndAttach(pd.conn, exId, options)
	return err
}

func (pd *PodmanDriver) GetMachineLogs(m driver.Machine,
	follow bool, tail int) (err error) {
	exists, err := pd.MachineExists(m)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name)
	}
	opts := new(containers.LogOptions)
	opts.WithStdout(true)
	opts.WithStderr(true)
	opts.WithTail(fmt.Sprint(tail))
	opts.WithFollow(follow)
	stdoutChan := make(chan string)
	stderrChan := make(chan string)
	go func() {
		for recv := range stdoutChan {
			fmt.Println(recv)
		}
	}()
	go func() {
		for recv := range stderrChan {
			fmt.Println(recv)
		}
	}()
	err = containers.Logs(pd.conn, m.Fullname(), opts, stdoutChan, stderrChan)
	return err
}

func (pd *PodmanDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", "", namespace, all) // TODO get namespace here
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.conn, opts)
	if err != nil {
		return machines, err
	}
	for _, c := range ctrs {
		name, _, lab := getInfoFromLabels(c.Labels)
		var mNetworks []string
		for _, n := range c.Networks {
			s := strings.Index(n, "netkit_")
			if s == -1 {
				mNetworks = append(mNetworks, n)
			} else {
				s = s + 7 // s should be 0, + 7 accounts for netkit_
				e := strings.Index(n[s:], "_")
				if e == -1 {
					mNetworks = append(mNetworks, n)
				} else {
					mNetworks = append(mNetworks, n[s:s+e])
				}
			}
		}
		machines = append(machines, driver.MachineInfo{
			Name:     name,
			Lab:      lab,
			Image:    c.Image[:strings.IndexByte(c.Image, ':')],
			Networks: mNetworks,
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

func (pd *PodmanDriver) MachineInfo(m driver.Machine) (info driver.MachineInfo, err error) {
	exists, err := pd.MachineExists(m)
	if err != nil {
		return info, err
	} else if !exists {
		return info, driver.ErrNotExists
	}
	inspect, err := containers.Inspect(pd.conn, m.Fullname(), nil)
	if err != nil {
		return info, err
	}
	info.Name = m.Name
	info.Image = inspect.ImageName
	info.State = inspect.State.Status
	return info, err
}

func (pd *PodmanDriver) CopyInFiles(m driver.Machine, hostlab string) error {
	machineDir := filepath.Join(hostlab, m.Name)
	mDirInfo, err := os.Stat(machineDir)
	if os.IsNotExist(err) {
		log.Warnf("Machine directory %s doesn't exist, creating machine %s without mounting custom files.\n", machineDir, m.Name)
		return nil
	} else if err != nil {
		return err
	}
	if !mDirInfo.IsDir() {
		return fmt.Errorf("%s is a file when it should be the machine directory for %s.", machineDir, m.Name)
	}
	opts := new(containers.CopyOptions)
	reader, writer := io.Pipe()
	var copts copier.GetOptions
	go func() {
		defer writer.Close()
		err := copier.Get("/", "", copts, []string{machineDir}, writer)
		if err != nil {
			log.Fatal(err)
		}
	}()
	cp, err := containers.CopyFromArchiveWithOptions(pd.conn, m.Fullname(), "/", reader, opts)
	if err != nil {
		return err
	}
	err = cp()
	if err != nil {
		return err
	}
	return nil
}

func (pd *PodmanDriver) ListAllNamespaces() (namespaces []string, err error) {
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", "", "", true)
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.conn, opts)
	if err != nil {
		return namespaces, err
	}
	for _, c := range ctrs {
		_, ns, _ := getInfoFromLabels(c.Labels)
		found := false
		for _, n := range namespaces {
			if ns == n {
				found = true
			}
		}
		if !found {
			namespaces = append(namespaces, ns)
		}

	}
	return namespaces, err
}

func (pd *PodmanDriver) GetCLICommand() (command *cobra.Command, err error) {
	return new(cobra.Command), nil
}