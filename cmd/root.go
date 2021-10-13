package cmd

import (
	"github.com/spf13/cobra"
)

var NetkitCLI = &cobra.Command{
	Use:   "netkit",
	Short: "Netkit is a network emulation tool",
}
var verbose bool

func init() {
	NetkitCLI.AddCommand(labCmd)
	NetkitCLI.AddCommand(connectCmd)
	NetkitCLI.AddCommand(logsCmd)
	NetkitCLI.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
