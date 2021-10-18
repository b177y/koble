package netkit

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

type Lab struct {
	Name          string `validate:"alphanum,max=30"`
	Directory     string
	CreatedAt     string             `yaml:"created_at,omitempty" validate:"datetime"`
	NetkitVersion string             `yaml:"netkit_version,omitempty"`
	Description   string             `yaml:"description,omitempty"`
	Authors       []string           `yaml:"authors,omitempty"`
	Emails        []string           `yaml:"emails,omitempty" validate:"email"`
	Web           []string           `yaml:"web,omitempty" validate:"url"`
	Machines      []Machine          `yaml:"machines,omitempty"`
	Networks      []Network          `yaml:"networks,omitempty"`
	DefaultImage  string             `yaml:"default_image,omitempty"`
	PodmanExtra   PodmanMachineExtra `yaml:"podman_extra,omitempty"`
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
			return fmt.Errorf("%s already exists as a directory. To initialise it as a Netkit lab directory, cd to it then run init with no name.", name)
		} else {
			return fmt.Errorf("A file named %s exists. Please use a different name to initialise the lab or rename the file.", name)
		}
	}
	fn := "lab.yml"
	if newDir {
		os.Mkdir(name, 0755)
		fn = name + "/" + fn
	}
	// TODO check if in script mode
	// ask for name, description etc
	lab := Lab{
		Description:   description,
		NetkitVersion: VERSION,
		Authors:       authors,
		Emails:        emails,
		Web:           web,
	}
	lab.CreatedAt = time.Now().Format("02-01-2006")
	fmt.Print(lab)
	bytes, err := yaml.Marshal(lab)
	if err != nil {
		return err
	}
	err = os.WriteFile(fn, bytes, 0644)
	return err
}

func AddMachineToLab(name string, networks []string, image string) error {
	lab := Lab{}
	exists, err := getLab(&lab)
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
	lab.Machines = append(lab.Machines, Machine{
		Name:     name,
		Image:    image,
		Networks: networks,
	})
	err = saveLab(&lab)
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
	exists, err := getLab(&lab)
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
	err = saveLab(&lab)
	if err != nil {
		return err
	}
	fmt.Printf("Created new network %s.\n", name)
	return nil
}
