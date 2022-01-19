package machine

import (
	"fmt"
	"os"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var mListAll bool
var mListJson bool

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list machines",
	Aliases: []string{"ls"},
	Example: `koble machine list --json
koble machine ls --all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !mListAll && !cli.NK.Config.Quiet {
			fmt.Fprintf(os.Stderr, "Listing all machines within the namespace (%s).\n", cli.NK.Config.Namespace)
			fmt.Fprintf(os.Stderr, "To see all machines use `koble machine list --all`\n\n")
		}
		return cli.NK.ListMachines(mListAll, mListJson)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&mListAll, "all", "a", false, "List from all namespaces. This overrides the --namespace option.")
	listCmd.Flags().BoolVar(&mListJson, "json", false, "Print machine list as json array to stdout")
	machineCmd.AddCommand(listCmd)
	cli.RootCmd.AddCommand(listCmd)
}
