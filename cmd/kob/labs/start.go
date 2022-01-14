package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/koble"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [options] MACHINE [MACHINE...]",
	Short: "start a lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.LabStart(args)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	cli.AddWaitFlag(startCmd)
	cli.AddTermFlags(startCmd, "lab_start")
	startCmd.Flags().Bool("launch", false, "launch terminal attach sessions to started lab machines (conflicts with terminal 'this')")
	koble.BindFlag("launch.lab_start", startCmd.Flags().Lookup("launch"))
	labCmd.AddCommand(startCmd)
}
