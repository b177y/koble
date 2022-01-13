package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var launch bool
var terminal string

var attachCmd = &cobra.Command{
	Use:               "attach MACHINE [options]",
	Short:             "attach to the main tty of a machine",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: cli.AutocompRunningMachine,
	Example: `koble attach a0 --terminal xterm
koble attach dh --launch=false`,
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := args[0]
		if cli.NK.Config.Terminal.Launch {
			return cli.NK.LaunchInTerm(machine)
		}
		return cli.NK.AttachToMachine(machine)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	cli.AddTermFlags(attachCmd)
	cli.AddWaitFlag(attachCmd)
	machineCmd.AddCommand(attachCmd)
	cli.Commands = append(cli.Commands, attachCmd)
}
