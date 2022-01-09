package machine

import (
	"errors"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var (
	useTerm bool
	useCon  bool
)

var attachCmd = &cobra.Command{
	Use:               "attach MACHINE [options]",
	Short:             "attach to the main tty of a machine",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: cli.AutocompRunningMachine,
	Example: `koble attach a0 --terminal
koble attach dh --console`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if useTerm && useCon {
			return errors.New("CLI Flags --terminal and --console cannot be used together.")
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
		machine := args[0]
		return cli.NK.AttachToMachine(machine)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	attachCmd.Flags().BoolVarP(&useTerm, "terminal", "t", false, "Launch shell in new terminal.")
	attachCmd.Flags().BoolVar(&useCon, "console", false, "Launch shell within current console.")
	machineCmd.AddCommand(attachCmd)
	cli.Commands = append(cli.Commands, attachCmd)
}
