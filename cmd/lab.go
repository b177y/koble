package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lstartCmd = &cobra.Command{
	Use:   "start",
	Short: "start a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting labbb...")
	},
}

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "The 'lab' subcommand is used to control netkit labs",
}

func init() {
	labCmd.AddCommand(lstartCmd)
}
