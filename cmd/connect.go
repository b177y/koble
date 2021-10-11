package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "The 'connect' subcommand is used to connect to netkit machines",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connecting")
	},
}
