package cli

import (
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var namespace string

var (
	RootCmd = &cobra.Command{
		Use:   "koble",
		Short: "Koble is a network emulation tool",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			NK, err = koble.Load()
			if err != nil {
				return err
			}
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return NK.Cleanup()
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
	RootCmd.PersistentFlags().String("namespace", "", "namespace to use")
	koble.BindFlag("namespace", RootCmd.PersistentFlags().Lookup("namespace"))
	RootCmd.RegisterFlagCompletionFunc("namespace", AutocompNamespace)
	RootCmd.PersistentFlags().CountP("verbose", "v", "verbose output")
	koble.BindFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	RootCmd.PersistentFlags().Bool("quiet", false, "only show errors in log errors")
	koble.BindFlag("quiet", RootCmd.PersistentFlags().Lookup("quiet"))
	RootCmd.PersistentFlags().String("driver", "", "disable interactive and coloured output")
	koble.BindFlag("driver.name", RootCmd.PersistentFlags().Lookup("driver"))
	// TODO add autocomp for --driver (list available drivers)
	RootCmd.PersistentFlags().Bool("plain", false, "disable interactive and coloured output")
	koble.BindFlag("noninteractive", RootCmd.PersistentFlags().Lookup("plain"))
	RootCmd.PersistentFlags().Bool("no-color", false, "disable coloured output")
	koble.BindFlag("nocolor", RootCmd.PersistentFlags().Lookup("no-color"))
}
