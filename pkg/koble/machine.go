package koble

import (
	"fmt"
	"io"

	"github.com/b177y/koble/driver"
	prettyjson "github.com/hokaccha/go-prettyjson"
)

//func (nk *Koble) StartMachine(name, image string, networks []string, out io.Writer) error {
func (nk *Koble) StartMachine(name string, conf driver.MachineConfig, out io.Writer) error {
	// Start with defaults
	m, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}

	for _, n := range conf.Networks {
		fmt.Fprintf(out, "creating network %s", n)
		err := nk.StartNetwork(n, driver.NetConfig{}) // TODO get netconfig from Lab
		if err != nil && err != driver.ErrExists {
			return err
		}
	}

	for _, dependency := range conf.Dependencies {
		dep, err := nk.Driver.Machine(dependency, nk.Namespace)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "waiting for dependency %s to boot", dependency)
		err = dep.WaitUntil(60*5, driver.BootedState(), nil)
		if err != nil {
			return fmt.Errorf("Error waiting for dependency %s to boot: %w", dependency, err)
		}
	}
	// conf.Volumes = append(conf.Volumes, spec.Mount{
	// 	Source:      nk.Lab.Directory,
	// 	Destination: "/hostlab",
	// })

	return m.Start(&conf)
}

func (nk *Koble) MachineInfo(name string, json bool) error {
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

func (nk *Koble) StopMachine(name string, force bool, out io.Writer) error {
	m, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}
	if force {
		fmt.Fprintf(out, "Crashing machine %s", name)
	} else {
		fmt.Fprintf(out, "Halting machine %s", name)
	}
	return m.Stop(force)
}

func (nk *Koble) RemoveMachine(name string, out io.Writer) error {
	m, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "removing machine %s", name)
	return m.Remove()
}

func (nk *Koble) DestroyMachine(machine string, out io.Writer) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		fmt.Fprintf(out, "no machine to remove")
		return nil
	}
	fmt.Fprintf(out, "Crashing machine %s", m.Name())
	err = m.Stop(true)
	if err != nil {
		return err
	}
	err = m.WaitUntil(120, driver.ExitedState(), nil)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Removing machine %s", m.Name())
	return m.Remove()
}

func (nk *Koble) MachineLogs(machine string, follow bool, tail int) error {
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

func (nk *Koble) ListMachines(all, json bool) error {
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
