package cmd

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:               "attach [options] MACHINE",
	Short:             "The 'attach' subcommand is used to attach to the main tty on a netkit machine",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: autocompMachine,
	PreRun: func(cmd *cobra.Command, args []string) {
		if useTerm && useCon {
			err := errors.New("CLI Flags --terminal and --console cannot be used together.")
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
		machine := args[0]
		err := nk.AttachToMachine(machine)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	attachCmd.Flags().BoolVarP(&useTerm, "terminal", "t", false, "Launch shell in new terminal.")
	attachCmd.Flags().BoolVar(&useCon, "console", false, "Launch shell within current console.")
}
