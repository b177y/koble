package networks

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "manage networks",
}

// var nlistCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "list koble networks",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if !nListAll {
// 			if nk.Lab.Name == "" {
// 				fmt.Fprintln(os.Stderr, "Listing all networks which are not associated with a lab.")
// 				fmt.Fprintf(os.Stderr, "To see all machines use `koble net list --all`\n\n")
// 			} else {
// 				fmt.Fprintf(os.Stderr, "Listing all networks within this lab (%s).\n", nk.Lab.Name)
// 				fmt.Fprintf(os.Stderr, "To see all machines use `koble net list --all`\n\n")
// 			}
// 		}
// 		err := nk.ListNetworks(nListAll)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	},
// }

func init() {
	cli.RootCmd.AddCommand(netCmd)
	// netCmd.AddCommand(nlistCmd)
	// netCmd.AddCommand(ninfoCmd)

	// nlistCmd.Flags().BoolVarP(&nListAll, "all", "a", false, "List all networks (from all labs / non-labs)")
}
