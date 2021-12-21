package netkit

import "github.com/b177y/netkit/driver"

func (nk *Netkit) AttachToMachine(machine string) error {
	m := driver.Machine{
		Name:      machine,
		Lab:       nk.Lab.Name,
		Namespace: nk.Namespace,
	}
	return nk.Driver.AttachToMachine(m)
}

func (nk *Netkit) Exec(machine, command, user string,
	detach bool, workdir string) error {
	m := driver.Machine{
		Name:      machine,
		Lab:       nk.Lab.Name,
		Namespace: nk.Namespace,
	}
	return nk.Driver.Exec(m, command, user, detach, workdir)
}

func (nk *Netkit) Shell(machine, user, workdir string) error {
	m := driver.Machine{
		Name:      machine,
		Lab:       nk.Lab.Name,
		Namespace: nk.Namespace,
	}
	return nk.Driver.Shell(m, user, workdir)
}
