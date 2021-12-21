package cmd

import (
	"errors"

	"github.com/b177y/netkit/pkg/netkit"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var nk *netkit.Netkit
var verbose bool
var quiet bool
var namespace string

var NetkitCLI = &cobra.Command{
	Use:   "netkit",
	Short: "Netkit is a network emulation tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose && quiet {
			log.Fatal(errors.New("CLI Flags --verbose and --quiet cannot be used together."))
		}
		if verbose {
			log.SetLevel(log.DebugLevel)
		} else if quiet {
			log.SetLevel(log.ErrorLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}
		var err error
		nk, err = netkit.NewNetkit(namespace)
		if err != nil {
			log.Fatal(err)
		}
	},
	Version: netkit.VERSION,
}

var useTerm bool
var useCon bool

// Shared flag variables
var machine string
var labName string

func init() {
	NetkitCLI.AddCommand(labCmd)
	NetkitCLI.AddCommand(shellCmd)
	NetkitCLI.AddCommand(execCmd)
	NetkitCLI.AddCommand(attachCmd)
	NetkitCLI.AddCommand(logsCmd)
	NetkitCLI.AddCommand(machineCmd)
	NetkitCLI.AddCommand(netCmd)
	NetkitCLI.PersistentFlags().StringVar(&namespace, "namespace", "", "namespace to use")
	NetkitCLI.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	NetkitCLI.PersistentFlags().BoolVar(&quiet, "quiet", false, "only show warnings and errors")
	NetkitCLI.RegisterFlagCompletionFunc("namespace", autocompNamespace)
}
