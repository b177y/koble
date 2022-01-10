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
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/yaml.v2"

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

func InitLab(options InitOpts) error {
	newDir := true
	if options.Name == "" {
		log.Debug("Name not given, initialising lab in current directory.")
		newDir = false
		exists := fileExists("lab.yml")
		if exists {
			return errors.New("lab.yml already exists in this directory.")
		}
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		options.Name = filepath.Base(dir)
	}
	err := validator.New().Var(options.Name, "alphanum,max=30")
	if err != nil {
		return err
	}
	exists := fileExists(options.Name)
	if exists {
		info, err := os.Stat(options.Name)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fmt.Errorf("%s already exists as a directory. To initialise it as a Koble lab directory, cd to it then run init with no name.", options.Name)
		} else {
			return fmt.Errorf("A file named %s exists. Please use a different name to initialise the lab or rename the file.", options.Name)
		}
	}
	pathPrefix := ""
	if newDir {
		os.Mkdir(options.Name, 0755)
		pathPrefix = options.Name
	}
	// TODO check if in script mode
	// ask for name, description etc
	lab := Lab{
		Description:  options.Description,
		KobleVersion: VERSION,
		Authors:      options.Authors,
		Emails:       options.Emails,
		Web:          options.Webs,
	}
	lab.CreatedAt = time.Now().Format("02-01-2006")
	bytes, err := yaml.Marshal(lab)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(pathPrefix, "lab.yml"), bytes, 0644)
	err = os.Mkdir(filepath.Join(pathPrefix, "shared"), 0755)
	if err != nil {
		// TODO warn but not error if already exists
		return err
	}
	err = os.WriteFile(filepath.Join(pathPrefix, "shared.startup"), []byte(SHARED_STARTUP), 0644)
	if err != nil {
		// TODO warn but not error if already exists
		return err
	}
	return err
}

// redo with new viper config
func AddMachineToLab(name string, networks []string, image string) error {
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
		return fmt.Errorf("Machine name %s must be alphanumeric and shorter than 30 characters: %w", name, err)
	}
	err = os.Mkdir(name, 0755)
	if err != nil {
		// TODO warn but not error if already exists
		return err
	}
	fn := name + ".startup"
	err = os.WriteFile(fn, []byte(DEFAULT_STARTUP), 0644)
	if err != nil {
		// TODO warn but not error if already exists
		return err
	}

	for mn := range lab.Machines {
		if mn == name {
			return fmt.Errorf("A machine with the name %s already exists.", name)
		}
	}
	lab.Machines[name] = driver.MachineConfig{
		Networks: networks,
		Image:    image,
	}
	err = SaveLab(&lab)
	// TODO print help for getting started with machine
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
	for nn, _ := range lab.Networks {
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
			m, err := nk.Driver.Machine(name, nk.Namespace)
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
			m, err := nk.Driver.Machine(name, nk.Namespace)
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
