package main

import (
	"github.com/b177y/koble/cmd/kob/cli"
	_ "github.com/b177y/koble/cmd/kob/labs"
	_ "github.com/b177y/koble/cmd/kob/machines"
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var namespace string

var (
	rootCmd = &cobra.Command{
		Use:   "koble",
		Short: "Koble is a network emulation tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			cli.NK, err = koble.Load()
			if err != nil {
				return err
			}
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return cli.NK.Cleanup()
		},
		Version:       koble.VERSION,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

// Shared flag variables
var machine string
var labName string

func init() {
	rootCmd.PersistentFlags().String("namespace", "", "namespace to use")
	koble.BindFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace"))
	rootCmd.RegisterFlagCompletionFunc("namespace", cli.AutocompNamespace)
	rootCmd.PersistentFlags().CountP("verbose", "v", "verbose output")
	koble.BindFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	rootCmd.PersistentFlags().Bool("quiet", false, "only show errors in log errors")
	koble.BindFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	rootCmd.PersistentFlags().String("driver", "", "disable interactive and coloured output")
	koble.BindFlag("driver.name", rootCmd.PersistentFlags().Lookup("driver"))
	// TODO add autocomp for --driver (list available drivers)
	rootCmd.PersistentFlags().Bool("plain", false, "disable interactive and coloured output")
	koble.BindFlag("noninteractive", rootCmd.PersistentFlags().Lookup("plain"))
	rootCmd.PersistentFlags().Bool("no-color", false, "disable coloured output")
	koble.BindFlag("nocolor", rootCmd.PersistentFlags().Lookup("no-color"))
	for _, c := range cli.Commands {
		rootCmd.AddCommand(c)
	}
}
