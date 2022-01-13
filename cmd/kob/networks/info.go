package networks

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var infoJson bool

var infoCmd = &cobra.Command{
	Use:                   "info [options] NETWORK",
	Short:                 "get info about a network",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.NetworkInfo(args[0], infoJson)
	},
}

func init() {
	infoCmd.Flags().BoolVar(&infoJson, "json", false, "print network info as json object to stdout")
	netCmd.AddCommand(infoCmd)
}
