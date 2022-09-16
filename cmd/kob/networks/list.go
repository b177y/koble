package networks

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var listAll bool
var listJson bool

var listCmd = &cobra.Command{
	Use:                   "list [options]",
	Short:                 "list all koble networks",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.ListNetworks(listAll, listJson)
	},
}

func init() {
	listCmd.Flags().BoolVar(&listJson, "json", false, "print network info as json object to stdout")
	listCmd.Flags().BoolVarP(&listAll, "all", "a", false, "list networks from all namespaces")
	netCmd.AddCommand(listCmd)
}
