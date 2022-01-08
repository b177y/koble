package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy [options] MACHINE [MACHINE...]",
	Short: "crash and remove all machines in a koble lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.LabDestroy(args, labAllMachines)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	destroyCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Destroy all koble machines, including those not in the current lab.")
	labCmd.AddCommand(destroyCmd)
}
