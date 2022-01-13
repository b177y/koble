package machine

import (
	"strings"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/koble"

	"github.com/spf13/cobra"
)

var detachMode bool

var execCmd = &cobra.Command{
	Use:               "exec [options] MACHINE [COMMAND [ARG...]]",
	Short:             "run a command on a machine",
	ValidArgsFunction: cli.AutocompRunningMachine,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cli.NK.Config.Terminal.Launch {
			return cli.NK.LaunchInTerm(args[0])
		}
		return cli.NK.Exec(args[0], strings.Join(args[1:], " "), user, detachMode, workDir)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	execCmd.Flags().SetInterspersed(false)
	execCmd.Flags().StringVarP(&user, "user", "u", "", "User to execute shell as.")
	execCmd.Flags().BoolVarP(&detachMode, "detach", "d", false, "Run the command in detached mode (backgrounded)")
	execCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	execCmd.Flags().StringP("terminal", "t", "", "terminal to launch")
	koble.BindFlag("terminal.name", execCmd.Flags().Lookup("terminal"))
	execCmd.Flags().BoolVar(&launch, "launch", false, "launch terminal for attach session")
	koble.BindFlag("terminal.launch", execCmd.Flags().Lookup("launch"))
	execCmd.Flags().StringToString("term-opt", map[string]string{}, "option to pass to terminal")
	koble.BindFlag("term_opts", execCmd.Flags().Lookup("term-opt"))
	machineCmd.AddCommand(execCmd)
	cli.Commands = append(cli.Commands, execCmd)
}
