package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var stopForce bool

var stopCmd = &cobra.Command{
	Use:   "stop [options] MACHINE [MACHINE...]",
	Short: "stop machines in a koble lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.LabStop(args, stopForce, labAllMachines)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	stopCmd.Flags().BoolVarP(&stopForce, "force", "f", false, "Force halt machines.")
	labCmd.AddCommand(stopCmd)
}
