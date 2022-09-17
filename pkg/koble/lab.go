package koble

import (
	"errors"
	"fmt"
	"sync"

	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/output"
)

func (nk *Koble) LabStart(mlist []string) error {
	return nk.ForMachine(nk.Lab.Header, "Starting", mlist, func(name string,
		mconf driver.MachineConfig,
		out output.Output) (err error) {
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Started machine %s", name))
			}
		}()
		out.Start()
		attachTerm := ""
		if nk.Config.Launch.LabStart {
			attachTerm = nk.Config.Terminal.LabStart
		}
		return nk.StartMachine(name, mconf, attachTerm, out)
	})
}

func filterMachines(machines map[string]driver.MachineConfig,
	filter []string) (mList map[string]driver.MachineConfig) {
	// if no machines in filter then all machines are included
	if len(filter) == 0 {
		return machines
	}
	// only keep machines which are in the filter list
	for _, name := range filter {
		if val, ok := machines[name]; ok {
			mList[name] = val
		}
	}
	return mList
}

func (nk *Koble) GetMachineList(mlist []string,
	all bool) (machines []driver.Machine, err error) {
	if len(mlist) == 0 && nk.Lab.Name == "" && !all {
		return machines, errors.New("You are not in a lab. Use --all or specify machines.")
	} else if all && len(mlist) != 0 {
		return machines, errors.New("You cannot specify machines when using --all")
	}
	// output, err := nk.Driver.ListMachines(nk.Lab.Name, all)
	// for _, m := range output {
	// 	machines = append(machines, driver.Machine{
	// 		Name: m.Name,
	// 		Lab:  m.Lab,
	// 	})
	// }
	// machines = filterMachines(machines, mlist)
	return machines, nil
}

func (nk *Koble) LabDestroy(mlist []string) error {
	return nk.ForMachine(nk.Lab.Header, "Destroying", mlist, func(name string,
		mconf driver.MachineConfig,
		out output.Output) (err error) {
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Destroyed machine %s", name))
			}
		}()
		out.Start()
		return nk.DestroyMachine(name, out)
	})
}

func (nk *Koble) LabRemove(mlist []string) error {
	return nk.ForMachine(nk.Lab.Header, "Removing", mlist, func(name string,
		mconf driver.MachineConfig,
		out output.Output) (err error) {
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Removed machine %s", name))
			}
		}()
		out.Start()
		return nk.RemoveMachine(name, out)
	})
}

func (nk *Koble) LabStop(mlist []string, force bool) error {
	return nk.ForMachine(nk.Lab.Header, "stopping", mlist, func(name string,
		mconf driver.MachineConfig,
		out output.Output) (err error) {
		defer func() {
			if err != nil {
				out.Error(err)
			} else {
				out.Success(fmt.Sprintf("Stopped machine %s", name))
			}
		}()
		out.Start()
		return nk.StopMachine(name, force, out)
	})
}

func (nk *Koble) LabInfo() error {
	if nk.LabRoot == "" {
		return errors.New("You are not currently in a lab directory.")
	}
	err := nk.ListMachines(false, false)
	if err != nil {
		return err
	}
	fmt.Printf("\n")
	err = nk.ListNetworks(false, false)
	fmt.Printf("\n")
	return err
}

func (nk *Koble) ForMachine(headerFunc func(string) string, titlePref string, filterList []string, toRun func(name string, mconf driver.MachineConfig, out output.Output) error) error {
	if nk.LabRoot == "" {
		return errors.New("You are not currently in a lab directory.")
	}

	oc := output.NewContainer(headerFunc, titlePref, nk.Config.NonInteractive)
	oc.Start()
	defer oc.Stop()
	machines := filterMachines(nk.Lab.Machines, filterList)
	var wg sync.WaitGroup
	for name, mconf := range machines {
		wg.Add(1)
		go func(name string, mconf driver.MachineConfig, c output.Container) error {
			defer wg.Done()
			title := fmt.Sprintf("%s machine %s", titlePref, name)
			return output.WithStdout(title, oc, func(out output.Output) error {
				return toRun(name, mconf, out)
			})
		}(name, mconf, oc)
	}
	wg.Wait()
	return nil
}
