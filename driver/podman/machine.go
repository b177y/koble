package podman

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	copier "github.com/containers/buildah/copier"
	"github.com/containers/image/v5/manifest"
	"github.com/creasty/defaults"

	"github.com/b177y/netkit/driver"
	"github.com/containers/podman/v3/pkg/api/handlers"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/bindings/images"
	"github.com/containers/podman/v3/pkg/specgen"
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
	return "netkit_" + m.namespace + "_" + m.name
}

func (m *Machine) Exists() (bool, error) {
	return containers.Exists(m.pd.conn, m.Id(), nil)
}

func (m *Machine) Running() bool {
	// TODO add err
	inspect, err := containers.Inspect(m.pd.conn, m.Id(), nil)
	if err != nil {
		return false
	}
	return inspect.State.Running
}

func (m *Machine) State() (state string, err error) {
	inspect, err := containers.Inspect(m.pd.conn, m.Id(), nil)
	if err != nil {
		return "", err
	}
	if inspect.State.Status == "running" {
		hc, err := containers.RunHealthCheck(m.pd.conn, m.Id(), nil)
		if err != nil {
			return "", err
		}
		if hc.Status != "healthy" {
			return "booting", nil
		} else {
			return "running", nil
		}
	} else {
		return inspect.State.Status, nil
	}
}

func (m *Machine) getLabels() map[string]string {
	labels := make(map[string]string)
	labels["netkit"] = "true"
	labels["netkit:name"] = m.Name()
	// if m.Lab != "" {
	// 	labels["netkit:lab"] = m.Lab
	// }
	labels["netkit:namespace"] = m.namespace
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

func (m *Machine) Start(opts *driver.StartOptions) (err error) {
	if opts == nil {
		opts = new(driver.StartOptions)
	}
	if err := defaults.Set(opts); err != nil {
		return err
	}
	if opts.Image == "" {
		opts.Image = m.pd.DefaultImage
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if exists {
		if m.Running() {
			return nil
		} else {
			prev := log.GetLevel()
			log.SetLevel(log.ErrorLevel)
			err = containers.Start(m.pd.conn, m.Id(), nil)
			log.SetLevel(prev)
			return err
		}
	}
	if opts.Image == "" {
		opts.Image = m.pd.DefaultImage
	}
	imExists, err := images.Exists(m.pd.conn, opts.Image, nil)
	if err != nil {
		return err
	}
	if !imExists {
		fmt.Println("Image", opts.Image, "does not already exist, attempting to pull...")
		_, err = images.Pull(m.pd.conn, opts.Image, nil)
		if err != nil {
			return err
		}
	}
	s := specgen.NewSpecGenerator(opts.Image, false)
	s.Name = m.Id()
	s.Hostname = m.Name()
	s.Command = []string{"/sbin/init"}
	s.CapAdd = []string{"NET_ADMIN", "SYS_ADMIN", "CAP_NET_BIND_SERVICE", "CAP_NET_RAW", "CAP_SYS_NICE", "CAP_IPC_LOCK", "CAP_CHOWN"}
	s.NetNS = specgen.Namespace{
		NSMode: specgen.Bridge,
	}
	s.UseImageHosts = true
	s.UseImageResolvConf = true
	for _, n := range opts.Networks {
		net, err := m.pd.Network(n, m.namespace)
		if err != nil {
			return err
		}
		s.CNINetworks = append(s.CNINetworks, net.Id())
	}
	s.ContainerHealthCheckConfig.HealthConfig = &manifest.Schema2HealthConfig{
		Test:    []string{"CMD-SHELL", "test", "$(systemctl show -p ExecMainCode --value netkit-startup-phase2.service)", "-eq", "1"},
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
	createResponse, err := containers.CreateWithSpec(m.pd.conn, s, nil)
	if err != nil {
		return err
	}
	// TODO make m.CopyInFiles
	err = m.CopyInFiles(opts.Hostlab)
	if err != nil {
		return err
	}
	// temporary fix to https://github.com/containers/podman/issues/12204
	prev := log.GetLevel()
	log.SetLevel(log.ErrorLevel)
	err = containers.Start(m.pd.conn, createResponse.ID, nil)
	log.SetLevel(prev)
	return err
}

func (m *Machine) Stop(force bool) error {
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist", m.Name())
	}
	if !m.Running() {
		return fmt.Errorf("Can't stop %s as it isn't running", m.Name())
	}
	if force {
		return containers.Kill(m.pd.conn, m.Id(), nil)
	}
	return containers.Stop(m.pd.conn, m.Id(), nil)
}

func (m *Machine) Remove() error {
	exists, err := m.Exists()
	if err != nil {
		fmt.Println("hmm error checking exists", err)
		return err
	}
	if !exists {
		fmt.Println("doesnt exist so not removing")
		return nil
	}
	fmt.Println("container exists so removing it now")
	return containers.Remove(m.pd.conn, m.Id(), nil)
}

func (m *Machine) Info() (info driver.MachineInfo, err error) {
	s, err := containers.Inspect(m.pd.conn, m.Id(), nil)
	if err != nil {
		return info, err
	}
	info = driver.MachineInfo{
		Name:      m.name,
		Pid:       s.State.Pid,
		Status:    s.State.Status, // TODO make the same as UML
		Running:   s.State.Running,
		StartedAt: s.State.StartedAt,
		ExitCode:  s.State.ExitCode,
		Image:     s.ImageName,
		State:     s.State.Status,
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
	fmt.Printf("Attaching to %s, Use key sequence <ctrl><p>, <ctrl><q> to detach.\n", m.Name())
	fmt.Printf("You might need to hit <enter> once attached to get a prompt.\n\n")
	err = containers.Attach(m.pd.conn, m.Id(), os.Stdin, os.Stdout, os.Stderr, nil, aOpts)
	return err
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
	exId, err := containers.ExecCreate(m.pd.conn, m.Id(), ec)
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
	err = containers.ExecStartAndAttach(m.pd.conn, exId, options)
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
	exId, err := containers.ExecCreate(m.pd.conn, m.Id(), ec)
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
	err = containers.ExecStartAndAttach(m.pd.conn, exId, options)
	return err
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
	err = containers.Logs(m.pd.conn, m.Id(), lOpts, stdoutChan, stderrChan)
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
	cp, err := containers.CopyFromArchiveWithOptions(m.pd.conn, m.Id(), "/", reader, opts)
	if err != nil {
		return err
	}
	err = cp()
	if err != nil {
		return err
	}
	return nil
}

func (m *Machine) WaitUntil(state string, timeout time.Duration) error {
	// TODO make this global method within driver package?
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	for {
		// once condition is met return
		mState, err := m.State()
		if err != nil {
			return err
		}
		if mState == state {
			return nil
		}
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("timed out waiting for %s to be in state %s (currently in state %s): %w",
				m.name, state, mState, err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (m *Machine) Networks() ([]driver.Network, error) {
	return []driver.Network{}, driver.ErrNotImplemented
}
