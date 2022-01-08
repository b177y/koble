package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [options] MACHINE [MACHINE...]",
	Short: "stop machines in a koble lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.LabHalt(args, labHaltForce, labAllMachines)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	stopCmd.Flags().BoolVarP(&labHaltForce, "force", "f", false, "Force halt machines.")
	stopCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Halt all koble machines, including those not in the current lab.")
	labCmd.AddCommand(stopCmd)
}
