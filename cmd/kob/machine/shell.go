package machine

import (
	"errors"
	"os"

	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var user string
var workDir string

var shellCmd = &cobra.Command{
	Use:               "shell [options] MACHINE [COMMAND [ARG...]]",
	Short:             "get a shell on a koble machine",
	ValidArgsFunction: kob.AutocompRunningMachine,
	PreRun: func(cmd *cobra.Command, args []string) {
		if useTerm && useCon {
			err := errors.New("CLI Flags --terminal and --console cannot be used together.")
			log.Fatal(err)
		} else if (useTerm && detachMode) || (useCon && detachMode) {
			err := errors.New("CLI Flag --detach cannot be used with --terminal or --console.")
			log.Fatal(err)
		} else if useTerm {
			nk.Config.OpenTerms = true
		} else if useCon {
			nk.Config.OpenTerms = false
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if nk.Config.OpenTerms {
			err := nk.LaunchInTerm()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
		err := nk.Shell(args[0], user, workDir)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	shellCmd.Flags().StringVarP(&user, "user", "u", "", "User to execute shell as.")
	shellCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	shellCmd.Flags().BoolVarP(&useTerm, "terminal", "t", false, "Launch shell in new terminal.")
	shellCmd.Flags().BoolVar(&useCon, "console", false, "Launch shell within current console.")
	KobleCLI.AddCommand(shellCmd)
	machineCmd.AddCommand(shellCmd)
}
