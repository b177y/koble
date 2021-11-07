package netkit

import "github.com/b177y/netkit/driver"

func (nk *Netkit) AttachToMachine(machine string) error {
	m := driver.Machine{
		Name: machine,
		Lab:  nk.Lab.Name,
	}
	err := nk.Driver.AttachToMachine(m)
	return err
}

func (nk *Netkit) ExecMachineShell(machine, command, user string,
	detach bool, workdir string) error {
	m := driver.Machine{
		Name: machine,
		Lab:  nk.Lab.Name,
	}
	err := nk.Driver.MachineExecShell(m, command, user, detach, workdir)
	return err
}
