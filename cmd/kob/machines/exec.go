package machine

import (
	"strings"

	"github.com/b177y/koble/cmd/kob/cli"

	"github.com/spf13/cobra"
)

var detachMode bool

var execCmd = &cobra.Command{
	Use:               "exec [options] MACHINE [COMMAND [ARG...]]",
	Short:             "run a command on a machine",
	ValidArgsFunction: cli.AutocompRunningMachine,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cli.NK.Config.OpenTerms {
			return cli.NK.LaunchInTerm()
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
	machineCmd.AddCommand(execCmd)
	cli.Commands = append(cli.Commands, execCmd)
}
