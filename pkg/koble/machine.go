package koble

import (
	"fmt"
	"io"

	"github.com/b177y/koble/driver"
	prettyjson "github.com/hokaccha/go-prettyjson"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

// func (nk *Koble) StartMachineWithStatus(name, image string, networks []string, wait, plain bool) (err error) {
// 	oc := output.NewContainer(nil, plain)
// 	oc.Start()
// 	out := oc.AddOutput(fmt.Sprintf("Starting machine %s", name))
// 	out.Start()
// 	defer func() {
// 		if err != nil {
// 			out.Error(err)
// 		} else {
// 			out.Success(fmt.Sprintf("Started machine %s", name))
// 		}
// 		oc.Stop()
// 	}()
// 	err = machine.Start(out)
// 	if err != nil {
// 		return err
// 	}
// 	if wait {
// 		m, err := nk.Driver.Machine(name, nk.Namespace)
// 		if err != nil {
// 			return err
// 		}
// 		out.Write([]byte("booting"))
// 		return m.WaitUntil("running", 60*5)
// 	}
// 	return nil
// }

//func (nk *Koble) StartMachine(name, image string, networks []string, out io.Writer) error {
func (nk *Koble) StartMachine(name string, conf driver.MachineConfig, out io.Writer) error {
	// Start with defaults
	dm, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}

	for _, n := range conf.Networks {
		fmt.Fprintf(out, "creating network %s", n)
		err := nk.StartNetwork(n, driver.NetConfig{})
		if err != nil && err != driver.ErrExists {
			return err
		}
	}
	conf.Volumes = append(conf.Volumes, spec.Mount{
		Source:      nk.Lab.Directory,
		Destination: "/hostlab",
	})
	return dm.Start(&conf)
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

func (nk *Koble) HaltMachine(name string, force bool, out io.Writer) error {
	m, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}
	if force {
		fmt.Fprintf(out, "Crashing machine %s\n", name)
	} else {
		fmt.Fprintf(out, "Halting machine %s\n", name)
	}
	return m.Stop(force)
}

func (nk *Koble) RemoveMachine(name string, out io.Writer) error {
	m, err := nk.Driver.Machine(name, nk.Namespace)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "removing machine %s\n", name)
	return m.Remove()
}

func (nk *Koble) DestroyMachine(machine string, out io.Writer) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Crashing machine %s\n", m.Name())
	err = m.Stop(true)
	if err != nil {
		return err
	}
	// TODO workout best way to delay until machine stopped
	// m.WaitUntil() ?
	fmt.Fprintf(out, "Removing machine %s\n", m.Name())
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
