package koble

import (
	"context"
	"fmt"
	"time"

	"github.com/b177y/koble/driver"
)

func (nk *Koble) AttachToMachine(machine string) error {
	m, err := nk.Driver.Machine(machine, nk.Config.Namespace)
	if err != nil {
		return err
	}
	if waitTimeout := nk.Config.Wait; waitTimeout > 0 {
		fmt.Printf("waiting for %s to be created\r", machine)
		ctx, cancel := context.WithTimeout(context.Background(),
			waitTimeout*time.Second)
		defer cancel()
		for {
			time.Sleep(200 * time.Millisecond)
			// check timeout
			if err := ctx.Err(); err != nil {
				return fmt.Errorf("timed out waiting for %s to be created: %w",
					m.Name(), err)
			}
			exists, err := m.Exists()
			if err != nil {
				return err
			}
			if exists {
				break
			}
		}
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
	fmt.Printf("waiting for %s to boot\r", machine)
	if waitTimeout := nk.Config.Wait; waitTimeout > 0 {
		err = m.WaitUntil(waitTimeout, driver.BootedState(), driver.ExitedState())
		if err != nil {
			return err
		}
	}
	fmt.Printf("                         \r")
	return m.Shell(&driver.ShellOptions{
		User:    user,
		Workdir: workdir,
	})
}
