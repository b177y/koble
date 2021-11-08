package podman

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	copier "github.com/containers/buildah/copier"

	"github.com/b177y/netkit/driver"
	"github.com/containers/podman/v3/pkg/api/handlers"
	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/specgen"
	log "github.com/sirupsen/logrus"
)

type PodmanDriver struct {
	conn         context.Context
	Name         string
	DefaultImage string
	URI          string
}

type PDConf struct {
	URI string
}

func (pd *PodmanDriver) GetDefaultImage() string {
	return pd.DefaultImage
}

func (pd *PodmanDriver) SetupDriver(conf map[string]interface{}) (err error) {
	pd.Name = "Podman"
	pd.URI = fmt.Sprintf("unix://run/user/%s/podman/podman.sock",
		fmt.Sprint(os.Getuid()))
	pd.DefaultImage = "localhost/netkit-deb-test"
	// override uri with config option
	if val, ok := conf["uri"]; ok {
		if str, ok := val.(string); ok {
			pd.URI = str
		} else {
			return fmt.Errorf("Driver 'uri' in config must be a string.")
		}
	}
	if val, ok := conf["default_image"]; ok {
		if str, ok := val.(string); ok {
			pd.DefaultImage = str
		} else {
			return fmt.Errorf("Driver 'default_image' in config must be a string.")
		}
	}
	log.Debug("Attempting to connect to podman socket.")
	pd.conn, err = bindings.NewConnection(context.Background(), pd.URI)
	if err != nil {
		return driver.NewDriverError(err, pd.Name, "SetupDriver")
	}
	return nil
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

func getInfoFromLabels(labels map[string]string) (name, lab string) {
	if val, ok := labels["netkit:name"]; ok {
		name = val
	}
	if val, ok := labels["netkit:nolab"]; ok && val == "true" {
		lab = ""
	} else if val, ok := labels["netkit:lab"]; ok {
		lab = val
	}
	return name, lab
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
	imExists, err := images.Exists(pd.conn, m.Image, nil)
	if err != nil {
		return driver.NewDriverError(err, pd.Name, "StartMachine")
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
	for _, n := range m.Networks {
		net := driver.Network{
			Name: n,
			Lab:  m.Lab,
		}
		s.CNINetworks = append(s.CNINetworks, net.Fullname())
	}
	s.Terminal = true
	s.Labels = getLabels(m.Name, m.Lab)
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

func (pd *PodmanDriver) MachineExecShell(m driver.Machine, command,
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

func (pd *PodmanDriver) GetMachineLogs(m driver.Machine,
	stdoutChan, stderrChan chan string,
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
	err = containers.Logs(pd.conn, m.Fullname(), opts, stdoutChan, stderrChan)
	return err
}

func (pd *PodmanDriver) ListMachines(lab string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", lab, all)
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.conn, opts)
	if err != nil {
		return machines, err
	}
	for _, c := range ctrs {
		name, lab := getInfoFromLabels(c.Labels)
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
			Image:    c.Image,
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
