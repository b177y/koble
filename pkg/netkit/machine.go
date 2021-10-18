package netkit

import (
	"fmt"
	"os"

	"github.com/b177y/netkit/driver"
	spec "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
)

type PodmanMachineExtra struct {
	Caps []string `yaml:"caps,omitempty"`
}

type Machine struct {
	Name        string             `yaml:"name" validate:"alphanum,max=30"`
	Networks    []string           `yaml:"networks,omitempty" validate:"alphanum,max=30"`
	Image       string             `yaml:"image,omitempty"`
	Volumes     []spec.Mount       `yaml:"volumes,omitempty"`
	HostHome    bool               `yaml:"hosthome,omitempty"`
	PodmanExtra PodmanMachineExtra `yaml:"podman_extra,omitempty"`
}

func (nk *Netkit) StartMachine(name, image string, networks []string) error {
	// Start with defaults
	m := driver.Machine{
		Name:     name,
		Hostlab:  nk.Lab.Directory,
		Hosthome: false,
		Networks: []string{},
		Image:    nk.Driver.GetDefaultImage(),
	}
	log.Debug("defaults", m)

	// Add options from lab
	for _, lm := range nk.Lab.Machines {
		if lm.Name == m.Name {
			m.Volumes = lm.Volumes
			m.Hosthome = lm.HostHome
			if lm.Image != "" {
				m.Image = lm.Image
			}
			m.Networks = lm.Networks
		}
	}
	log.Debug("lab", m)
	// Add options from command line flags
	if image != "" {
		m.Image = image
	}
	if m.Hosthome {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		m.Volumes = append(m.Volumes, spec.Mount{
			Source:      home,
			Destination: "/hosthome",
		})
	}
	m.Volumes = append(m.Volumes, spec.Mount{
		Source:      nk.Lab.Directory,
		Destination: "/hostlab",
	})
	log.Debug("cli", m)
	_, err := nk.Driver.StartMachine(m, nk.Lab.Name)
	return err
}

func CrashMachine() error {
	return nil
}

func HaltMachine() error {
	return nil
}

func DestroyMachine() error {
	return nil
}

func MachineInfo() error {
	return nil
}

func (nk *Netkit) MachineLogs(machine string, follow bool, tail int) error {
	stdoutChan := make(chan string)
	stderrChan := make(chan string)
	go func() {
		for recv := range stdoutChan {
			fmt.Println(recv)
		}
	}()
	go func() {
		for recv := range stderrChan {
			fmt.Println(recv)
		}
	}()
	err := nk.Driver.GetMachineLogs(machine, nk.Lab.Name,
		stdoutChan, stderrChan, follow, tail)
	return err
}

func (nk *Netkit) ListMachines(all bool) error {
	if !all {
		if nk.Lab.Name == "" {
			fmt.Println("Listing all machines which are not associated with a lab.")
			fmt.Printf("To see all machines use `netkit machine list --all`\n\n")
		} else {
			fmt.Printf("Listing all machines within this lab (%s).\n", nk.Lab.Name)
			fmt.Printf("To see all machines use `netkit machine list --all`\n\n")
		}
	}
	machines, err := nk.Driver.ListMachines(nk.Lab.Name, all)
	var mlist [][]string
	var headers []string
	if all {
		mlist, headers = MachineInfoToStringArr(machines, true)
	} else {
		mlist, headers = MachineInfoToStringArr(machines, false)
	}
	RenderTable(headers, mlist)
	return err
}
