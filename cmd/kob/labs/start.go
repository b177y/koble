package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"

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
	labCmd.AddCommand(startCmd)
}
