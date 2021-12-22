package uml

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/uml/shim"
	"github.com/b177y/netkit/driver/uml/vecnet"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/docker/docker/pkg/reexec"
	ht "github.com/hpcloud/tail"
)

type UMLDriver struct {
	Name         string
	DefaultImage string
	Kernel       string
	RunDir       string
	StorageDir   string
}

func (ud *UMLDriver) GetDefaultImage() string {
	return ud.DefaultImage
}

func (ud *UMLDriver) SetupDriver(conf map[string]interface{}) (err error) {
	err = vecnet.CreateAndEnterUserNS("netkit")
	if err != nil {
		return err
	}
	ud.Name = "UserMode Linux"
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	ud.Kernel = fmt.Sprintf("%s/netkit-jh/kernel/netkit-kernel", homedir)
	ud.DefaultImage = fmt.Sprintf("%s/netkit-jh/fs/netkit-fs", homedir)
	ud.RunDir = fmt.Sprintf("/run/user/%s/uml", os.Getenv("UML_ORIG_UID"))
	ud.StorageDir = fmt.Sprintf("%s/.local/share/uml", homedir)
	// override kernel with config option
	if val, ok := conf["kernel"]; ok {
		if str, ok := val.(string); ok {
			ud.Kernel = str
		} else {
			return fmt.Errorf("Driver 'kernel' in config must be a string.")
		}
	}
	if val, ok := conf["default_image"]; ok {
		if str, ok := val.(string); ok {
			ud.DefaultImage = str
		} else {
			return fmt.Errorf("Driver 'default_image' in config must be a string.")
		}
	}
	if val, ok := conf["run_dir"]; ok {
		if str, ok := val.(string); ok {
			ud.RunDir = str
		} else {
			return fmt.Errorf("Driver 'run_dir' in config must be a string.")
		}
	}
	if val, ok := conf["storage_dir"]; ok {
		if str, ok := val.(string); ok {
			ud.StorageDir = str
		} else {
			return fmt.Errorf("Driver 'storage_dir' in config must be a string.")
		}
	}
	err = os.MkdirAll(filepath.Join(ud.StorageDir, "overlay"), 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	err = os.MkdirAll(filepath.Join(ud.RunDir, "ns", "GLOBAL"), 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	return nil
}

func (ud *UMLDriver) MachineExists(m driver.Machine) (exists bool,
	err error) {

	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	mDir := filepath.Join(ud.RunDir, "machine", mHash)
	if _, err := os.Stat(mDir); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (ud *UMLDriver) getKernelCMD(m driver.Machine, networks []string) (cmd []string, err error) {
	cmd = []string{ud.Kernel}
	cmd = append(cmd, "name="+m.Name)
	cmd = append(cmd, "title="+m.Name)
	cmd = append(cmd, "umid="+m.Name)
	cmd = append(cmd, "mem=132M")
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	diskPath := filepath.Join(ud.StorageDir, "overlay", mHash+".disk")
	// fsPath := filepath.Join(ud.StorageDir, "images", ud.DefaultImage)
	cmd = append(cmd, fmt.Sprintf("ubd0=%s,%s", diskPath, ud.DefaultImage))
	cmd = append(cmd, "root=98:0")
	umlDir := filepath.Join(ud.RunDir, "machine", mHash)
	cmd = append(cmd, "uml_dir="+umlDir)
	cmd = append(cmd, "con0=fd:0,fd:1")
	cmd = append(cmd, "con1=null")
	cmd = append(cmd, networks...)
	if m.HostHome {
		home, err := os.UserHomeDir()
		if err != nil {
			return []string{}, err
		}
		cmd = append(cmd, "hosthome="+home)
	}
	if m.Hostlab != "" {
		cmd = append(cmd, "hostlab="+m.Hostlab)
	}
	cmd = append(cmd, "net.ifnames=0")
	cmd = append(cmd, "SELINUX_INIT=0")
	return cmd, nil
}

func runInShim(sockPath, namespace string, kernelCmd []string) error {
	return vecnet.WithNetNS(namespace, func(ns.NetNS) error {
		_ = shim.IMPORT
		c := reexec.Command("umlShim")
		c.Args = append(c.Args, sockPath)
		c.Args = append(c.Args, kernelCmd...)
		c.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}
		return c.Start()
	})
}

func (ud *UMLDriver) StartMachine(m driver.Machine) (err error) {
	exists, err := ud.MachineExists(m)
	if err != nil {
		return err
	}
	if exists {
		state, err := ud.GetMachineState(m)
		if err != nil {
			return err
		}
		if state.Running {
			return nil
		}
	}
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	nsMdir := filepath.Join(ud.RunDir, "ns", m.Namespace)
	mDir := filepath.Join(ud.RunDir, "machine", mHash)
	err = os.MkdirAll(nsMdir, 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	err = os.MkdirAll(mDir, 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	// Remove symlink if it already exists
	if _, err := os.Stat(filepath.Join(nsMdir, m.Name)); err == nil {
		err = os.Remove(filepath.Join(nsMdir, m.Name))
		if err != nil {
			return err
		}
	}
	err = os.Symlink(mDir, filepath.Join(nsMdir, m.Name))
	if err != nil {
		return err
	}
	configBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(mDir, "config.json"), configBytes, 0644)
	if err != nil {
		return err
	}
	var networks []string
	for i, n := range m.Networks {
		// setup tap
		ifaceName, err := vecnet.AddHostToNet(m.Name, n, m.Namespace)
		if err != nil {
			return fmt.Errorf("Could not add machine %s to network %s: %w", m.Name, n, err)
		}
		cmd := fmt.Sprintf("vec%d:transport=tap,ifname=%s", i, ifaceName)
		// add to networks for cmdline
		networks = append(networks, cmd)
	}
	ifaceName, err := vecnet.SetupMgmtIface(m.Name, m.Namespace, filepath.Join(mDir, "slirp.sock"))
	if err != nil {
		return fmt.Errorf("Could not setup management interface: %w", err)
	}
	networks = append(networks, fmt.Sprintf("vec%d:transport=tap,ifname=%s,mac=00:03:B8:FA:CA:DE",
		len(networks), ifaceName))
	// for _, mnt := range m.Volumes {
	// 	if mnt.Type == "" {
	// 		mnt.Type = "bind"
	// 	}
	// }
	kernelcmd, err := ud.getKernelCMD(m, networks)
	if err != nil {
		return err
	}
	// fmt.Println("Got kernelcmd", kernelcmd)
	err = runInShim(mDir, m.Namespace, kernelcmd)
	if err != nil {
		return err
	}
	return err
}

func (ud *UMLDriver) HaltMachine(m driver.Machine, force bool) error {
	exists, err := ud.MachineExists(m)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist", m.Name)
	}
	state, err := ud.GetMachineState(m)
	if err != nil {
		return err
	}
	if !state.Running {
		return fmt.Errorf("Can't stop %s as it isn't running", m.Name)
	}
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	pidFile := filepath.Join(ud.RunDir, "machine", mHash, m.Name, "pid")
	pidBytes, err := os.ReadFile(pidFile)
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(strings.TrimSuffix(string(pidBytes), "\n"))
	if err != nil {
		return err
	}
	// Check if process exists
	killErr := syscall.Kill(pid, 0)
	if killErr == nil {
		// Send shutdown signal to UML instance
		sig := syscall.SIGTERM
		if force {
			sig = syscall.SIGKILL
		}
		err = syscall.Kill(pid, sig)
		if err != nil {
			return err
		}
		for i := 0; i < 10; i++ {
			// wait for kill 0 to give err (shows pid no longer running)
			err = syscall.Kill(pid, 0)
			if err != nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		if err == nil {
			return fmt.Errorf("Could not kill machine %s (%d): %w",
				m.Name, pid, err)
		}
	} else {
		err = os.WriteFile(filepath.Join(ud.RunDir, "machine", mHash, "state"), []byte("killed"), 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ud *UMLDriver) RemoveMachine(m driver.Machine) error {
	// TODO return non fatal errors?
	state, _ := ud.GetMachineState(m)
	if state.Running {
		return errors.New("Machine can't be removed as it's running")
	}
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	nsMdir := filepath.Join(ud.RunDir, "ns", m.Namespace, m.Name)
	mDir := filepath.Join(ud.RunDir, "machine", mHash)
	os.RemoveAll(mDir)
	os.RemoveAll(nsMdir)
	for _, n := range m.Networks {
		vecnet.RemoveHostTap(m.Name, n, m.Namespace)
	}
	os.Remove(filepath.Join(ud.StorageDir, "overlay", mHash+".disk"))
	return nil
}

func (ud *UMLDriver) GetMachineState(m driver.Machine) (state driver.MachineState, err error) {
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	mDir := filepath.Join(ud.RunDir, "machine", mHash)
	stateFile := filepath.Join(mDir, "state")
	p, err := os.ReadFile(stateFile)
	if err != nil {
		return state, err
	}
	state.Running = false
	state.Status = string(p)
	if string(p) == "running" {
		state.Running = true
	} else if string(p) == "exited" {
		ecFile := filepath.Join(mDir, "exitcode")
		p, err := os.ReadFile(ecFile)
		if err == nil {
			ec, err := strconv.ParseInt(string(p), 10, 32)
			if err == nil {
				state.ExitCode = int32(ec)
			}
		}
	}
	// TODO use shirou/gopsutil to get process create_time and work out uptime
	return state, nil
}

func (ud *UMLDriver) AttachToMachine(m driver.Machine) (err error) {
	exists, err := ud.MachineExists(m)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Machine %s does not exist.", m.Name)
	}
	fmt.Printf("Attaching to %s, Use key sequence <ctrl><p>, <ctrl><q> to detach.\n", m.Name)
	fmt.Printf("You might need to hit <enter> once attached to get a prompt.\n\n")
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	err = shim.Attach(filepath.Join(ud.RunDir, "machine", mHash, "attach.sock"))
	if err.Error() == "read escape sequence" {
		return nil
	} else {
		return err
	}
}

func (ud *UMLDriver) Exec(m driver.Machine, command,
	user string, detach bool, workdir string) (err error) {
	return vecnet.ExecCommand(m.Name, user, command, m.Namespace)
}

func (ud *UMLDriver) Shell(m driver.Machine, user, workdir string) (err error) {
	return vecnet.RunShell(m.Name, user, m.Namespace)
}

func (ud *UMLDriver) GetMachineLogs(m driver.Machine,
	follow bool, tail int) (err error) {
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	fn := filepath.Join(ud.RunDir, "machine", mHash, "machine.log")
	if follow {
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
		if tail >= 0 && tail < len(lines) {
			startLine = len(lines) - tail
		}
		for i := startLine; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	}
	return nil
}

func (ud *UMLDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	// TODO if all, look in all subdirs
	var machines []driver.MachineInfo
	if all {
		namespaces, err := os.ReadDir(filepath.Join(ud.RunDir, "ns"))
		if err != nil {
			return machines, err
		}
		for _, n := range namespaces {
			namespaceMachines, err := ud.ListMachines(n.Name(), false)
			if err != nil {
				return machines, err
			}
			machines = append(machines, namespaceMachines...)
		}
	} else {
		dir := filepath.Join(ud.RunDir, "ns", namespace)
		entries, err := os.ReadDir(dir)
		if err != nil {
			return machines, err
		}
		for _, e := range entries {
			// TODO (temp fix)
			if e.Name() != "netns.bind" {
				info, err := ud.MachineInfo(driver.Machine{
					Name:      e.Name(),
					Namespace: namespace,
				})
				if err != nil && err != driver.ErrNotExists {
					return machines, err
				}
				machines = append(machines, info)
			}
		}
	}
	return machines, nil
}

func (ud *UMLDriver) MachineInfo(m driver.Machine) (info driver.MachineInfo, err error) {
	info.Name = m.Name
	info.Namespace = m.Namespace
	exists, err := ud.MachineExists(m)
	if err != nil {
		return info, err
	} else if !exists {
		return info, driver.ErrNotExists
	}
	state, err := ud.GetMachineState(m)
	if err != nil {
		return info, err
	}
	info.State = state.Status
	info.ExitCode = state.ExitCode
	mHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(m.Name+"-"+m.Namespace)))
	mDir := filepath.Join(ud.RunDir, "machine", mHash)
	content, err := ioutil.ReadFile(filepath.Join(mDir, "config.json"))
	var mConfig driver.Machine
	err = json.Unmarshal(content, &mConfig)
	if err != nil {
		return info, err
	}
	info.Networks = mConfig.Networks
	info.Image = mConfig.Image
	info.Lab = mConfig.Lab
	return info, nil
}

func (ud *UMLDriver) ListAllNamespaces() (namespaces []string, err error) {
	nsEntries, err := os.ReadDir(filepath.Join(ud.RunDir, "ns"))
	for _, n := range nsEntries {
		namespaces = append(namespaces, n.Name())
	}
	return namespaces, nil
}
