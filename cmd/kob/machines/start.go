package machine

import (
	"fmt"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/koble"
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
	cli.AddTermFlags(startCmd, "machine_start")
	cli.AddWaitFlag(startCmd)
	startCmd.Flags().Bool("launch", false, "launch attach session to machine in terminal")
	koble.BindFlag("launch.machine_start", startCmd.Flags().Lookup("launch"))
	machineCmd.AddCommand(startCmd)
	cli.Commands = append(cli.Commands, startCmd)
}

var start = func(cmd *cobra.Command, args []string) error {
	return output.WithSimpleContainer(
		fmt.Sprintf("Starting machine %s", args[0]),
		nil,
		cli.NK.Config.NonInteractive,
		func(out output.Output) (err error) {
			if cli.NK.Config.Launch.MachineStart {
				err := cli.NK.AttachToMachine(args[0], cli.NK.Config.Terminal.MachineStart)
				if err != nil {
					return err
				}
			}
			err = cli.NK.StartMachine(args[0], startOpts)
			if err != nil {
				return err
			}
			out.Success("Started machine " + args[0])
			return nil
		})
}
