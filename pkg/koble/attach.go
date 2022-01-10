package koble

import "github.com/b177y/koble/driver"

func (nk *Koble) AttachToMachine(machine string) error {
	m, err := nk.Driver.Machine(machine, nk.Config.Namespace)
	if err != nil {
		return err
	}
	return m.Attach(nil)
}

func (nk *Koble) Exec(machine, command, user string,
	detach bool, workdir string) error {
	m, err := nk.Driver.Machine(machine, nk.Config.Namespace)
	if err != nil {
		return err
	}
	return m.Exec(command, &driver.ExecOptions{
		User:    user,
		Detach:  detach,
		Workdir: workdir,
	})
}

func (nk *Koble) Shell(machine, user, workdir string) error {
	m, err := nk.Driver.Machine(machine, nk.Config.Namespace)
	if err != nil {
		return err
	}
	return m.Shell(&driver.ShellOptions{
		User:    user,
		Workdir: workdir,
	})
}
