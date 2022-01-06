package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logsFollow bool
var logsTail int

var logsCmd = &cobra.Command{
	Use:   "logs [options] MACHINE",
	Short: "get logs from a koble machine",
	Args:  cobra.ExactArgs(1),
	Example: `koble logs a0 -f
	koble logs dh --tail 10`,
	ValidArgsFunction: autocompMachine,
	Run: func(cmd *cobra.Command, args []string) {
		machine := args[0]
		err := nk.MachineLogs(machine, logsFollow, logsTail)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

func init() {
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "Follow log output")
	logsCmd.Flags().IntVar(&logsTail, "tail", -1, "Output the specified number of LINES at the end of the logs.  Defaults to -1, which prints all lines")
	KobleCLI.AddCommand(logsCmd)
	machineCmd.AddCommand(logsCmd)
}
