package cmd

import (
	"errors"

	"github.com/b177y/koble/pkg/koble"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var nk *koble.Koble
var verbose bool
var quiet bool
var namespace string

var KobleCLI = &cobra.Command{
	Use:   "koble",
	Short: "Koble is a network emulation tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose && quiet {
			log.Fatal(errors.New("CLI Flags --verbose and --quiet cannot be used together."))
		}
		if verbose {
			log.SetLevel(log.DebugLevel)
		} else if quiet {
			log.SetLevel(log.ErrorLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}
		var err error
		nk, err = koble.NewKoble(namespace)
		if err != nil {
			log.Fatal(err)
		}
	},
	Version: koble.VERSION,
}

var useTerm bool
var useCon bool

// Shared flag variables
var machine string
var labName string

func init() {
	KobleCLI.AddCommand(labCmd)
	KobleCLI.AddCommand(machineCmd)
	KobleCLI.AddCommand(netCmd)
	KobleCLI.PersistentFlags().StringVar(&namespace, "namespace", "", "namespace to use")
	KobleCLI.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	KobleCLI.PersistentFlags().BoolVar(&quiet, "quiet", false, "only show warnings and errors")
	KobleCLI.RegisterFlagCompletionFunc("namespace", autocompNamespace)
}
