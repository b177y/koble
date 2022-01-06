package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "modify user settings in koble.yml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doing config stuff")
	},
}
