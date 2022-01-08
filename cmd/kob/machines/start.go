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
	Use:                   "start [options] MACHINENAME",
	Short:                 "start a machine",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cli.AutocompNonRunningMachine,
	RunE:                  start,
}

func init() {
	startCmd.Flags().StringVar(&startOpts.Image, "image", "", "Image to run machine with.")
	startCmd.Flags().StringArrayVar(&startOpts.Networks, "networks", []string{}, "Networks to attach to machine")
	startCmd.Flags().BoolVarP(&wait, "wait", "w", false, "wait for machine to boot")

	machineCmd.AddCommand(startCmd)
	cli.Commands = append(cli.Commands, startCmd)
}

var start = func(cmd *cobra.Command, args []string) error {
	return output.WithSimpleContainer(
		fmt.Sprintf("Starting machine %s", args[0]),
		nil,
		cli.Plain,
		func(c output.Container, out output.Output) (err error) {
			defer func() {
				if err == nil {
					out.Success("Started machine " + args[0])
				}
			}()
			err = cli.NK.StartMachine(args[0], startOpts, out)
			if err != nil {
				return err
			}
			if wait {
				m, err := cli.NK.Driver.Machine(args[0], cli.NK.Namespace)
				if err != nil {
					return err
				}
				fmt.Fprintf(out, "booting")
				return m.WaitUntil("running", 60*5)
			}
			return nil

		})
}
