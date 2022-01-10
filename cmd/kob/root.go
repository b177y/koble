package main

import (
	"errors"

	"github.com/b177y/koble/cmd/kob/cli"
	_ "github.com/b177y/koble/cmd/kob/labs"
	_ "github.com/b177y/koble/cmd/kob/machines"
	"github.com/b177y/koble/pkg/koble"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verbose bool
var quiet bool
var namespace string

var (
	rootCmd = &cobra.Command{
		Use:   "koble",
		Short: "Koble is a network emulation tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
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
			cli.NK, err = koble.Load(namespace)
			if err != nil {
				return err
			}
			err = cli.NK.LoadLab()
			if err != nil {
				return err
			}
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return cli.NK.ExitLab()
		},
		Version:       koble.VERSION,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

var useTerm bool
var useCon bool

// Shared flag variables
var machine string
var labName string

func init() {
	rootCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "namespace to use")
	rootCmd.RegisterFlagCompletionFunc("namespace", cli.AutocompNamespace)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "only show warnings and errors")
	rootCmd.PersistentFlags().Bool("plain", false, "disable interactive and coloured output")
	viper.BindPFlag("noninteractive", rootCmd.PersistentFlags().Lookup("plain"))
	rootCmd.PersistentFlags().Bool("no-color", false, "disable coloured output")
	viper.BindPFlag("nocolor", rootCmd.PersistentFlags().Lookup("no-color"))
	for _, c := range cli.Commands {
		rootCmd.AddCommand(c)
	}
}
