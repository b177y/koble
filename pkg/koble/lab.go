package koble

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/b177y/koble/driver"
	"github.com/b177y/koble/pkg/output"
	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"

	log "github.com/sirupsen/logrus"
)

var blue = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
var magBold = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
var mag = color.New(color.FgHiMagenta).SprintFunc()
var MAXPRINTWIDTH = 100

type InitOpts struct {
	Name        string
	Description string
	Authors     []string
	Emails      []string
	Webs        []string
}

func (nk *Koble) InitLab(options InitOpts) error {
	if nk.LabRoot == nk.InitialWorkDir {
		fmt.Println("nk labroot", nk.LabRoot, nk.InitialWorkDir)
		return fmt.Errorf("lab.yml already exists in this directory.")
	} else if nk.LabRoot != "" {
		log.Warnf("There is already a lab at %s, creating a new lab in %s\n", nk.LabRoot, nk.InitialWorkDir)
		err := os.Chdir(nk.InitialWorkDir)
		if err != nil {
			return err
		}
	}
	if options.Name == "" {
		log.Debug("Name not given, initialising lab in current directory.")
		options.Name = filepath.Base(nk.InitialWorkDir)
		err := validator.New().Var(options.Name, "alphanum,max=30")
		if err != nil {
			return err
		}
	} else {
		err := validator.New().Var(options.Name, "alphanum,max=30")
		if err != nil {
			return err
		}
		if fileExists(options.Name) {
			return fmt.Errorf("file or directory %s already exists", options.Name)
		}
		err = os.Mkdir(options.Name, 0755)
		if err != nil {
			return err
		}
		err = os.Chdir(options.Name)
		if err != nil {
			return err
		}
	}
	// TODO check if in script mode
	// ask for name, description etc
	vpl := viper.New()
	vpl.Set("koble_version", VERSION)
	vpl.Set("created_at", time.Now().Format("02-01-2006"))
	if options.Description != "" {
		vpl.Set("description", options.Description)
	}
	if len(options.Authors) != 0 {
		vpl.Set("authors", options.Authors)
	}
	if len(options.Emails) != 0 {
		vpl.Set("emails", options.Emails)
	}
	if len(options.Webs) != 0 {
		vpl.Set("webs", options.Webs)
	}
	err := vpl.SafeWriteConfigAs("lab.yml")
	if err != nil {
		return err
	}
	err = os.Mkdir("shared", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	err = os.WriteFile("shared.startup", []byte(SHARED_STARTUP), 0644)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	return nil
}

// redo with new viper config
func (nk *Koble) AddMachineToLab(name string, conf driver.MachineConfig) error {
	if nk.LabRoot == "" {
		return errors.New("lab.yml does not exist, are you in a lab directory?")
	}
	err := validator.New().Var(name, "alphanum,max=30")
	if err != nil {
		return err
	}

	if _, ok := nk.Lab.Machines[name]; ok {
		return fmt.Errorf("a machine named %s already exists", name)
	}

	err = os.Mkdir(name, 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	fn := name + ".startup"
	err = os.WriteFile(fn, []byte(DEFAULT_STARTUP), 0644)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	vpl := viper.New()
	vpl.SetConfigName("lab")
	vpl.SetConfigType("yaml")
	vpl.AddConfigPath(nk.LabRoot)
	err = vpl.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read lab.yml: %w", err)
	}
	// write networks even if empty, so that machine gets added to lab.yml
	vpl.Set("machines."+name+".networks", conf.Networks)
	if conf.Image != "" {
		vpl.Set("machines."+name+".image", conf.Image)
	}
	err = vpl.WriteConfigAs("lab.yml")
	if err != nil {
		return err
	}
	fmt.Printf("Created new machine %s, with directory for machine files and %s.startup as the machine startup script.\n", name, name)
	return nil
}

// redo with new viper config
func AddNetworkToLab(name string, external bool, gateway net.IP, subnet net.IPNet, ipv6 bool) error {
	if gateway.String() != "<nil>" {
		if subnet.IP == nil {
			return errors.New("To use a specified gateway you need to also specify a subnet.")
		} else if !subnet.Contains(gateway) {
			return fmt.Errorf("Gateway %s is not in subnet %s.", gateway.String(), subnet.String())
		}
	}
	lab := Lab{}
	// exists, err := GetLab(&lab)
	// if err != nil {
	// 	return err
	// }
	// if !exists {
	// 	return errors.New("lab.yml does not exist, are you in a lab directory?")
	// }
	err := validator.New().Var(name, "alphanum,max=30")
	if err != nil {
		return err
	}
	for nn := range lab.Networks {
		if nn == name {
			return fmt.Errorf("A network with the name %s already exists.", name)
		}
	}
	net := driver.NetConfig{
		External: external,
		//Gateway:  gateway,
		Subnet: subnet.String(),
		//IPv6:     ipv6,
	}

	if net.Subnet == "<nil>" {
		net.Subnet = ""
	}
	lab.Networks[name] = net
	err = SaveLab(&lab)
	if err != nil {
		return err
	}
	fmt.Printf("Created new network %s.\n", name)
	return nil
}

func barText(char byte, msg string, length int) string {
	remaining := length - len(msg) - 4
	if remaining <= 0 {
		remaining = 0
	}
	space := " "
	if msg == "" {
		space = "="
	}
	if remaining%2 != 0 {
		msg = msg + space
	}
	msg = space + msg + space
	padding := strings.Repeat(string(char), remaining/2)
	return blue(fmt.Sprintf(" %s%s%s \n", padding, msg, padding))
}

func itemText(key, value string, width int) string {
	if value == "" {
		value = "<unknown>"
	}
	remaining := width - len(key+value) - 3
	if remaining <= 0 {
		remaining = 0
	}
	padding := strings.Repeat(" ", remaining)
	return fmt.Sprintf(" %s:%s%s \n", magBold(key), padding, mag(value))
}

func itemTextArray(key string, values []string, width int) string {
	if len(values) >= 2 {
		key = key + "s"
	}
	return itemText(key, strings.Join(values, ", "), width)
}

func (lab *Lab) Header() string {
	var header string
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return fmt.Sprintf("Could not get terminal size to render lab header: %v", err)
	}
	if width > MAXPRINTWIDTH {
		width = MAXPRINTWIDTH
	}
	header += barText('=', "Starting Lab", width)
	header += itemText("Lab Directory", lab.Directory, width)
	header += itemText("Created At", lab.CreatedAt, width)
	header += itemText("Version", lab.KobleVersion, width)
	header += itemTextArray("Author", lab.Authors, width)
	header += itemTextArray("Email", lab.Emails, width)
	header += itemTextArray("Web", lab.Web, width)
	header += itemText("Description", lab.Description, width)
	header += barText('=', "", width)
	return header + "\n"
}

func (nk *Koble) LabStart(mlist []string, wait bool) error {
	return nk.ForMachine(nk.Lab.Header, mlist, func(name string,
		mconf driver.MachineConfig,
		c output.Container) (err error) {
		out := c.AddOutput(fmt.Sprintf("Starting machine %s", name))
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Started machine %s", name))
			}
		}()
		out.Start()
		err = nk.StartMachine(name, mconf, out)
		if err != nil {
			return err
		}
		if wait {
			m, err := nk.Driver.Machine(name, nk.Config.Namespace)
			if err != nil {
				return err
			}
			out.Write([]byte("booting"))
			return m.WaitUntil(60*5, driver.BootedState(), driver.ExitedState())
		}
		return nil
	})
}

func contains(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func filterMachines(machines map[string]driver.MachineConfig,
	filter []string) (mList map[string]driver.MachineConfig) {
	// if no machines in filter then all machines are included
	if len(filter) == 0 {
		return machines
	}
	// only keep machines which are in the filter list
	for _, name := range filter {
		if val, ok := machines[name]; ok {
			mList[name] = val
		}
	}
	return mList
}

func (nk *Koble) GetMachineList(mlist []string,
	all bool) (machines []driver.Machine, err error) {
	if len(mlist) == 0 && nk.Lab.Name == "" && !all {
		return machines, errors.New("You are not in a lab. Use --all or specify machines.")
	} else if all && len(mlist) != 0 {
		return machines, errors.New("You cannot specify machines when using --all")
	}
	// output, err := nk.Driver.ListMachines(nk.Lab.Name, all)
	// for _, m := range output {
	// 	machines = append(machines, driver.Machine{
	// 		Name: m.Name,
	// 		Lab:  m.Lab,
	// 	})
	// }
	// machines = filterMachines(machines, mlist)
	return machines, nil
}

func (nk *Koble) LabDestroy(mlist []string) error {
	return nk.ForMachine(nk.Lab.Header, mlist, func(name string,
		mconf driver.MachineConfig,
		c output.Container) (err error) {
		out := c.AddOutput(fmt.Sprintf("Destroying machine %s", name))
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Destroyed machine %s", name))
			}
		}()
		out.Start()
		return nk.DestroyMachine(name, out)
	})
}

func (nk *Koble) LabRemove(mlist []string) error {
	return nk.ForMachine(nk.Lab.Header, mlist, func(name string,
		mconf driver.MachineConfig,
		c output.Container) (err error) {
		out := c.AddOutput(fmt.Sprintf("Removing machine %s", name))
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Removed machine %s", name))
			}
		}()
		out.Start()
		return nk.RemoveMachine(name, out)
	})
}

func (nk *Koble) LabStop(mlist []string,
	force, wait bool) error {
	return nk.ForMachine(nk.Lab.Header, mlist, func(name string,
		mconf driver.MachineConfig,
		c output.Container) (err error) {
		out := c.AddOutput(fmt.Sprintf("Stopping machine %s", name))
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Stopped machine %s", name))
			}
		}()
		out.Start()
		err = nk.StopMachine(name, force, out)
		if err != nil {
			return err
		}
		if wait {
			m, err := nk.Driver.Machine(name, nk.Config.Namespace)
			if err != nil {
				return err
			}
			out.Write([]byte("halting"))
			return m.WaitUntil(60*5, driver.ExitedState(), nil)
		}
		return nil
	})
}

func (nk *Koble) LabInfo() error {
	if nk.LabRoot == "" {
		return errors.New("You are not currently in a lab directory.")
	}
	err := nk.ListMachines(false, false)
	if err != nil {
		return err
	}
	fmt.Printf("\n")
	err = nk.ListNetworks(false)
	fmt.Printf("\n")
	return err
}

func (nk *Koble) ForMachine(headerFunc func() string, filterList []string, toRun func(name string, mconf driver.MachineConfig, c output.Container) error) error {
	if nk.LabRoot == "" {
		return errors.New("You are not currently in a lab directory.")
	}
	fmt.Println("gonna run for each with", nk.Lab.Machines)
	oc := output.NewContainer(nk.Lab.Header, nk.Config.NonInteractive)
	oc.Start()
	defer oc.Stop()
	machines := filterMachines(nk.Lab.Machines, filterList)
	var wg sync.WaitGroup
	for name, mconf := range machines {
		wg.Add(1)
		go func(name string, mconf driver.MachineConfig, c output.Container) error {
			defer wg.Done()
			return toRun(name, mconf, oc)
		}(name, mconf, oc)
	}
	wg.Wait()
	return nil
}
