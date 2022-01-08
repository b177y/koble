package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy [options] MACHINE [MACHINE...]",
	Short: "crash and remove machines in a lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.LabDestroy(args)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	labCmd.AddCommand(destroyCmd)
}
