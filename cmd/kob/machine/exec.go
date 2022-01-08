package machine

import (
	"os"
	"strings"

	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var detachMode bool

var execCmd = &cobra.Command{
	Use:               "exec [options] MACHINE [COMMAND [ARG...]]",
	Short:             "run a command on a koble machine",
	ValidArgsFunction: kob.AutocompRunningMachine,
	Run: func(cmd *cobra.Command, args []string) {
		if kob.NK.Config.OpenTerms {
			err := kob.NK.LaunchInTerm()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
		err := kob.NK.Exec(args[0], strings.Join(args[1:], " "), user, detachMode, workDir)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	execCmd.Flags().SetInterspersed(false)
	execCmd.Flags().StringVarP(&user, "user", "u", "", "User to execute shell as.")
	execCmd.Flags().BoolVarP(&detachMode, "detach", "d", false, "Run the command in detached mode (backgrounded)")
	execCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	machineCmd.AddCommand(execCmd)
	kob.RootCmd.AddCommand(execCmd)
}
