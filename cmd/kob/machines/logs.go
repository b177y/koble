package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var logsFollow bool
var logsTail int

var logsCmd = &cobra.Command{
	Use:   "logs MACHINE [options]",
	Short: "get logs from a machine",
	Args:  cobra.ExactArgs(1),
	Example: `koble machine logs a0 --follow
koble machine logs dh --tail 10`,
	ValidArgsFunction: cli.AutocompMachine,
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := args[0]
		return cli.NK.MachineLogs(machine, logsFollow, logsTail)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "Follow log output")
	logsCmd.Flags().IntVar(&logsTail, "tail", -1, "Output the specified number of LINES at the end of the logs.  Defaults to -1, which prints all lines")
	machineCmd.AddCommand(logsCmd)
	cli.Commands = append(cli.Commands, logsCmd)
}
