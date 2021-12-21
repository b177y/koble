package cmd

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var user string
var workDir string

var shellCmd = &cobra.Command{
	Use:               "shell [options] MACHINE [COMMAND [ARG...]]",
	Short:             "The 'shell' subcommand is used to connect to a shell on a netkit machine",
	ValidArgsFunction: autocompRunningMachine,
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
}
