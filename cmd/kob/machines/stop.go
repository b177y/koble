package machine

import (
	"fmt"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/driver"
	"github.com/b177y/koble/pkg/output"
	"github.com/spf13/cobra"
)

var forceStop bool

var stopCmd = &cobra.Command{
	Use:                   "stop [options] MACHINE",
	Short:                 "stop a machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompRunningMachine,
	DisableFlagsInUseLine: true,
	RunE:                  stop,
}

func init() {
	stopCmd.Flags().BoolVarP(&wait, "wait", "w", false, "wait for machine to stop")
	stopCmd.Flags().BoolVarP(&forceStop, "force", "f", false, "force stop machine")
	machineCmd.AddCommand(stopCmd)
	cli.Commands = append(cli.Commands, stopCmd)
}

var stop = func(cmd *cobra.Command, args []string) error {
	return output.WithSimpleContainer(
		fmt.Sprintf("Stopping machine %s", args[0]),
		nil,
		cli.NK.Config.NonInteractive,
		func(c output.Container, out output.Output) (err error) {
			defer func() {
				if err == nil {
					out.Success("Stopped machine " + args[0])
				}
			}()
			err = cli.NK.StopMachine(args[0], forceStop, out)
			if err != nil {
				return err
			}
			if wait {
				m, err := cli.NK.Driver.Machine(args[0], cli.NK.Config.Namespace)
				if err != nil {
					return err
				}
				return m.WaitUntil(60*5, driver.ExitedState(), nil)
			}
			return nil
		})
}
