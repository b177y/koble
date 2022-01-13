package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"

	"github.com/spf13/cobra"
)

var user string
var workDir string

var shellCmd = &cobra.Command{
	Use:               "shell [options] MACHINE [COMMAND [ARG...]]",
	Short:             "get a shell on a machine",
	ValidArgsFunction: cli.AutocompRunningMachine,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cli.NK.Config.Terminal.Launch {
			return cli.NK.LaunchInTerm(args[0])
		}
		return cli.NK.Shell(args[0], user, workDir)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	shellCmd.Flags().StringVarP(&user, "user", "u", "", "User to execute shell as.")
	shellCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	cli.AddTermFlags(shellCmd)
	cli.AddWaitFlag(shellCmd)
	machineCmd.AddCommand(shellCmd)
	cli.Commands = append(cli.Commands, shellCmd)
}
