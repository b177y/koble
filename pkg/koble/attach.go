package koble

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/output"
	log "github.com/sirupsen/logrus"
)

func (nk *Koble) AttachToMachine(machine, term string) error {
	if os.Getenv("_KOBLE_IN_TERM") != "true" && term != "this" {
		log.WithFields(log.Fields{"machine": machine}).
			Debug("attach not in terminal, relaunching now")
		command, err := nk.reexecAttach(machine)
		if err != nil {
			return err
		}
		return nk.LaunchInTerm(machine, nk.Config.Terminal.Attach, command)
	}
	m, err := nk.Driver.Machine(machine, nk.Config.Namespace)
	if err != nil {
		return err
	}
	if waitTimeout := nk.Config.Wait; waitTimeout > 0 {
		err := output.WithSimpleContainer(fmt.Sprintf("attaching to %s", machine),
			nil, nk.Config.NonInteractive, func(o output.Output) error {
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
				o.Success("attached to machine " + machine)
				return nil
			})
		state, err := m.State()
		if err != nil {
			return err
		}
		if state.State != nil {
			if *state.State != "running" {
				err = m.WaitUntil(5*time.Second, driver.BootingState(),
					driver.ExitedState())
				if err != nil {
					return err
				}
			}
		}
	}
	err = nk.MachineLogs(m.Name(), false, -1)
	if err != nil {
		return err
	}
	return m.Attach(nil)
}

func (nk *Koble) Exec(machine, command, user string,
	detach bool, workdir string) error {
	if os.Getenv("_KOBLE_IN_TERM") != "true" && nk.Config.Terminal.Exec != "this" {
		log.WithFields(log.Fields{"machine": machine}).
			Debug("exec not in terminal, relaunching now")
		command, err := nk.reexecExec(machine, command, user, detach, workdir)
		if err != nil {
			return err
		}
		return nk.LaunchInTerm(machine, nk.Config.Terminal.Exec, command)
	}
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
	if os.Getenv("_KOBLE_IN_TERM") != "true" && nk.Config.Terminal.Shell != "this" {
		log.WithFields(log.Fields{"machine": machine}).
			Debug("shell not in terminal, relaunching now")
		command, err := nk.reexecShell(machine, user, workdir)
		if err != nil {
			return err
		}
		return nk.LaunchInTerm(machine, nk.Config.Terminal.Shell, command)
	}
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
