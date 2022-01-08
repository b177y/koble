package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [options] MACHINE [MACHINE...]",
	Aliases: []string{"rm"},
	Short:   "remove machines in a lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.LabRemove(args)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	labCmd.AddCommand(removeCmd)
}
