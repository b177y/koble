package cmd

import (
	"github.com/b177y/netkit/pkg/netkit"
	"github.com/spf13/cobra"
)

var NetkitCLI = &cobra.Command{
	Use:     "netkit",
	Short:   "Netkit is a network emulation tool",
	Version: netkit.VERSION,
}

var verbose bool

// Shared flag variables
var machine string
var labName string

func init() {
	NetkitCLI.AddCommand(labCmd)
	NetkitCLI.AddCommand(shellCmd)
	NetkitCLI.AddCommand(attachCmd)
	NetkitCLI.AddCommand(logsCmd)
	NetkitCLI.AddCommand(machineCmd)
	NetkitCLI.AddCommand(netCmd)
	NetkitCLI.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
