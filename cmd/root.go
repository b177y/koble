package cmd

import (
	"github.com/b177y/netkit/pkg/netkit"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var nk *netkit.Netkit

var NetkitCLI = &cobra.Command{
	Use:   "netkit",
	Short: "Netkit is a network emulation tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		nk, err = netkit.NewNetkit()
		if err != nil {
			log.Fatal(err)
		}
	},
	Version: netkit.VERSION,
}

var verbose bool
var useTerm bool
var noTerm bool

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
