package kob

import (
	"errors"

	"github.com/b177y/koble/pkg/koble"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var NK *koble.Koble
var verbose bool
var quiet bool
var plain bool
var noColor bool
var namespace string

var (
	RootCmd = &cobra.Command{
		Use:   "koble",
		Short: "Koble is a network emulation tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
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
			if plain || noColor {
				color.NoColor = true
			}
			var err error
			NK, err = koble.NewKoble(namespace)
			return err
		},
		Version: koble.VERSION,
	}
)

var useTerm bool
var useCon bool

// Shared flag variables
var machine string
var labName string

func init() {
	RootCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "namespace to use")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "only show warnings and errors")
	RootCmd.PersistentFlags().BoolVar(&plain, "plain", false, "disable interactive / coloured output")
	RootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable interactive / coloured output")
	RootCmd.RegisterFlagCompletionFunc("namespace", autocompNamespace)
}
