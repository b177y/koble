package machine

import (
	"errors"

	"github.com/b177y/koble/cmd/kob/cli"

	"github.com/spf13/cobra"
)

var user string
var workDir string

var shellCmd = &cobra.Command{
	Use:               "shell [options] MACHINE [COMMAND [ARG...]]",
	Short:             "get a shell on a machine",
	ValidArgsFunction: cli.AutocompRunningMachine,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if useTerm && useCon {
			return errors.New("CLI Flags --terminal and --console cannot be used together.")
		} else if (useTerm && detachMode) || (useCon && detachMode) {
			return errors.New("CLI Flag --detach cannot be used with --terminal or --console.")
		} else if useTerm {
			cli.NK.Config.OpenTerms = true
		} else if useCon {
			cli.NK.Config.OpenTerms = false
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if cli.NK.Config.OpenTerms {
			return cli.NK.LaunchInTerm()
		}
		return cli.NK.Shell(args[0], user, workDir)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	shellCmd.Flags().StringVarP(&user, "user", "u", "", "User to execute shell as.")
	shellCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	shellCmd.Flags().BoolVarP(&useTerm, "terminal", "t", false, "Launch shell in new terminal.")
	shellCmd.Flags().BoolVar(&useCon, "console", false, "Launch shell within current console.")
	machineCmd.AddCommand(shellCmd)
	cli.Commands = append(cli.Commands, shellCmd)
}
