package cmd

import (
	"github.com/b177y/netkit/pkg/netkit"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logsFollow bool
var logsTail int

var logsCmd = &cobra.Command{
	Use:   "logs [MACHINE]",
	Short: "The 'logs' subcommand is used to get logs from a netkit",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		machine := args[0]
		err := netkit.MachineLogs(machine, logsFollow, logsTail)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// TODO change this to positional arg
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "Follow log output")
	logsCmd.Flags().IntVar(&logsTail, "tail", -1, "Output the specified number of LINES at the end of the logs.  Defaults to -1, which prints all lines")
}
