package machine

import (
	"fmt"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/driver"
	"github.com/b177y/koble/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	startCmd.Flags().Int("wait", 300, "seconds to wait for machine to boot before timeout (default 300, -1 is don't wait)")
	viper.BindPFlag("wait", startCmd.Flags().Lookup("wait"))

	machineCmd.AddCommand(startCmd)
	cli.Commands = append(cli.Commands, startCmd)
}

var start = func(cmd *cobra.Command, args []string) error {
	return output.WithSimpleContainer(
		fmt.Sprintf("Starting machine %s", args[0]),
		nil,
		cli.NK.Config.NonInteractive,
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
			if cli.NK.Config.Wait > 0 {
				m, err := cli.NK.Driver.Machine(args[0], cli.NK.Config.Namespace)
				if err != nil {
					return err
				}
				fmt.Fprintf(out, "booting")
				return m.WaitUntil(cli.NK.Config.Wait, driver.BootedState(), driver.ExitedState())
			}
			return nil

		})
}
