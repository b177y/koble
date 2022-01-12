package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var attachCmd = &cobra.Command{
	Use:               "attach MACHINE [options]",
	Short:             "attach to the main tty of a machine",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: cli.AutocompRunningMachine,
	Example: `koble attach a0 --terminal
koble attach dh --console`,
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := args[0]
		if cli.NK.Config.Terminal.Launch {
			return cli.NK.LaunchInTerm(machine)
		}
		return cli.NK.AttachToMachine(machine)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	attachCmd.Flags().StringP("terminal", "t", "", "terminal to launch")
	viper.BindPFlag("terminal.name", attachCmd.Flags().Lookup("terminal"))
	attachCmd.Flags().Bool("launch", false, "launch terminal for attach session")
	viper.BindPFlag("terminal.launch", attachCmd.Flags().Lookup("launch"))
	attachCmd.Flags().StringToString("term-opt", map[string]string{}, "option to pass to terminal")
	viper.BindPFlag("term_opts", attachCmd.Flags().Lookup("term-opt"))
	machineCmd.AddCommand(attachCmd)
	cli.Commands = append(cli.Commands, attachCmd)
}
