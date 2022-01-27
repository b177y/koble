package uml

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/b177y/go-uml-utilities/pkg/mconsole"
	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/driver/uml/shim"
	"github.com/b177y/koble/pkg/driver/uml/vecnet"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/creasty/defaults"
	"github.com/docker/docker/pkg/reexec"
	ht "github.com/hpcloud/tail"
	spec "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
)

func init() {
	reexec.Register("umlShim", shim.RunShim)
	if reexec.Init() {
		os.Exit(0)
	}
}

type Machine struct {
	name      string
	namespace string
	ud        *UMLDriver
}

func (m *Machine) Name() string {
	return m.name
}

func (m *Machine) Id() string {
	return fmt.Sprintf("%x",
		md5.Sum([]byte(m.name+"."+m.namespace)))
}

func (m *Machine) Pid() int {
	return processBySubstring("umid="+m.Id(),
		"UMLNAMESPACE="+m.namespace)
}

func (m *Machine) Exists() (bool, error) {
	if m.Pid() > 0 {
		return true, nil
	}
	// check uml rundir for machine
	if _, err := os.Stat(m.mDir()); err == nil {
		return true, nil
	} else if err != nil && !os.IsNotExist(err) {
		return false, err
	}
	return false, nil
}

func (m *Machine) Running() bool {
	return m.Pid() > 0
}

func getKernelCMD(m *Machine, opts driver.MachineConfig, networks []string) (cmd []string, err error) {
	log.Debugf("generating kernel command for %s (namespace %s)", m.Name(), m.namespace)
	cmd = []string{filepath.Join(m.ud.Config.StorageDir, "kernel", m.ud.Config.Kernel)}
	cmd = append(cmd, "name="+m.name, "title="+m.name, "umid="+m.Id())
	cmd = append(cmd, "mem=132M")
	// fsPath := filepath.Join(ud.StorageDir, "images", ud.DefaultImage)
	cmd = append(cmd, fmt.Sprintf("ubd0=%s,%s", m.diskPath(),
		filepath.Join(m.ud.Config.StorageDir, "images", opts.Image)))
	cmd = append(cmd, "root=98:0")
	cmd = append(cmd, "uml_dir="+m.mDir())
	cmd = append(cmd, "con0=fd:0,fd:1", "con1=null")
	cmd = append(cmd, networks...)
	// TODO support any volume (need to modify koble-fs phase1 startup)
	for _, v := range opts.Volumes {
		if v.Destination == "/hosthome" {
			cmd = append(cmd, "hosthome="+v.Source)
		} else if v.Destination == "/hostlab" {
			cmd = append(cmd, "hostlab="+v.Source)
		}
	}
	cmd = append(cmd, "SELINUX_INIT=0")
	cmd = append(cmd, "UMLNAMESPACE="+m.namespace)
	if opts.Lab != "" {
		cmd = append(cmd, "UMLLAB="+opts.Lab)
	}
	if log.GetLevel() <= log.WarnLevel {
		cmd = append(cmd, "quiet")
	}
	return cmd, nil
}

func runInShim(mDir, namespace string, kernelCmd []string) error {
	return vecnet.WithNetNS(namespace, func(ns.NetNS) error {
		c := reexec.Command("umlShim")
		c.Args = append(c.Args, mDir)
		c.Args = append(c.Args, kernelCmd...)
		c.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}
		log.Info("starting shim with kernelCmd: ", kernelCmd)
		return c.Start()
	})
}

func (m *Machine) Start(opts *driver.MachineConfig) (err error) {
	log.WithFields(log.Fields{"options": fmt.Sprintf("%+v", opts)}).Infof(
		"Starting machine %s in namespace %s\n", m.Name(), m.namespace,
	)
	if opts == nil {
		opts = new(driver.MachineConfig)
	}
	if err := defaults.Set(opts); err != nil {
		return err
	}
	if opts.Image == "" {
		opts.Image = m.ud.Config.DefaultImage
	}
	if m.Running() {
		log.WithFields(log.Fields{"machine": m.Name(), "namespace": m.namespace}).
			Debugf("machine already running, not starting")
		return nil
	}
	defer func() {
		if err != nil {
			log.WithFields(log.Fields{"machine": m.Name(), "namespace": m.namespace}).
				Debugf("error in machine start, removing machine: %v\n", err)
			m.Remove()
			os.MkdirAll(m.mDir(), 0744)
			os.WriteFile(filepath.Join(m.mDir(), "state"), []byte("failed"), 0644)
		}
	}()
	nsMdir := filepath.Join(m.ud.Config.RunDir, "ns", m.namespace)
	err = os.MkdirAll(nsMdir, 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	err = os.MkdirAll(m.mDir(), 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	err = os.WriteFile(filepath.Join(m.mDir(), "state"), []byte("starting"), 0644)
	if err != nil {
		return err
	}
	// Remove symlink if it already exists
	if _, err := os.Stat(m.nsDir()); err == nil {
		err = os.Remove(m.nsDir())
		if err != nil {
			return err
		}
	}
	err = os.Symlink(m.mDir(), m.nsDir())
	if err != nil {
		return err
	}
	err = saveInfo(m.mDir(), opts)
	if err != nil {
		return err
	}
	var networks []string
	for i, n := range opts.Networks {
		// setup tap
		ifaceName, err := vecnet.AddHostToNet(m.name, n, m.namespace)
		if err != nil {
			return fmt.Errorf("Could not add machine %s to network %s: %w", m.Name(), n, err)
		}
		// TODO find a way to rename vec devices on boot
		// cmd := fmt.Sprintf("vec%d:transport=tap,ifname=%s", i, ifaceName)
		cmd := fmt.Sprintf("eth%d=tuntap,%s", i, ifaceName)
		// add to networks for cmdline
		networks = append(networks, cmd)
	}
	ifaceName, mgmtIp, err := vecnet.SetupMgmtIface(m.name, m.namespace, filepath.Join(m.mDir(), "slirp.sock"))
	if err != nil {
		return fmt.Errorf("Could not setup management interface: %w", err)
	}
	// TODO autoconf with custom ip
	networks = append(networks, fmt.Sprintf("vec%d:transport=tap,ifname=%s,mac=00:03:B8:FA:CA:DE autoconf_koble0=%s",
		len(networks), ifaceName, mgmtIp))
	if opts.HostHome {
		opts.Volumes = append(opts.Volumes, spec.Mount{
			Source:      os.Getenv("UML_ORIG_HOME"),
			Destination: "/hosthome",
		})
	}
	kernelcmd, err := getKernelCMD(m, *opts, networks)
	if err != nil {
		return fmt.Errorf("could not generate kernel cmd: %w", err)
	}
	// fmt.Println("Got kernelcmd", kernelcmd)
	err = runInShim(m.mDir(), m.namespace, kernelcmd)
	if err != nil {
		return fmt.Errorf("failed to start instance in shim: %w", err)
	}
	return err
}

func (m *Machine) Stop(force bool) (err error) {
	// defer func() {
	// 	if err == nil {
	// 		// TODO remove this once test kernel patch reverted
	// 		os.RemoveAll(filepath.Join(m.mDir(), m.Id()))
	// 	}
	// }()
	exists, err := m.Exists()
	if err != nil {
		return err
	} else if !exists {
		// make force stop immutable (like how `rm -f` doesn't error if file doesn't exist)
		if force {
			return nil
		}
		return fmt.Errorf("can't stop %s as it does not exist", m.name)
	}
	if !m.Running() {
		// make force stop immutable
		if force {
			return nil
		}
		return fmt.Errorf("can't stop %s as it isn't running", m.name)
	}
	umlDir := filepath.Join(m.ud.Config.RunDir, "machine", m.Id(), m.Id())
	if !force {
		_, err = mconsole.CommandWithSock(mconsole.CtrlAltDel(),
			filepath.Join(umlDir, "mconsole"))
		// if socket timeout return nil
		// TODO patch UML kernel to respond before executing cad action
		// if err, ok := err.(net.Error); ok && err.Timeout() {
		// string error comparison is bad practice but above does not work
		// no documentation found for unix socket deadline exceeded errors
		if err.Error() == "read socket timeout" {
			return nil
		}
		return err
	}
	pid := m.Pid()
	// Check if process exists
	killErr := syscall.Kill(-pid, 0)
	if killErr != nil {
		return fmt.Errorf("Could not crash machine %s (%d): %w", m.name, pid, killErr)
	}
	// Send shutdown signal to UML instance
	sig := syscall.SIGKILL
	err = syscall.Kill(-pid, sig)
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		// wait for kill 0 to give err (shows pid no longer running)
		err = syscall.Kill(-pid, 0)
		if err != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err == nil { // kill 0 error == nil means still running
		return fmt.Errorf("Could not kill machine %s (%d)",
			m.name, pid)
	}
	return nil
}

func (m *Machine) Remove() error {
	// TODO WARN on non fatal errors (cannot remove paths etc)
	if m.Running() {
		return errors.New("Machine can't be removed as it's running")
	}
	os.RemoveAll(m.mDir())
	os.RemoveAll(filepath.Join(m.nsDir(), m.name))
	err := vecnet.RemoveSlirp(filepath.Join(m.mDir(), "slirp.sock"))
	if err != nil {
		log.Warnf("Could not remove slirp for machine %s: %w\n",
			m.Name(), err)
	}
	err = vecnet.RemoveMachineNets(m.Name(), m.namespace, true)
	if err != nil {
		log.Warnf("Could not remove networks for machine %s: %w\n",
			m.Name(), err)
	}
	os.Remove(m.diskPath())
	return nil
}

func (m *Machine) Attach(opts *driver.AttachOptions) (err error) {
	if opts == nil {
		opts = new(driver.AttachOptions)
	}
	if !m.Running() {
		return fmt.Errorf("cannot attach to machine %s: not running", m.name)
	}
	if os.Getenv("_KOBLE_IN_TERM") == "" && (log.GetLevel() > log.ErrorLevel) {
		fmt.Printf("Attaching to %s, Use key sequence <ctrl><p>, <ctrl><q> to detach.\n", m.name)
		fmt.Printf("You might need to hit <enter> once attached to get a prompt.\n\n")
	}
	return shim.Attach(filepath.Join(m.mDir(), "attach.sock"))
}

func (m *Machine) Exec(command string,
	opts *driver.ExecOptions) (err error) {
	// TODO check opts and fill with defaults
	if opts == nil {
		opts = new(driver.ExecOptions)
	}
	return vecnet.ExecCommand(m.name, opts.User, command, m.namespace)
}

func (m *Machine) Shell(opts *driver.ShellOptions) (err error) {
	// TODO check opts and fill with defaults
	if opts == nil {
		opts = new(driver.ShellOptions)
	}
	return vecnet.RunShell(m.name, opts.User, m.namespace)
}

func (m *Machine) Logs(opts *driver.LogOptions) (err error) {
	// TODO check opts and fill with defaults
	if opts == nil {
		opts = new(driver.LogOptions)
	}
	fn := filepath.Join(m.ud.Config.RunDir, "machine", m.Id(), "machine.log")
	if opts.Follow {
		t, err := ht.TailFile(fn, ht.Config{Follow: true})
		if err != nil {
			return err
		}
		for line := range t.Lines {
			fmt.Println(line.Text)
		}
	} else {
		f, err := os.Open(fn)
		if err != nil {
			return err
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		startLine := 0
		if opts.Tail >= 0 && opts.Tail < len(lines) {
			startLine = len(lines) - opts.Tail
		}
		for i := startLine; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
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
	return []driver.Network{}, nil
}