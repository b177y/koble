package netkit

import "github.com/b177y/netkit/driver"

func (nk *Netkit) AttachToMachine(machine string) error {
	m, err := nk.Driver.Machine(machine)
	if err != nil {
		return err
	}
	return m.Attach(nil)
}

func (nk *Netkit) Exec(machine, command, user string,
	detach bool, workdir string) error {
	m, err := nk.Driver.Machine(machine)
	if err != nil {
		return err
	}
	return m.Exec(command, &driver.ExecOptions{
		User:    user,
		Detach:  detach,
		Workdir: workdir,
	})
}

func (nk *Netkit) Shell(machine, user, workdir string) error {
	m, err := nk.Driver.Machine(machine)
	if err != nil {
		return err
	}
	return m.Shell(&driver.ShellOptions{
		User:    user,
		Workdir: workdir,
	})
}
