package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "The 'config' subcommand is used to modify settings in netkit.yml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doing config stuff")
	},
}
