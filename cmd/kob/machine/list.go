package machine

import (
	"fmt"
	"os"

	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var mListAll bool
var mListJson bool

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List koble machines",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if !mListAll {
			if kob.NK.Namespace == "" {
				fmt.Fprintln(os.Stderr, "Listing all machines in the GLOBAL namespace.")
				fmt.Fprintf(os.Stderr, "To see all machines use `koble machine list --all`\n\n")
			} else {
				fmt.Fprintf(os.Stderr, "Listing all machines within the namespace (%s).\n", kob.NK.Namespace)
				fmt.Fprintf(os.Stderr, "To see all machines use `koble machine list --all`\n\n")
			}
		}
		err := kob.NK.ListMachines(mListAll, mListJson)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&mListAll, "all", "a", false, "List all machines (from all labs / non-labs)")
	listCmd.Flags().BoolVar(&mListJson, "json", false, "Print machine list as json array to stdout")
	machineCmd.AddCommand(listCmd)
	kob.RootCmd.AddCommand(listCmd)
}
