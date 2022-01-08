package machine

import (
	"os"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/driver"
	"github.com/spf13/cobra"
)

var startOpts driver.MachineConfig

var startCmd = &cobra.Command{
	Use:                   "start [options] MACHINENAME",
	Short:                 "start a koble machine",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cli.AutocompNonRunningMachine,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.StartMachine(args[0], startOpts, os.Stdout) // TODO put plain into nk
	},
}

func init() {
	startCmd.Flags().StringVar(&startOpts.Image, "image", "", "Image to run machine with.")
	startCmd.Flags().StringArrayVar(&startOpts.Networks, "networks", []string{}, "Networks to attach to machine")
	startCmd.Flags().BoolVarP(&machineWait, "wait", "w", false, "wait for machine to boot up")

	machineCmd.AddCommand(startCmd)
	cli.Commands = append(cli.Commands, startCmd)
}
