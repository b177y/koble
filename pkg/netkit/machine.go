package netkit

import (
	"fmt"
	"os"

	"github.com/b177y/netkit/driver"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

func (nk *Netkit) StartMachine(name, image string, networks []string) error {
	// Start with defaults
	m, err := nk.Driver.Machine(name)
	if err != nil {
		return err
	}
	opts := driver.StartOptions{
		Lab:      nk.Lab.Name,
		Hostlab:  nk.Lab.Directory,
		HostHome: true,
		// Networks: networks,
		Image: nk.Driver.GetDefaultImage(),
	}

	for _, n := range networks {
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
		if lm.Name == m.Name() {
			// opts.Volumes = lm.Volumes
			opts.HostHome = lm.HostHome
			if lm.Image != "" {
				opts.Image = lm.Image
			}
			// opts.Networks = lm.Networks
		}
	}
	// Add options from command line flags
	if image != "" {
		opts.Image = image
	}
	if opts.HostHome {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		opts.Volumes = append(opts.Volumes, spec.Mount{
			Source:      home,
			Destination: "/hosthome",
		})
	}
	opts.Volumes = append(opts.Volumes, spec.Mount{
		Source:      nk.Lab.Directory,
		Destination: "/hostlab",
	})
	fmt.Println("Starting machine", m.Name())
	return m.Start(&opts)
}

func (nk *Netkit) MachineInfo(name string) error {
	m, err := nk.Driver.Machine(name)
	var infoTable [][]string
	infoTable = append(infoTable, []string{"Name", m.Name()})
	// if nk.Lab.Name != "" {
	// 	for _, lm := range nk.Lab.Machines {
	// 		if lm.Name == m.Name() {
	// 			// lm.Lab = m.Lab
	// 			// m = lm
	// 			if lm.Image != "" {
	// 				infoTable = append(infoTable,
	// 					[]string{"Image", lm.Image})
	// 			}
	// 			// if len(lm.Dependencies) != 0 {
	// 			// 	infoTable = append(infoTable,
	// 			// 		[]string{"Dependencies", strings.Join(lm.Dependencies, ",")})
	// 			// }
	// 			if len(lm.Networks) != 0 {
	// 				infoTable = append(infoTable,
	// 					[]string{"Networks", strings.Join(lm.Networks, ",")})
	// 			}
	// 			if len(lm.Volumes) != 0 {
	// 				var vols []string
	// 				for _, v := range lm.Volumes {
	// 					vols = append(vols, v.Source+":"+v.Destination)
	// 				}
	// 				infoTable = append(infoTable,
	// 					[]string{"Volumes", strings.Join(vols, ",")})
	// 			}
	// 		}
	// 	}
	// }
	info, err := m.Info()
	if err != nil && err != driver.ErrNotExists {
		return err
	}
	if info.Image != "" {
		infoTable = append(infoTable, []string{"Image", info.Image})
	}
	if info.State != "" {
		infoTable = append(infoTable, []string{"State", info.State})
	}
	RenderTable([]string{}, infoTable)
	return nil
}

func (nk *Netkit) HaltMachine(machine string, force bool) error {
	m, err := nk.Driver.Machine(machine)
	if err != nil {
		return err
	}
	fmt.Println("Halting machine", m.Name)
	return m.Stop(force)
}

func (nk *Netkit) DestroyMachine(machine string) error {
	m, err := nk.Driver.Machine(machine)
	if err != nil {
		return err
	}
	fmt.Println("Crashing machine", m.Name)
	err = m.Stop(true)
	if err != nil {
		return err
	}
	// TODO workout best way to delay until machine stopped
	// time.Sleep(time.Second)
	fmt.Println("Removing machine", m.Name)
	return m.Remove()
}

func (nk *Netkit) MachineLogs(machine string, follow bool, tail int) error {
	m, err := nk.Driver.Machine(machine)
	if err != nil {
		return err
	}
	opts := driver.LogOptions{
		Follow: follow,
		Tail:   tail,
	}
	return m.Logs(&opts)
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
