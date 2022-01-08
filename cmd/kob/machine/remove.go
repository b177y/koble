package machine

import (
	"os"

	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:                   "remove [options] MACHINE",
	Short:                 "remove a koble machine",
	Aliases:               []string{"rm"},
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     kob.AutocompNonRunningMachine,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.RemoveMachine(args[0], os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	machineCmd.AddCommand(removeCmd)
	kob.RootCmd.AddCommand(removeCmd)
}
