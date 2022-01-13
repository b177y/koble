package machine

import (
	"fmt"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/driver"
	"github.com/b177y/koble/pkg/output"
	"github.com/spf13/cobra"
)

var startOpts driver.MachineConfig

var startCmd = &cobra.Command{
	Use:                   "start [options] MACHINE",
	Short:                 "start a machine",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cli.AutocompNonRunningMachine,
	RunE:                  start,
}

func init() {
	startCmd.Flags().StringVar(&startOpts.Image, "image", "", "image to run machine with")
	startCmd.Flags().StringArrayVar(&startOpts.Networks, "network", []string{}, "networks to attach to machine")
	cli.AddWaitFlag(startCmd)
	machineCmd.AddCommand(startCmd)
	cli.Commands = append(cli.Commands, startCmd)
}

var start = func(cmd *cobra.Command, args []string) error {
	return output.WithSimpleContainer(
		fmt.Sprintf("Starting machine %s", args[0]),
		nil,
		cli.NK.Config.NonInteractive,
		func(c output.Container, out output.Output) (err error) {
			err = cli.NK.StartMachine(args[0], startOpts, out)
			if err == nil {
				out.Success("Started machine " + args[0])
			}
			return err
		})
}
