package cmd

import (
	"errors"
	"os"

	"github.com/b177y/netkit/pkg/netkit"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var command string
var user string
var detachMode bool
var workDir string

var shellCmd = &cobra.Command{
	Use:   "shell [MACHINE]",
	Short: "The 'shell' subcommand is used to connect to a shell on a netkit machine",
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
		// TODO get driver from ?cmd?
		if nk.Config.OpenTerms {
			err := netkit.LaunchInTerm()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}
		machine := args[0]
		err := nk.ExecMachineShell(machine, command, user, detachMode, workDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	shellCmd.Flags().StringVarP(&command, "command", "c", "/bin/bash", "Command to execute in shell.")
	shellCmd.Flags().StringVarP(&user, "user", "u", "", "User to execute shell as.")
	shellCmd.Flags().BoolVarP(&detachMode, "detach", "d", false, "Run the exec session in detached mode (backgrounded)")
	shellCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	shellCmd.Flags().BoolVarP(&useTerm, "terminal", "t", false, "Launch shell in new terminal.")
	shellCmd.Flags().BoolVar(&noTerm, "console", false, "Launch shell within current console.")
}
