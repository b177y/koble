package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var stopForce bool

var stopCmd = &cobra.Command{
	Use:   "stop [options] MACHINE [MACHINE...]",
	Short: "stop a lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.LabStop(args, stopForce, wait)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	stopCmd.Flags().BoolVarP(&stopForce, "force", "f", false, "Force halt machines.")
	stopCmd.Flags().BoolVarP(&wait, "wait", "w", false, "wait for all machines to shutdown")
	labCmd.AddCommand(stopCmd)
}