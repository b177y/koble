package netkit

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

type PodmanMachineExtra struct {
	AddCaps       []string `yaml:"add_caps"`
	RmCaps        []string `yaml:"remove_caps"`
	Volumes       []string `yaml:"volumes"`
	MountHostHome bool     `yaml:"mount_host_home"`
}

type Network struct {
	Name     string `yaml:"name" validate:"alphanum,max=30"`
	Internal bool   `yaml:"external,omitempty"`
	Gateway  string `yaml:"gateway,omitempty" validate:"ip"`
	IpRange  string `yaml:"ip_range,omitempty" validate:"cidr"`
	Subnet   string `yaml:"subnet,omitempty" validate:"cidr"`
	IPv6     string `yaml:"ipv6,omitempty" validate:"ipv6"`
}

type Machine struct {
	Name        string             `yaml:"name" validate:"alphanum,max=30"`
	Networks    []Network          `yaml:"networks,omitempty,flow"`
	Image       string             `yaml:"image,omitempty"`
	PodmanExtra PodmanMachineExtra `yaml:"podman_extra,omitempty,flow"`
}

type Lab struct {
	Name          string    `yaml:"name,omitempty" validate:"alphanum,max=30"`
	CreatedAt     string    `yaml:"created_at,omitempty,flow" validate:"datetime"`
	NetkitVersion string    `yaml:"netkit_version,omitempty,flow"`
	Description   string    `yaml:"description,omitempty"`
	Authors       []string  `yaml:"authors,omitempty,flow"`
	Emails        []string  `yaml:"emails,omitempty,flow" validate:"email"`
	Web           []string  `yaml:"web,omitempty,flow" validate:"url"`
	Machines      []Machine `yaml:"machines,omitempty,flow"`
}

func InitLab(name string, description string, authors []string, emails []string, web []string) error {
	// check if names are empty etc
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
	// check if in script mode
	lab := Lab{
		Description:   description,
		Machines:      []Machine{},
		NetkitVersion: VERSION,
	}
	lab.CreatedAt = time.Now().Format("02-01-2006")
	fmt.Print(lab)
	bytes, err := yaml.Marshal(lab)
	if err != nil {
		return err
	}
	err = os.WriteFile(fn, bytes, 0644)
	// ask for name, description etc
	return err
}
