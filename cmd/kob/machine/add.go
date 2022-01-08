package machine

import (
	"github.com/b177y/koble/cmd/kob"
	"github.com/b177y/koble/pkg/koble"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:                   "add [options] MACHINENAME",
	Short:                 "add a new machine to a lab",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := koble.AddMachineToLab(args[0], machineNetworks, machineImage)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	addCmd.Flags().StringVar(&addMachineImage, "image", "", "Image to use for new machine.")
	addCmd.Flags().StringArrayVar(&addMachineNetworks, "networks", []string{}, "Networks to add to new machine.")
	machineCmd.AddCommand(addCmd)
	kob.RootCmd.AddCommand(addCmd)
}
