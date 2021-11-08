package netkit

import (
	"fmt"
	"os"

	"github.com/b177y/netkit/driver"
	spec "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
)

func (nk *Netkit) StartMachine(name, image string, networks []string) error {
	// Start with defaults
	m := driver.Machine{
		Name:     name,
		Lab:      nk.Lab.Name,
		Hostlab:  nk.Lab.Directory,
		HostHome: false,
		Networks: []string{},
		Image:    nk.Driver.GetDefaultImage(),
	}
	log.Debug("defaults", m)

	for _, n := range networks {
		fmt.Printf("Starting network %s\n", n)
		err := nk.StartNetwork(n)
		if err != nil && err != driver.ErrExists {
			return err
		}
		_, err = nk.Driver.GetNetworkState(driver.Network{
			Name: n,
			Lab:  nk.Lab.Name,
		})
		if err != nil {
			return err
		}
	}

	// Add options from lab
	for _, lm := range nk.Lab.Machines {
		if lm.Name == m.Name {
			m.Volumes = lm.Volumes
			m.HostHome = lm.HostHome
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
	if m.HostHome {
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
	fmt.Printf("Starting %s...\n", m.Name)
	err := nk.Driver.StartMachine(m)
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
	m := driver.Machine{
		Name: machine,
		Lab:  nk.Lab.Name,
	}
	err := nk.Driver.GetMachineLogs(m, stdoutChan,
		stderrChan, follow, tail)
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
	if err != nil {
		return err
	}
	mlist, headers := MachineInfoToStringArr(machines, all)
	RenderTable(headers, mlist)
	return nil
}
