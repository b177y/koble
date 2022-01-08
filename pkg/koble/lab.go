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

type Lab struct {
	Name         string    `yaml:"name,omitempty" validate:"alphanum,max=30"`
	Directory    string    `yaml:"dir,omitempty"`
	CreatedAt    string    `yaml:"created_at,omitempty" validate:"datetime"`
	KobleVersion string    `yaml:"koble_version,omitempty"`
	Description  string    `yaml:"description,omitempty"`
	Authors      []string  `yaml:"authors,omitempty"`
	Emails       []string  `yaml:"emails,omitempty" validate:"email"`
	Web          []string  `yaml:"web,omitempty" validate:"url"`
	Machines     []Machine `yaml:"machines,omitempty"`
	Networks     []Network `yaml:"networks,omitempty"`
	DefaultImage string    `yaml:"default_image,omitempty"`
}

type Machine struct {
	Name         string
	Image        string
	Networks     []Network
	Volumes      []string
	Dependencies []string
	HostHome     bool
}

func InitLab(name string, description string, authors []string, emails []string, web []string) error {
	newDir := true
	if name == "" {
		log.Debug("Name not given, initialising lab in current directory.")
		newDir = false
		exists, err := fileExists("lab.yml")
		if err != nil {
			return err
		}
		if exists {
			return errors.New("lab.yml already exists in this directory.")
		}
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		name = filepath.Base(dir)
	}
	err := validator.New().Var(name, "alphanum,max=30")
	if err != nil {
		return err
	}
	exists, err := fileExists(name)
	if err != nil {
		return err
	}
	if exists {
		info, err := os.Stat(name)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fmt.Errorf("%s already exists as a directory. To initialise it as a Koble lab directory, cd to it then run init with no name.", name)
		} else {
			return fmt.Errorf("A file named %s exists. Please use a different name to initialise the lab or rename the file.", name)
		}
	}
	pathPrefix := ""
	if newDir {
		os.Mkdir(name, 0755)
		pathPrefix = name
	}
	// TODO check if in script mode
	// ask for name, description etc
	lab := Lab{
		Description:  description,
		KobleVersion: VERSION,
		Authors:      authors,
		Emails:       emails,
		Web:          web,
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

func AddMachineToLab(name string, networks []string, image string) error {
	lab := Lab{}
	exists, err := GetLab(&lab)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("lab.yml does not exist, are you in a lab directory?")
	}
	err = validator.New().Var(name, "alphanum,max=30")
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

	for _, m := range lab.Machines {
		if m.Name == name {
			return fmt.Errorf("A machine with the name %s already exists.", name)
		}
	}
	lab.Machines = append(lab.Machines, Machine{Name: name})
	err = SaveLab(&lab)
	// TODO print help for getting started with machine
	if err != nil {
		return err
	}
	fmt.Printf("Created new machine %s, with directory for machine files and %s.startup as the machine startup script.\n", name, name)
	return nil
}

func AddNetworkToLab(name string, external bool, gateway net.IP, subnet net.IPNet, ipv6 bool) error {
	if gateway.String() != "<nil>" {
		if subnet.IP == nil {
			return errors.New("To use a specified gateway you need to also specify a subnet.")
		} else if !subnet.Contains(gateway) {
			return fmt.Errorf("Gateway %s is not in subnet %s.", gateway.String(), subnet.String())
		}
	}
	lab := Lab{}
	exists, err := GetLab(&lab)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("lab.yml does not exist, are you in a lab directory?")
	}
	err = validator.New().Var(name, "alphanum,max=30")
	if err != nil {
		return err
	}
	for _, n := range lab.Networks {
		if n.Name == name {
			return fmt.Errorf("A network with the name %s already exists.", name)
		}
	}
	net := Network{
		Name:     name,
		External: external,
		Gateway:  gateway,
		Subnet:   subnet.String(),
		IPv6:     ipv6,
	}

	if net.Subnet == "<nil>" {
		net.Subnet = ""
	}
	lab.Networks = append(lab.Networks, net)
	err = SaveLab(&lab)
	if err != nil {
		return err
	}
	fmt.Printf("Created new network %s.\n", name)
	return nil
}

func (nk *Koble) Validate() error {
	// do some extra validation here
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

func (nk *Koble) LabHeader() (header string) {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return fmt.Sprintf("Could not get terminal size to render lab header: %v", err)
	}
	if width > MAXPRINTWIDTH {
		width = MAXPRINTWIDTH
	}
	header += barText('=', "Starting Lab", width)
	header += itemText("Lab Directory", nk.Lab.Directory, width)
	header += itemText("Version", nk.Lab.KobleVersion, width)
	header += itemTextArray("Author", nk.Lab.Authors, width)
	header += itemTextArray("Email", nk.Lab.Emails, width)
	header += itemTextArray("Web", nk.Lab.Web, width)
	header += itemText("Description", nk.Lab.Description, width)
	header += barText('=', "", width)
	return header + "\n"
}

func (nk *Koble) LabStart(mlist []string) error {
	oc := output.NewContainer(false)
	oc.Start()
	defer oc.Stop()
	if nk.Lab.Name == "" {
		return errors.New("You are not currently in a lab directory.")
	}
	fmt.Printf(nk.LabHeader())
	machines := filterMachines(nk.Lab.Machines, mlist)
	var wg sync.WaitGroup
	for _, m := range machines {
		wg.Add(1)
		go func(m Machine) error {
			defer wg.Done()
			output := oc.AddOutput(fmt.Sprintf("Starting machine %s", m.Name), fmt.Sprintf("Started machine %s", m.Name))
			output.Start()
			time.Sleep(1 * time.Second)
			output.Write([]byte("loading"))
			time.Sleep(1 * time.Second)
			output.Write([]byte("waiting"))
			time.Sleep(1 * time.Second)
			output.Write([]byte("starting networks"))
			time.Sleep(1 * time.Second)
			output.Write([]byte("starting machine"))
			time.Sleep(1 * time.Second)
			output.Write([]byte("booting"))
			time.Sleep(2 * time.Second)
			output.Finished()
			return nil
			// m, err := nk.Driver.Machine(m.Name, nk.Namespace)
			// if err != nil {
			// 	return err
			// }
			// err = m.Start(nil)
			// if err != nil {
			// 	return err
			// }
		}(m)
	}
	wg.Wait()
	return nil
}

func contains(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func filterMachines(machines []Machine,
	filter []string) (mList []Machine) {
	// if no machines in filter then all machines are included
	if len(filter) == 0 {
		return machines
	}
	// only keep machines which are in the filter list
	for _, m := range machines {
		if contains(filter, m.Name) {
			mList = append(mList, m)
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

func (nk *Koble) LabDestroy(mlist []string, all bool) error {
	machines := filterMachines(nk.Lab.Machines, mlist)
	for _, m := range machines {
		err := nk.DestroyMachine(m.Name)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func (nk *Koble) LabHalt(mlist []string,
	force, all bool) error {
	machines := filterMachines(nk.Lab.Machines, mlist)
	for _, m := range machines {
		err := nk.HaltMachine(m.Name, force)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func (nk *Koble) LabInfo() error {
	if nk.Lab.Name == "" {
		return errors.New("You are not in a lab right now...")
	}
	fmt.Println("============================ Lab ===============================")
	var info [][]string
	info = append(info, []string{"Name", nk.Lab.Name})
	info = append(info, []string{"Directory", nk.Lab.Directory})
	info = append(info, []string{"Created At", nk.Lab.CreatedAt})
	info = append(info, []string{"Koble Version", nk.Lab.KobleVersion})
	authorHeading, authors := multiHeading("Author", nk.Lab.Authors)
	info = append(info, []string{authorHeading, authors})
	emailHeading, emails := multiHeading("Email", nk.Lab.Emails)
	info = append(info, []string{emailHeading, emails})
	webHeading, web := multiHeading("URL", nk.Lab.Web)
	info = append(info, []string{webHeading, web})
	info = append(info, []string{"Description", nk.Lab.Description})
	RenderTable([]string{}, info)
	fmt.Printf("================================================================\n\n")
	err := nk.ListMachines(false, false)
	if err != nil {
		return err
	}
	fmt.Printf("\n")
	err = nk.ListNetworks(false)
	fmt.Printf("\n")
	return err
}