package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var detachMode bool

var execCmd = &cobra.Command{
	Use:               "exec [options] MACHINE [COMMAND [ARG...]]",
	Short:             "run a command on a koble machine",
	ValidArgsFunction: autocompRunningMachine,
	Run: func(cmd *cobra.Command, args []string) {
		if nk.Config.OpenTerms {
			err := nk.LaunchInTerm()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
		err := nk.Exec(args[0], strings.Join(args[1:], " "), user, detachMode, workDir)
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
	KobleCLI.AddCommand(execCmd)
	machineCmd.AddCommand(execCmd)
}
