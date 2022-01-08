package machine

import (
	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:                   "info [options] MACHINE",
	Short:                 "get info about a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     kob.AutocompMachine,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.MachineInfo(args[0], mInfoJson)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	infoCmd.Flags().BoolVar(&mInfoJson, "json", false, "Print machine info as json object to stdout")
	machineCmd.AddCommand(infoCmd)
	kob.RootCmd.AddCommand(infoCmd)
}
