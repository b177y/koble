package netkit

func (nk *Netkit) AttachToMachine(machine string) error {
	err := nk.Driver.AttachToMachine(machine, nk.Lab.Name)
	return err
}

func (nk *Netkit) ExecMachineShell(machine, command, user string,
	detach bool, workdir string) error {
	err := nk.Driver.MachineExecShell(machine, nk.Lab.Name, command, user, detach, workdir)
	return err
}
