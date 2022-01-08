package machine

import (
	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:                   "start [options] MACHINENAME",
	Short:                 "start a koble machine",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     kob.AutocompNonRunningMachine,
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.StartMachineWithStatus(args[0], machineImage, machineNetworks, machineWait, false) // TODO put plain into nk
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	startCmd.Flags().StringVar(&machineImage, "image", "", "Image to run machine with.")
	startCmd.Flags().StringArrayVar(&machineNetworks, "networks", []string{}, "Networks to attach to machine")
	startCmd.Flags().BoolVarP(&machineWait, "wait", "w", false, "wait for machine to boot up")

	machineCmd.AddCommand(startCmd)
	kob.RootCmd.AddCommand(startCmd)
}
