package lab

import "github.com/spf13/cobra"

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "control koble labs",
}

func init() {
	kob.rootCmd.AddCommand(labCmd)
}
