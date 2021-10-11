package cmd

import (
	"github.com/spf13/cobra"
)

var NetkitCLI = &cobra.Command{
	Use:   "netkit",
	Short: "Netkit is a network emulation tool",
}

func init() {
	NetkitCLI.AddCommand(labCmd)
	NetkitCLI.AddCommand(connectCmd)
}
