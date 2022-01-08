package machine

import (
	"os"

	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:                   "destroy [options] MACHINE",
	Short:                 "force stop and remove a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     AutocompMachine,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.DestroyMachine(args[0], os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	machineCmd.AddCommand(destroyCmd)
	kob.RootCmd.AddCommand(destroyCmd)
}
