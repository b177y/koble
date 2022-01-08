package lab

import (
	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy [options] MACHINE [MACHINE...]",
	Short: "crash and remove all machines in a koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.LabDestroy(args, labAllMachines)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	destroyCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Destroy all koble machines, including those not in the current lab.")
	labCmd.AddCommand(destroyCmd)
}
