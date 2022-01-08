package koble

import (
	"fmt"
	"io"
	"os"

	"github.com/b177y/koble/driver"
	"github.com/b177y/koble/pkg/output"
	prettyjson "github.com/hokaccha/go-prettyjson"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

func (nk *Koble) StartMachineWithStatus(name, image string, networks []string, wait, plain bool) (err error) {
	oc := output.NewContainer(nil, plain)
	oc.Start()
	out := oc.AddOutput(fmt.Sprintf("Starting machine %s", name))
	out.Start()
	defer func() {
		if err != nil {
			out.Error(err)
		} else {
			out.Success(fmt.Sprintf("Started machine %s", name))
		}
		oc.Stop()
	}()
	machine := Machine{
		Name:  name,
		Image: image,
		//Networks: networks,
	}
	err = nk.StartMachine(machine, out)
	if err != nil {
		return err
	}
	if wait {
		m, err := nk.Driver.Machine(name, nk.Namespace)
		if err != nil {
			return err
		}
		out.Write([]byte("booting"))
		return m.WaitUntil("running", 60*5)
	}
	return nil
}

func netStrList(nets []Network) (names []string) {
	for _, n := range nets {
		names = append(names, n.Name)
	}
	return names
}

//func (nk *Koble) StartMachine(name, image string, networks []string, out io.Writer) error {
func (nk *Koble) StartMachine(machine Machine, out io.Writer) error {
	// Start with defaults
	m, err := nk.Driver.Machine(machine.Name, nk.Namespace)
	if err != nil {
		return err
	}
	opts := driver.StartOptions{
		Lab:      nk.Lab.Name,
		Hostlab:  nk.Lab.Directory,
		HostHome: true,
		Networks: netStrList(machine.Networks),
	}

	for _, n := range machine.Networks {
		out.Write([]byte("creating network " + n.Name))
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
	if machine.Image != "" {
		opts.Image = machine.Image
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
	// fmt.Printf("Starting machine %s\n", m.Name())
	return m.Start(&opts)
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

func (nk *Koble) HaltMachine(machine string, force bool, out io.Writer) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	if force {
		fmt.Fprintf(out, "Crashing machine %s\n", m.Name())
	} else {
		fmt.Fprintf(out, "Halting machine %s\n", m.Name())
	}
	return m.Stop(force)
}

func (nk *Koble) RemoveMachine(machine string, out io.Writer) error {
	m, err := nk.Driver.Machine(machine, nk.Namespace)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "removing machine %s\n", m.Name())
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
