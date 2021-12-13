package netkit

import (
	"fmt"
	"os"
	"strings"

	"github.com/b177y/netkit/driver"
	spec "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
)

func (nk *Netkit) StartMachine(name, image string, networks []string) error {
	// Start with defaults
	m := driver.Machine{
		Name:      name,
		Lab:       nk.Lab.Name,
		Namespace: nk.Namespace,
		Hostlab:   nk.Lab.Directory,
		HostHome:  false,
		Networks:  []string{},
		Image:     nk.Driver.GetDefaultImage(),
	}
	log.Debug("defaults", m)

	for _, n := range networks {
		fmt.Printf("Starting network %s\n", n)
		err := nk.StartNetwork(n)
		if err != nil && err != driver.ErrExists {
			return err
		}
		_, err = nk.Driver.GetNetworkState(driver.Network{
			Name:      n,
			Namespace: nk.Namespace,
			Lab:       nk.Lab.Name,
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

func (nk *Netkit) MachineInfo(name string) error {
	m := driver.Machine{
		Name:      name,
		Namespace: nk.Namespace,
		Lab:       nk.Lab.Name,
	}
	var infoTable [][]string
	infoTable = append(infoTable, []string{"Name", m.Name})
	if nk.Lab.Name != "" {
		for _, lm := range nk.Lab.Machines {
			if lm.Name == m.Name {
				lm.Lab = m.Lab
				m = lm
				if lm.Image != "" {
					infoTable = append(infoTable,
						[]string{"Image", lm.Image})
				}
				if len(lm.Dependencies) != 0 {
					infoTable = append(infoTable,
						[]string{"Dependencies", strings.Join(lm.Dependencies, ",")})
				}
				if len(lm.Networks) != 0 {
					infoTable = append(infoTable,
						[]string{"Networks", strings.Join(lm.Networks, ",")})
				}
				if len(lm.Volumes) != 0 {
					var vols []string
					for _, v := range lm.Volumes {
						vols = append(vols, v.Source+":"+v.Destination)
					}
					infoTable = append(infoTable,
						[]string{"Volumes", strings.Join(vols, ",")})
				}
			}
		}
	}
	info, err := nk.Driver.MachineInfo(m)
	fmt.Println(info)
	if err != nil && err != driver.ErrNotExists {
		return err
	}
	if info.Image != "" && m.Image == "" {
		infoTable = append(infoTable, []string{"Image", info.Image})
	}
	if info.State != "" {
		infoTable = append(infoTable, []string{"State", info.State})
	}
	RenderTable([]string{}, infoTable)
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
		Name:      machine,
		Namespace: nk.Namespace,
		Lab:       nk.Lab.Name,
	}
	err := nk.Driver.GetMachineLogs(m, stdoutChan,
		stderrChan, follow, tail)
	return err
}

func (nk *Netkit) ListMachines(all bool) error {
	machines, err := nk.Driver.ListMachines(nk.Namespace, all)
	if err != nil {
		return err
	}
	mlist, headers := MachineInfoToStringArr(machines, all)
	RenderTable(headers, mlist)
	return nil
}
