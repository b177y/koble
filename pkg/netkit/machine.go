package netkit

import (
	"fmt"
	"os"

	"github.com/b177y/netkit/driver"
	prettyjson "github.com/hokaccha/go-prettyjson"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

func (nk *Netkit) StartMachine(name, image string, networks []string) error {
	// Start with defaults
	m, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}
	opts := driver.StartOptions{
		Lab:      nk.Lab.Name,
		Hostlab:  nk.Lab.Directory,
		HostHome: true,
		Networks: networks,
	}

	for _, n := range networks {
		err := nk.StartNetwork(n)
		if err != nil && err != driver.ErrExists {
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
	fmt.Printf("Starting machine %s\n", m.Name())
	return m.Start(&opts)
}

func (nk *Netkit) MachineInfo(name string, json bool) error {
	m, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}
	info, err := m.Info()
	if err != nil {
		return err
	}
	if json {
		s, err := prettyjson.Marshal(info)
		if err != nil {
			return err
		}
		fmt.Println(string(s))
	} else {
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
		if info.Image != "" {
			infoTable = append(infoTable, []string{"Image", info.Image})
		}
		if info.State != "" {
			infoTable = append(infoTable, []string{"State", info.State})
		}
		RenderTable([]string{}, infoTable)
	}
	return nil
}

func (nk *Netkit) HaltMachine(machine string, force bool) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	fmt.Printf("Halting machine %s\n", m.Name())
	return m.Stop(force)
}

func (nk *Netkit) RemoveMachine(machine string) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	fmt.Printf("Removing machine %s\n", m.Name())
	return m.Remove()
}

func (nk *Netkit) DestroyMachine(machine string) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	fmt.Printf("Crashing machine %s\n", m.Name())
	err = m.Stop(true)
	if err != nil {
		return err
	}
	// TODO workout best way to delay until machine stopped
	// m.WaitUntil() ?
	fmt.Printf("Removing machine %s\n", m.Name())
	return m.Remove()
}

func (nk *Netkit) MachineLogs(machine string, follow bool, tail int) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	opts := driver.LogOptions{
		Follow: follow,
		Tail:   tail,
	}
	return m.Logs(&opts)
}

func (nk *Netkit) ListMachines(all, json bool) error {
	machines, err := nk.Driver.ListMachines(nk.Namespace, all)
	if err != nil {
		return err
	}
	if json {
		s, err := prettyjson.Marshal(machines)
		if err != nil {
			return err
		}
		fmt.Println(string(s))
	} else {
		mlist, headers := MachineInfoToStringArr(machines, all)
		RenderTable(headers, mlist)
	}
	return nil
}
