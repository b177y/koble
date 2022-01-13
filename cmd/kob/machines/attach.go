package machine

import (
	"fmt"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var launch bool
var terminal string

var attachCmd = &cobra.Command{
	Use:               "attach MACHINE [options]",
	Short:             "attach to the main tty of a machine",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: cli.AutocompRunningMachine,
	Example: `koble attach a0 --terminal
koble attach dh --console`,
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := args[0]
		fmt.Println("opts", cli.NK.Config.Terminal.Launch)
		fmt.Println("cob", cmd.Flags().Lookup("launch").Value)
		if cli.NK.Config.Terminal.Launch {
			return cli.NK.LaunchInTerm(machine)
		}
		return cli.NK.AttachToMachine(machine)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	attachCmd.Flags().StringVarP(&terminal, "terminal", "t", "gnome", "terminal to launch")
	koble.BindFlag("terminal.name", attachCmd.Flags().Lookup("terminal"))
	attachCmd.Flags().BoolVar(&launch, "launch", false, "launch terminal for attach session")
	koble.BindFlag("terminal.launch", attachCmd.Flags().Lookup("launch"))
	attachCmd.Flags().StringToString("term-opt", map[string]string{}, "option to pass to terminal")
	koble.BindFlag("term_opts", attachCmd.Flags().Lookup("term-opt"))
	machineCmd.AddCommand(attachCmd)
	cli.Commands = append(cli.Commands, attachCmd)
}
