package koble

import (
	"fmt"
	"io"
	"strings"

	"github.com/b177y/koble/pkg/driver"
	"github.com/dustin/go-humanize"
	prettyjson "github.com/hokaccha/go-prettyjson"
	spec "github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
)

func mergeMachineConf(base driver.MachineConfig,
	overrides driver.MachineConfig) (merged driver.MachineConfig) {
	if base.Image == "" && overrides.Image != "" {
		base.Image = overrides.Image
	}
	if len(base.Networks) == 0 && len(overrides.Networks) != 0 {
		base.Networks = overrides.Networks
	}
	return base
}

func (nk *Koble) StartMachine(name string, conf driver.MachineConfig,
	attachTerm string, out io.Writer) error {
	m, err := nk.Driver.Machine(name, nk.Config.Namespace)
	if err != nil {
		return err
	}

	// merge lab machine config with cli options
	if labM, ok := nk.Lab.Machines[name]; ok {
		conf = mergeMachineConf(labM, conf)
	}

	for _, n := range conf.Networks {
		net := driver.NetConfig{}
		if labN, ok := nk.Lab.Networks[n]; ok {
			net = labN
			log.Debugf("using lab network %s: %+v\n", n, net)
		} else {
			log.Debugf("network %s not found in lab, using default\n", n)
		}
		err := nk.StartNetwork(n, net)
		if err != nil && err != driver.ErrExists {
			return err
		}
	}

	for _, dependency := range conf.Dependencies {
		dep, err := nk.Driver.Machine(dependency, nk.Config.Namespace)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "waiting for dependency %s to boot", dependency)
		err = dep.WaitUntil(60*5, driver.BootedState(), nil)
		if err != nil {
			return fmt.Errorf("Error waiting for dependency %s to boot: %w", dependency, err)
		}
	}

	if conf.Hostlab {
		conf.Volumes = append(conf.Volumes, spec.Mount{
			Source:      nk.LabRoot,
			Destination: "/hostlab",
		})
	}

	err = m.Start(&conf)
	if err != nil {
		return err
	}
	if attachTerm != "" && attachTerm != "this" {
		err = nk.AttachToMachine(name, attachTerm)
		if err != nil {
			return err
		}
	}
	if waitTimeout := nk.Config.Wait; waitTimeout > 0 &&
		attachTerm != "this" {
		fmt.Fprintln(out, "booting")
		return m.WaitUntil(waitTimeout, driver.BootedState(), driver.ExitedState())
	}
	return nil
}

func (nk *Koble) MachineInfo(name string, json bool) error {
	m, err := nk.Driver.Machine(name, nk.Config.Namespace)
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
		infoTable = append(infoTable, []string{"NAME", m.Name()})
		infoTable = append(infoTable, []string{"PID", fmt.Sprint(info.Pid)})
		infoTable = append(infoTable, []string{"IMAGE", info.Image})
		infoTable = append(infoTable, []string{"STATE", info.State})
		if !info.StartedAt.IsZero() {
			infoTable = append(infoTable, []string{"CREATED AT", humanize.Time(info.CreatedAt)})
		} else {
			infoTable = append(infoTable, []string{"CREATED AT", ""})
		}
		if !info.StartedAt.IsZero() {
			infoTable = append(infoTable, []string{"STARTED AT", humanize.Time(info.StartedAt)})
		} else {
			infoTable = append(infoTable, []string{"STARTED AT", ""})
		}
		// infoTable = append(infoTable, []string{"VOLUMES", strings.Join(info.Volumes, ", ")})
		// infoTable = append(infoTable, []string{"PORTS", ""})
		infoTable = append(infoTable, []string{"NETWORKS", strings.Join(info.Networks, ", ")})
		RenderTable([]string{}, infoTable)
	}
	return nil
}

func (nk *Koble) StopMachine(name string, force bool, out io.Writer) error {
	m, err := nk.Driver.Machine(name, nk.Config.Namespace)
	if err != nil {
		return err
	}
	if force {
		fmt.Fprintf(out, "Crashing machine %s\n", name)
	} else {
		fmt.Fprintf(out, "Halting machine %s\n", name)
	}
	err = m.Stop(force)
	if err != nil {
		return err
	}
	if waitTimeout := nk.Config.Wait; waitTimeout > 0 {
		return m.WaitUntil(waitTimeout, driver.ExitedState(), nil)
	}
	return nil
}

func (nk *Koble) RemoveMachine(name string, out io.Writer) error {
	m, err := nk.Driver.Machine(name, nk.Config.Namespace)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "removing machine %s\n", name)
	return m.Remove()
}

func (nk *Koble) DestroyMachine(machine string, out io.Writer) error {
	m, err := nk.Driver.Machine(machine, nk.Config.Namespace)
	if err != nil {
		return err
	}
	exists, err := m.Exists()
	if err != nil {
		return err
	}
	if !exists {
		fmt.Fprintf(out, "no machine to remove\n")
		return nil
	}
	fmt.Fprintf(out, "Crashing machine %s\n", m.Name())
	err = m.Stop(true)
	if err != nil {
		return err
	}
	err = m.WaitUntil(120, driver.ExitedState(), nil)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Removing machine %s\n", m.Name())
	return m.Remove()
}

func (nk *Koble) MachineLogs(machine string, follow bool, tail int) error {
	m, err := nk.Driver.Machine(machine, nk.Config.Namespace)
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
	machines, err := nk.Driver.ListMachines(nk.Config.Namespace, all)
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
