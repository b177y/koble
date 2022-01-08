package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"

	"github.com/spf13/cobra"
)

var labHaltForce bool
var labAllMachines bool

var lstartCmd = &cobra.Command{
	Use:   "start [options] MACHINE [MACHINE...]",
	Short: "start a koble lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.Lab.Start(args)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	labCmd.AddCommand(lstartCmd)
}
