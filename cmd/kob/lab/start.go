package lab

import (
	log "github.com/sirupsen/logrus"

	"github.com/b177y/koble/cmd/kob"
	"github.com/spf13/cobra"
)

var labHaltForce bool
var labAllMachines bool

var lstartCmd = &cobra.Command{
	Use:   "start [options] MACHINE [MACHINE...]",
	Short: "start a koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.LabStart(args)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	labCmd.AddCommand(lstartCmd)
}
