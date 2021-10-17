package cmd

import (
	"errors"
	"os"

	"github.com/b177y/netkit/pkg/netkit"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach [MACHINE]",
	Short: "The 'attach' subcommand is used to attach to the main tty on a netkit machine",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if useTerm && noTerm {
			err := errors.New("CLI Flags --terminal and --console cannot be used together.")
			log.Fatal(err)
		} else if useTerm {
			nk.Config.OpenTerms = true
		} else if noTerm {
			nk.Config.OpenTerms = false
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if nk.Config.OpenTerms {
			err := netkit.LaunchInTerm()
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
}

func init() {
	attachCmd.Flags().BoolVarP(&useTerm, "terminal", "t", false, "Launch shell in new terminal.")
	attachCmd.Flags().BoolVar(&noTerm, "console", false, "Launch shell within current console.")
}
