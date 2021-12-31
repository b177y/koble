package uml

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/b177y/go-uml-utilities/pkg/mconsole"
	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/uml/shim"
	"github.com/b177y/netkit/driver/uml/vecnet"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/docker/docker/pkg/reexec"
	ht "github.com/hpcloud/tail"
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
		md5.Sum([]byte(m.name+"-"+m.namespace)))
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

func getKernelCMD(m *Machine, opts driver.StartOptions) (cmd []string, err error) {
	cmd = []string{m.ud.Kernel}
	cmd = append(cmd, "name="+m.name, "title="+m.name, "umid="+m.Id())
	cmd = append(cmd, "mem=132M")
	// fsPath := filepath.Join(ud.StorageDir, "images", ud.DefaultImage)
	cmd = append(cmd, fmt.Sprintf("ubd0=%s,%s", m.diskPath(), m.ud.DefaultImage))
	cmd = append(cmd, "root=98:0")
	cmd = append(cmd, "uml_dir="+m.mDir())
	cmd = append(cmd, "con0=fd:0,fd:1", "con1=null")
	var networks []string
	for _, n := range opts.Networks {
		networks = append(networks, n.Name)
	}
	cmd = append(cmd, networks...)
	if opts.HostHome {
		home, err := os.UserHomeDir()
		if err != nil {
			return []string{}, err
		}
		cmd = append(cmd, "hosthome="+home)
	}
	if opts.Hostlab != "" {
		cmd = append(cmd, "hostlab="+opts.Hostlab)
	}
	cmd = append(cmd, "SELINUX_INIT=0")
	cmd = append(cmd, "UMLNAMESPACE="+m.namespace)
	if opts.Lab != "" {
		cmd = append(cmd, "UMLLAB="+opts.Lab)
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
		return c.Start()
	})
}

func (m *Machine) Start(opts *driver.StartOptions) (err error) {
	if opts == nil {
		opts = new(driver.StartOptions)
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if exists {
		if err != nil {
			return fmt.Errorf("could not get machine state: %w", err)
		}
		if m.Running() {
			return nil
		}
	}
	defer func() {
		if err != nil {
			m.Remove()
		}
	}()
	nsMdir := filepath.Join(m.ud.RunDir, "ns", m.namespace)
	err = os.MkdirAll(nsMdir, 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	err = os.MkdirAll(m.mDir(), 0744)
	if err != nil && err != os.ErrExist {
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
	configBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(m.mDir(), "config.json"),
		configBytes, 0644)
	if err != nil {
		return err
	}
	var networks []string
	for i, n := range opts.Networks {
		// setup tap
		ifaceName, err := vecnet.AddHostToNet(m.name, n.Name, m.namespace)
		if err != nil {
			return fmt.Errorf("Could not add machine %s to network %s: %w", m.name, n.Name, err)
		}
		cmd := fmt.Sprintf("vec%d:transport=tap,ifname=%s", i, ifaceName)
		// add to networks for cmdline
		networks = append(networks, cmd)
	}
	ifaceName, err := vecnet.SetupMgmtIface(m.name, m.namespace, filepath.Join(m.mDir(), "slirp.sock"))
	if err != nil {
		return fmt.Errorf("Could not setup management interface: %w", err)
	}
	// TODO autoconf with custom ip
	networks = append(networks, fmt.Sprintf("vec%d:transport=tap,ifname=%s,mac=00:03:B8:FA:CA:DE autoconf_netkit0=10.22.2.110/24",
		len(networks), ifaceName))
	// for _, mnt := range m.Volumes {
	// 	if mnt.Type == "" {
	// 		mnt.Type = "bind"
	// 	}
	// }
	kernelcmd, err := getKernelCMD(m, *opts)
	if err != nil {
		return err
	}
	// fmt.Println("Got kernelcmd", kernelcmd)
	err = runInShim(m.mDir(), m.namespace, kernelcmd)
	if err != nil {
		return err
	}
	return err
}

func (m *Machine) Stop(force bool) error {
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
	umlDir := filepath.Join(m.ud.RunDir, "machine", m.Id(), m.Id())
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
	// get networks for machine
	// for _, n := range m.Networks {
	// 	vecnet.RemoveHostTap(m.Name, n, m.Namespace)
	// }
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
	fmt.Printf("Attaching to %s, Use key sequence <ctrl><p>, <ctrl><q> to detach.\n", m.name)
	fmt.Printf("You might need to hit <enter> once attached to get a prompt.\n\n")
	err = shim.Attach(filepath.Join(m.mDir(), "attach.sock"))
	if err.Error() == "read escape sequence" {
		return nil
	} else {
		return err
	}
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
	fn := filepath.Join(m.ud.RunDir, "machine", m.Id(), "machine.log")
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

func (m *Machine) WaitUntil(state string,
	timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	for {
		// once condition is met return
		if m.State() == state {
			return nil
		}
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("timed out waiting for %s to be in state %s (currently in state %s): %w",
				m.name, state, m.State(), err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (m *Machine) Networks() ([]driver.Network, error) {
	return []driver.Network{}, nil
}
