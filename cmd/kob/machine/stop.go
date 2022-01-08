package machine

import (
	"os"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:                   "stop [options] MACHINE",
	Short:                 "stop a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompRunningMachine,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.HaltMachine(args[0], false, os.Stdout)
	},
}

func init() {
	machineCmd.AddCommand(stopCmd)
	cli.Commands = append(cli.Commands, stopCmd)
}
