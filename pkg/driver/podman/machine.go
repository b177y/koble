package podman

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	copier "github.com/containers/buildah/copier"
	"github.com/containers/image/v5/manifest"
	"github.com/creasty/defaults"
	"github.com/cri-o/ocicni/pkg/ocicni"
	units "github.com/docker/go-units"
	"github.com/opencontainers/runtime-spec/specs-go"

	"github.com/b177y/koble/pkg/driver"
	"github.com/containers/podman/v4/pkg/api/handlers"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/bindings/images"
	"github.com/containers/podman/v4/pkg/bindings/network"
	"github.com/containers/podman/v4/pkg/specgen"
	log "github.com/sirupsen/logrus"
)

type Machine struct {
	name      string
	namespace string
	pd        *PodmanDriver
}

func (m *Machine) Name() string {
	return m.name
}

func (m *Machine) Id() string {
	return "koble." + m.namespace + "." + m.name + "." + m.pd.DriverName
}

func (m *Machine) Exists() (bool, error) {
	return containers.Exists(m.pd.Conn, m.Id(), nil)
}

func (m *Machine) Running() bool {
	// TODO add err
	inspect, err := containers.Inspect(m.pd.Conn, m.Id(), nil)
	if err != nil {
		return false
	}
	return inspect.State.Running
}

func (m *Machine) State() (state driver.MachineState, err error) {
	exists, err := m.Exists()
	if err != nil {
		return state, err
	} else if !exists {
		return driver.MachineState{Exists: false}, driver.ErrNotExists
	}
	state.Exists = true
	inspect, err := containers.Inspect(m.pd.Conn, m.Id(), nil)
	if err != nil {
		return state, err
	}
	if inspect.State.Status == "running" {
		hc, err := containers.RunHealthCheck(m.pd.Conn, m.Id(), nil)
		if err != nil {
			return state, err
		}
		var s string
		if hc.Status == "healthy" {
			s = "running"
		} else {
			s = "booting"
		}
		state.State = &s
	} else {
		state.State = &inspect.State.Status
	}
	state.ExitCode = &inspect.State.ExitCode
	state.Running = &inspect.State.Running
	return state, nil
}

func (m *Machine) getLabels() map[string]string {
	labels := make(map[string]string)
	labels["koble"] = "true"
	labels["koble:name"] = m.Name()
	labels["koble:driver"] = "podman"
	labels["koble:namespace"] = m.namespace
	return labels
}

func getInfoFromLabels(labels map[string]string) (name, namespace string) {
	if val, ok := labels["koble:name"]; ok {
		name = val
	}
	if val, ok := labels["koble:namespace"]; ok {
		namespace = val
	}
	return name, namespace
}

func getFilters(machine, namespace, driver string, all bool) map[string][]string {
	filters := make(map[string][]string)
	var labelFilters []string
	labelFilters = append(labelFilters, "koble=true")
	labelFilters = append(labelFilters, "koble:driver="+driver)
	if !all {
		labelFilters = append(labelFilters, "koble:namespace="+namespace)
		if machine != "" {
			labelFilters = append(labelFilters, "koble:name="+machine)
		}
	}
	filters["label"] = labelFilters
	return filters
}

func (m *Machine) Start(opts *driver.MachineConfig) (err error) {
	if opts == nil {
		opts = new(driver.MachineConfig)
	}
	if err := defaults.Set(opts); err != nil {
		return err
	}
	if opts.Image == "" {
		opts.Image = m.pd.Config.DefaultImage
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if exists {
		if m.Running() {
			return nil
		} else {
			return containers.Start(m.pd.Conn, m.Id(), nil)
		}
	}
	if opts.Image == "" {
		opts.Image = m.pd.Config.DefaultImage
	}
	imExists, err := images.Exists(m.pd.Conn, opts.Image, nil)
	if err != nil {
		return err
	}
	if !imExists {
		fmt.Println("Image", opts.Image, "does not already exist, attempting to pull...")
		_, err = images.Pull(m.pd.Conn, opts.Image, nil)
		if err != nil {
			return err
		}
	}
	s := specgen.NewSpecGenerator(opts.Image, false)
	s.Name = m.Id()
	s.Hostname = m.Name()
	s.Command = []string{"/sbin/init"}
	s.CapAdd = []string{"NET_ADMIN", "SYS_ADMIN", "CAP_NET_BIND_SERVICE", "CAP_NET_RAW", "CAP_SYS_NICE", "CAP_IPC_LOCK", "CAP_CHOWN", "CAP_MKNOD"}
	if len(opts.Networks) != 0 {
		s.NetNS = specgen.Namespace{
			NSMode: specgen.Bridge,
		}
		s.CNINetworks = make([]string, 0)
		for _, n := range opts.Networks {
			net, err := m.pd.Network(n, m.namespace)
			if err != nil {
				return err
			}
			s.CNINetworks = append(s.CNINetworks, net.Id())
		}
	} else {
		s.NetNS = specgen.Namespace{
			NSMode: specgen.NoNetwork,
		}
	}
	s.UseImageHosts = true
	s.Sysctl = make(map[string]string, 0)
	s.Sysctl["net.ipv4.conf.all.forwarding"] = "1"
	s.UseImageResolvConf = true
	s.ContainerHealthCheckConfig.HealthConfig = &manifest.Schema2HealthConfig{
		Test:    []string{"CMD-SHELL", "test", "$(systemctl show -p ExecMainCode --value koble-startup-phase2.service)", "-eq", "1"},
		Timeout: 3 * time.Second,
	}
	s.Terminal = true
	s.Labels = m.getLabels()
	for _, mnt := range opts.Volumes {
		if mnt.Type == "" {
			mnt.Type = "bind"
		}
		s.Mounts = append(s.Mounts, mnt)
	}
	s.Env = make(map[string]string, 0)
	s.Env["kstart-driver"] = "podman"
	s.Env["kstart-quiet"] = strconv.FormatBool(log.GetLevel() <= log.WarnLevel)
	limit, err := units.RAMInBytes("132M") // TODO allow user set mem
	if err != nil {
		return fmt.Errorf("invalid memory format %s: %w", "132M", err)
	}
	s.ResourceLimits = &specs.LinuxResources{
		Memory: &specs.LinuxMemory{Limit: &limit},
	}
	createResponse, err := containers.CreateWithSpec(m.pd.Conn, s, nil)
	if err != nil {
		return err
	}
	// TODO make m.CopyInFiles
	// err = m.CopyInFiles(opts.Hostlab)
	// if err != nil {
	// 	return err
	// }
	return containers.Start(m.pd.Conn, createResponse.ID, nil)
}

func (m *Machine) Stop(force bool) error {
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		// make force stop immutable (like how `rm -f` doesn't error if file doesn't exist)
		if force {
			return nil
		}
		return fmt.Errorf("Machine %s does not exist", m.Name())
	}
	if !m.Running() {
		// make force stop immutable
		if force {
			return nil
		}
		return fmt.Errorf("Can't stop %s as it isn't running", m.Name())
	}
	if force {
		return containers.Kill(m.pd.Conn, m.Id(), nil)
	}
	return containers.Stop(m.pd.Conn, m.Id(), nil)
}

func (m *Machine) Remove() error {
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	_, err = containers.Remove(m.pd.Conn, m.Id(), nil) // throw away RmReport
	if err != nil {
		return fmt.Errorf("could not remove container: %w", err)
	}
	opts := network.PruneOptions{}
	opts.Filters = make(map[string][]string)
	opts.Filters["label"] = []string{"koble:driver=" + m.pd.DriverName}
	_, err = network.Prune(m.pd.Conn, nil)
	if err != nil {
		if strings.Contains(err.Error(), "no such container") {
			log.Tracef("error pruning networks: %w", err)
			return nil
		} else {
			return fmt.Errorf("could not prune networks: %w", err)
		}
	}
	return nil
}

func (m *Machine) Info() (info driver.MachineInfo, err error) {
	s, err := containers.Inspect(m.pd.Conn, m.Id(), nil)
	if err != nil {
		return info, err
	}
	var networks []string
	for key := range s.NetworkSettings.Networks {
		parts := strings.Split(key, ".")
		if len(parts) != 4 {
			return info, fmt.Errorf("network (%s) name format incorrect", key)
		}
		networks = append(networks, parts[2])
	}
	// TODO get healthcheck
	info = driver.MachineInfo{
		Name:      m.name,
		Namespace: m.namespace,
		Pid:       s.State.Pid,
		Status:    s.State.Status, // TODO make the same as UML
		Running:   s.State.Running,
		CreatedAt: s.Created,
		StartedAt: s.State.StartedAt,
		ExitCode:  s.State.ExitCode,
		Image:     s.ImageName,
		State:     s.State.Status,
		Networks:  networks,
		Ports:     []ocicni.PortMapping{},
		Mounts:    []string{},
	}
	return info, nil
}

func (m *Machine) Attach(opts *driver.AttachOptions) (err error) {
	if opts == nil {
		opts = new(driver.AttachOptions)
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name())
	}
	aOpts := new(containers.AttachOptions)
	if os.Getenv("_KOBLE_IN_TERM") == "" && (log.GetLevel() > log.ErrorLevel) {
		fmt.Printf("Attaching to %s, Use key sequence <ctrl><p>, <ctrl><q> to detach.\n", m.Name())
		fmt.Printf("You might need to hit <enter> once attached to get a prompt.\n\n")
	}
	return containers.Attach(m.pd.Conn, m.Id(),
		os.Stdin, os.Stdout, os.Stderr, nil, aOpts)
}

func (m *Machine) Shell(opts *driver.ShellOptions) (err error) {
	if opts == nil {
		opts = new(driver.ShellOptions)
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name())
	}
	ec := new(handlers.ExecCreateConfig)
	ec.Cmd = []string{"/bin/bash"}
	ec.User = opts.User
	ec.WorkingDir = opts.Workdir
	ec.AttachStderr = true
	ec.AttachStdin = true
	ec.AttachStdout = true
	ec.Tty = true
	exId, err := containers.ExecCreate(m.pd.Conn, m.Id(), ec)
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
	err = containers.ExecStartAndAttach(m.pd.Conn, exId, options)
	return err
}

func (m *Machine) Exec(command string,
	opts *driver.ExecOptions) (err error) {
	if opts == nil {
		opts = new(driver.ExecOptions)
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name())
	}
	ec := new(handlers.ExecCreateConfig)
	ec.Cmd = strings.Fields(command)
	ec.User = opts.User
	ec.Detach = opts.Detach
	ec.WorkingDir = opts.Workdir
	if !ec.Detach {
		ec.AttachStderr = true
		ec.AttachStdout = true
	}
	exId, err := containers.ExecCreate(m.pd.Conn, m.Id(), ec)
	if err != nil {
		return err
	}
	options := new(containers.ExecStartAndAttachOptions)
	if !ec.Detach {
		options.WithOutputStream(io.WriteCloser(os.Stdout))
		options.WithAttachOutput(true)
		options.WithErrorStream(io.WriteCloser(os.Stderr))
		options.WithAttachError(true)
	}
	return containers.ExecStartAndAttach(m.pd.Conn, exId, options)
}

func (m *Machine) Logs(opts *driver.LogOptions) (err error) {
	if opts == nil {
		opts = new(driver.LogOptions)
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name())
	}
	lOpts := new(containers.LogOptions)
	lOpts.WithStdout(true)
	lOpts.WithStderr(true)
	lOpts.WithTail(fmt.Sprint(opts.Tail))
	lOpts.WithFollow(opts.Follow)
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
	err = containers.Logs(m.pd.Conn, m.Id(), lOpts, stdoutChan, stderrChan)
	return err
}

func (m *Machine) CopyInFiles(hostlab string) error {
	machineDir := filepath.Join(hostlab, m.Name())
	mDirInfo, err := os.Stat(machineDir)
	if os.IsNotExist(err) {
		log.Warnf("Machine directory %s doesn't exist, creating machine %s without mounting custom files.\n", machineDir, m.Name())
		return nil
	} else if err != nil {
		return err
	}
	if !mDirInfo.IsDir() {
		return fmt.Errorf("%s is a file when it should be the machine directory for %s.", machineDir, m.Name())
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
	cp, err := containers.CopyFromArchiveWithOptions(m.pd.Conn, m.Id(), "/", reader, opts)
	if err != nil {
		return err
	}
	err = cp()
	if err != nil {
		return err
	}
	return nil
}

func (m *Machine) WaitUntil(timeout time.Duration,
	target, failOn *driver.MachineState) error {
	log.WithFields(log.Fields{
		"target": fmt.Sprintf("%+v", target),
		"failon": fmt.Sprintf("%+v", failOn),
	}).Infof(
		"WaitUntil for machine %s in namespace %s\n", m.Name(), m.namespace,
	)
	return driver.WaitUntil(m, timeout, target, failOn)
}

func (m *Machine) Networks() ([]driver.Network, error) {
	return []driver.Network{}, driver.ErrNotImplemented
}
