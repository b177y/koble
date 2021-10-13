package cmd

import (
	"github.com/spf13/cobra"
)

var NetkitCLI = &cobra.Command{
	Use:   "netkit",
	Short: "Netkit is a network emulation tool",
}

var verbose bool

// Shared flag variables
var machine string
var lab string

func init() {
	NetkitCLI.AddCommand(labCmd)
	NetkitCLI.AddCommand(shellCmd)
	NetkitCLI.AddCommand(attachCmd)
	NetkitCLI.AddCommand(logsCmd)
	NetkitCLI.AddCommand(machineCmd)
	NetkitCLI.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
