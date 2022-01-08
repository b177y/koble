package machine

import (
	"os"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:                   "destroy [options] MACHINE",
	Short:                 "force stop and remove a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompMachine,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.DestroyMachine(args[0], os.Stdout)
	},
}

func init() {
	machineCmd.AddCommand(destroyCmd)
	cli.Commands = append(cli.Commands, destroyCmd)
}
