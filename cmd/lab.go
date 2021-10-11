package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lstartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting labbb...")
	},
}

var lcleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cleaning lab")
	},
}

var lcrashCmd = &cobra.Command{
	Use:   "crash",
	Short: "Crash a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crashing lab")
	},
}

var lhaltCmd = &cobra.Command{
	Use:   "halt",
	Short: "Halt a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Halting lab")
	},
}

var linfoCmd = &cobra.Command{
	Use:   "info",
	Short: "View lab info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lab info")
	},
}

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "The 'lab' subcommand is used to control netkit labs",
}

func init() {
	labCmd.AddCommand(lstartCmd)
	labCmd.AddCommand(lcleanCmd)
	labCmd.AddCommand(lcrashCmd)
	labCmd.AddCommand(lhaltCmd)
	labCmd.AddCommand(linfoCmd)
}
