package uml

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/b177y/koble/pkg/driver"
	"github.com/containers/image/v5/manifest"
	"github.com/containers/podman/v4/pkg/api/handlers"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/bindings/images"
	"github.com/containers/podman/v4/pkg/specgen"
	"github.com/creasty/defaults"
	"github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
)

type Machine struct {
	name      string
	namespace string
	ud        *UMLDriver
	p         driver.Machine
}

func (m *Machine) Name() string {
	return m.p.Name()
}

func (m *Machine) Id() string {
	return m.p.Id()
	// return m.p.Id() + ".uml"
}

func (m *Machine) Exists() (bool, error) {
	return m.p.Exists()
}

func (m *Machine) Running() bool {
	return m.p.Running()
}

func (m *Machine) State() (state driver.MachineState, err error) {
	return m.p.State()
}

func (m *Machine) getLabels(opts driver.MachineConfig) map[string]string {
	labels := make(map[string]string)
	labels["koble"] = "true"
	labels["koble:name"] = m.Name()
	labels["koble:driver"] = "uml"
	labels["koble:namespace"] = m.namespace
	labels["uml:image"] = opts.Image
	return labels
}

func getKernelCMD(m *Machine, opts driver.MachineConfig, networks []string) (cmd []string, err error) {
	log.Debugf("generating kernel command for %s (namespace %s)", m.Name(), m.namespace)
	cmd = []string{"/entrypoint.sh", filepath.Join("/uml/kernel", m.ud.Config.Kernel)}
	cmd = append(cmd, "name="+m.name, "title="+m.name, "umid="+m.Id())
	cmd = append(cmd, "mem=132M")
	// fsPath := filepath.Join(ud.StorageDir, "images", ud.DefaultImage)
	cmd = append(cmd, fmt.Sprintf("ubd0=/overlay.disk,%s",
		filepath.Join("/uml/images", "koble-fs"))) // TODO
	cmd = append(cmd, "root=98:0")
	cmd = append(cmd, "uml_dir=/root") //TODO
	cmd = append(cmd, "con0=fd:0,fd:1", "con1=null")
	cmd = append(cmd, networks...)
	// TODO support any volume (need to modify koble-fs phase1 startup)
	for _, v := range opts.Volumes {
		if v.Destination == "/hosthome" {
			cmd = append(cmd, "hosthome="+v.Source)
		} else if v.Destination == "/hostlab" {
			cmd = append(cmd, "hostlab=/hostlab")
		}
	}
	if opts.Hostlab {
		cmd = append(cmd, "kstart:hostlab=true")
	}
	cmd = append(cmd, "kstart:driver=UML")
	cmd = append(cmd, "SELINUX_INIT=0")
	if log.GetLevel() < log.WarnLevel {
		cmd = append(cmd, "quiet")
	} else if log.GetLevel() < log.InfoLevel {
		cmd = append(cmd, "loglevel=6")
	} else {
		cmd = append(cmd, "loglevel=7")
	}
	return cmd, nil
}

func (m *Machine) Start(opts *driver.MachineConfig) (err error) {
	if opts == nil {
		opts = new(driver.MachineConfig)
	}
	if err := defaults.Set(opts); err != nil {
		return err
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if exists {
		if m.Running() {
			return nil
		} else {
			return containers.Start(m.ud.Podman.Conn, m.Id(), nil)
		}
	}
	imExists, err := images.Exists(m.ud.Podman.Conn, "docker.io/b177y/uml-runner", nil)
	if err != nil {
		return err
	}
	if !imExists {
		fmt.Println("Image docker.io/b177y/uml-runner does not already exist, attempting to pull...")
		_, err = images.Pull(m.ud.Podman.Conn, "docker.io/b177y/uml-runner", nil)
		if err != nil {
			return err
		}
	}
	opts.Image = m.ud.Config.DefaultImage
	s := specgen.NewSpecGenerator("docker.io/b177y/uml-runner", false)
	s.Name = m.Id()
	s.Hostname = m.Name()
	s.Command = []string{"/bin/bash", "-c"}
	// s.CapAdd = []string{"NET_ADMIN", "SYS_ADMIN", "CAP_NET_BIND_SERVICE", "CAP_NET_RAW", "CAP_SYS_NICE", "CAP_IPC_LOCK", "CAP_CHOWN", "CAP_SYS_PTRACE", "CAP_MKNOD"}
	s.Privileged = true // TODO
	if len(opts.Networks) != 0 {
		s.NetNS = specgen.Namespace{
			NSMode: specgen.Bridge,
		}
		s.CNINetworks = make([]string, 0)
		for _, n := range opts.Networks {
			net, err := m.ud.Network(n, m.namespace)
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
	s.Env = make(map[string]string, 0)
	s.Env["TMPDIR"] = "/tmp"
	s.Sysctl = make(map[string]string, 0)
	s.Sysctl["net.ipv4.conf.all.forwarding"] = "1"
	s.ContainerHealthCheckConfig.HealthConfig = &manifest.Schema2HealthConfig{
		Test:    []string{"CMD-SHELL", "test", "-f", "/run/uml/machine.ready"},
		Timeout: 3 * time.Second,
	}
	s.Terminal = true
	s.Labels = m.getLabels(*opts)
	for _, mnt := range opts.Volumes {
		if mnt.Type == "" {
			mnt.Type = "bind"
		}
		s.Mounts = append(s.Mounts, mnt)
	}
	s.Mounts = append(s.Mounts, specs.Mount{
		Source:      m.ud.Config.StorageDir,
		Destination: "/uml",
		Options:     []string{"exec", "ro"},
		Type:        "bind",
	})
	s.Mounts = append(s.Mounts, specs.Mount{
		Destination: "/tmp",
		Options:     []string{"exec", "rw", "nosuid"},
		Type:        "tmpfs",
	})
	var networks []string
	for i := range opts.Networks {
		// cmd := fmt.Sprintf("eth%d=tuntap,nk%d", i, i)
		cmd := fmt.Sprintf("vec%d:transport=raw,ifname=eth%d", i, i)
		networks = append(networks, cmd)
	}
	kernCmd, err := getKernelCMD(m, *opts, networks)
	s.Command = append(s.Command, strings.Join(kernCmd, " "))
	createResponse, err := containers.CreateWithSpec(m.ud.Podman.Conn, s, nil)
	if err != nil {
		return err
	}
	return containers.Start(m.ud.Podman.Conn, createResponse.ID, nil)
}

func (m *Machine) Stop(force bool) error {
	if !force {
		ec := new(handlers.ExecCreateConfig)
		ec.Cmd = []string{"mconsole"}
		ec.Cmd = append(ec.Cmd, filepath.Join("/root", m.Id(), "mconsole"))
		ec.Cmd = append(ec.Cmd, "cad")
		ec.Detach = true
		exId, err := containers.ExecCreate(m.ud.Podman.Conn, m.Id(), ec)
		if err != nil {
			return err
		}
		return containers.ExecStart(m.ud.Podman.Conn, exId, nil)
	}
	return m.p.Stop(true)
}

func (m *Machine) Remove() error {
	return m.p.Remove()
}

func (m *Machine) Info() (info driver.MachineInfo, err error) {
	info, err = m.p.Info()
	if err != nil {
		return info, err
	}
	inspect, err := containers.Inspect(m.ud.Podman.Conn, m.Id(), nil)
	if err != nil {
		return info, err
	}
	if val, ok := inspect.Config.Labels["uml:image"]; ok {
		info.Image = val
	} else {
		return info, fmt.Errorf("label uml:image not added for %s", m.Name())
	}
	return info, err
}

func (m *Machine) Attach(opts *driver.AttachOptions) (err error) {
	return m.p.Attach(opts)
}

func (m *Machine) Shell(opts *driver.ShellOptions) (err error) {
	return driver.ErrNotImplemented
}

func (m *Machine) Exec(command string,
	opts *driver.ExecOptions) (err error) {
	return driver.ErrNotImplemented
}

func (m *Machine) Logs(opts *driver.LogOptions) (err error) {
	return m.p.Logs(opts)
}

func (m *Machine) WaitUntil(timeout time.Duration,
	target, failOn *driver.MachineState) error {
	return m.p.WaitUntil(timeout, target, failOn)
}

func (m *Machine) Networks() ([]driver.Network, error) {
	return []driver.Network{}, driver.ErrNotImplemented
}

func getFilters(machine, namespace, driver string, all bool) map[string][]string {
	filters := make(map[string][]string)
	var labelFilters []string
	labelFilters = append(labelFilters, "koble=true")
	labelFilters = append(labelFilters, "koble:driver=uml")
	if !all {
		labelFilters = append(labelFilters, "koble:namespace="+namespace)
		if machine != "" {
			labelFilters = append(labelFilters, "koble:name="+machine)
		}
	}
	filters["label"] = labelFilters
	return filters
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

func (ud *UMLDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", namespace, "uml", all)
	opts.WithFilters(filters)
	ctrs, err := containers.List(ud.Podman.Conn, opts)
	if err != nil {
		return machines, err
	}
	for _, c := range ctrs {
		name, ns := getInfoFromLabels(c.Labels)
		m, err := ud.Machine(name, ns)
		if err != nil {
			return machines, err
		}
		info, err := m.Info()
		if err != nil {
			return machines, err
		}
		machines = append(machines, info)
	}
	return machines, nil
}

func (ud *UMLDriver) ListAllNamespaces() ([]string, error) {
	return ud.Podman.ListAllNamespaces()
}
