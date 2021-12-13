package uml

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/uml/shim"
	"github.com/docker/docker/pkg/reexec"
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
	ud.Name = "UserMode Linux"
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	ud.Kernel = fmt.Sprintf("%s/netkit-jh/kernel/netkit-kernel", homedir)
	ud.DefaultImage = fmt.Sprintf("%s/netkit-jh/fs/netkit-fs", homedir)
	ud.RunDir = fmt.Sprintf("/run/user/%s/netkit/uml", fmt.Sprint(os.Getuid()))
	ud.StorageDir = fmt.Sprintf("%s/.local/share/netkit/uml", homedir)
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
	return nil
}

func (ud *UMLDriver) MachineExists(m driver.Machine) (exists bool,
	err error) {
	mDir := filepath.Join(ud.RunDir, m.Namespace, m.Name+"-runtime")
	if _, err := os.Stat(mDir); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (ud *UMLDriver) getKernelCMD(m driver.Machine) (cmd []string, err error) {
	cmd = []string{ud.Kernel}
	cmd = append(cmd, "name="+m.Name)
	cmd = append(cmd, "title="+m.Name)
	cmd = append(cmd, "umid="+m.Name)
	cmd = append(cmd, "mem=132M")
	diskPath := filepath.Join(ud.StorageDir, "overlay", m.Namespace, m.Name+".disk")
	// fsPath := filepath.Join(ud.StorageDir, "images", ud.DefaultImage)
	cmd = append(cmd, fmt.Sprintf("ubd0=%s,%s", diskPath, ud.DefaultImage))
	cmd = append(cmd, "root=98:0")
	umlDir := filepath.Join(ud.RunDir, m.Namespace)
	cmd = append(cmd, "uml_dir="+umlDir)
	// TODO add networks
	cmd = append(cmd, "con0=fd:0,fd:1")
	cmd = append(cmd, "con1=null")
	if m.HostHome {
		home, err := os.UserHomeDir()
		if err != nil {
			return []string{}, err
		}
		cmd = append(cmd, "hosthome="+home)
	}
	cmd = append(cmd, "SELINUX_INIT=0")
	return cmd, nil
}

func runInShim(sockPath string, kernelCmd []string) error {
	_ = shim.IMPORT
	c := reexec.Command("umlShim")
	c.Args = append(c.Args, sockPath)
	c.Args = append(c.Args, kernelCmd...)
	c.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	// c.Stdout = os.Stdout
	// c.Stderr = os.Stderr
	return c.Start()
}

func (ud *UMLDriver) StartMachine(m driver.Machine) (err error) {
	kernelcmd, err := ud.getKernelCMD(m)
	if err != nil {
		return err
	}
	err = runInShim(filepath.Join(ud.RunDir, m.Namespace, m.Name+"-runtime"), kernelcmd)
	if err != nil {
		return err
	}
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
		} else {
			// err = containers.Start(ud.conn, m.Fullname(), nil)
			_, err := ud.getKernelCMD(m)
			if err != nil {
				return err
			}
			// fmt.Println(kernelcmd)
			return err
		}
	}
	if err != nil {
		return err
	}
	for _, n := range m.Networks {
		net := driver.Network{
			Name: n,
			Lab:  m.Lab,
		}
		fmt.Println("use", net)
	}
	for _, mnt := range m.Volumes {
		if mnt.Type == "" {
			mnt.Type = "bind"
		}
	}
	// err = containers.Start(ud.conn, createResponse.ID, nil)
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
	// err = containers.Stop(ud.conn, m.Fullname(), nil)
	return err
}

func (ud *UMLDriver) RemoveMachine(m driver.Machine) error {
	// err := containers.Remove(ud.conn, m.Fullname(), nil)
	state, err := ud.GetMachineState(m)
	if err != nil {
		return err
	}
	if state.Running {
		return fmt.Errorf("Cannot remove machine %s as it's still running", m.Name)
	}
	return nil
}

func (ud *UMLDriver) GetMachineState(m driver.Machine) (state driver.MachineState, err error) {
	mDir := filepath.Join(ud.RunDir, m.Namespace, m.Name+"-runtime")
	stateFile := filepath.Join(mDir, "state")
	p, err := os.ReadFile(stateFile)
	if err != nil {
		return state, err
	}
	state.Running = false
	state.Status = string(p)
	if string(p) == "running" {
		state.Running = true
	} else if string(p) == "exitted" {
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
	err = shim.Attach(filepath.Join(ud.RunDir, m.Namespace, m.Name+"-runtime", "attach.sock"))
	if err.Error() == "read escape sequence" {
		return nil
	} else {
		return err
	}
}

func (ud *UMLDriver) MachineExecShell(m driver.Machine, command,
	user string, detach bool, workdir string) (err error) {
	return driver.ErrNotImplemented
}

func (ud *UMLDriver) GetMachineLogs(m driver.Machine,
	stdoutChan, stderrChan chan string,
	follow bool, tail int) (err error) {
	return driver.ErrNotImplemented
}

func (ud *UMLDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	// TODO if all, look in all subdirs
	var machines []driver.MachineInfo
	dir := filepath.Join(ud.RunDir, namespace)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return machines, err
	}
	for _, e := range entries {
		n := e.Name()
		if strings.HasSuffix(n, "-runtime") {
			info, err := ud.MachineInfo(driver.Machine{
				Name:      strings.TrimSuffix(n, "-runtime"),
				Namespace: namespace,
			})
			if err != nil && err != driver.ErrNotExists {
				return machines, err
			}
			machines = append(machines, info)
		}
	}
	// ctrs, err := containers.List(ud.conn, opts)
	// if err != nil {
	// 	return machines, err
	// }
	return machines, nil
}

func (ud *UMLDriver) MachineInfo(m driver.Machine) (info driver.MachineInfo, err error) {
	info.Name = m.Name
	info.Lab = m.Lab
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
	return info, nil
}
