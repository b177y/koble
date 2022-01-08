package lab

import (
	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [options] MACHINE [MACHINE...]",
	Short: "stop machines in a koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.LabHalt(args, labHaltForce, labAllMachines)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	stopCmd.Flags().BoolVarP(&labHaltForce, "force", "f", false, "Force halt machines.")
	stopCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Halt all koble machines, including those not in the current lab.")
	labCmd.AddCommand(stopCmd)
}
