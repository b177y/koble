package machine

import (
	"os"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:                   "remove [options] MACHINE",
	Short:                 "remove a koble machine",
	Aliases:               []string{"rm"},
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompNonRunningMachine,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.RemoveMachine(args[0], os.Stdout)
	},
}

func init() {
	machineCmd.AddCommand(removeCmd)
	cli.Commands = append(cli.Commands, removeCmd)
}
