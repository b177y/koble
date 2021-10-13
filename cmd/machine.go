package cmd

import (
	"fmt"
	"log"

	"github.com/b177y/netkit/pkg/netkit"
	"github.com/spf13/cobra"
)

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "The 'machine' subcommand is used to start and manage netkit machines",
}

var mstartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting machine...")
		err := netkit.StartMachine(machineName, machineImage, machineNetworks)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mcrashCmd = &cobra.Command{
	Use:   "crash",
	Short: "Crash a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crashing machine...")
	},
}

var mhaltCmd = &cobra.Command{
	Use:   "halt",
	Short: "Halt a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Halting machine...")
	},
}

var minfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getting machine info...")
	},
}

var maddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new machine to a lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := netkit.AddMachineToLab(machineName, machineNetworks, machineImage)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	machineCmd.AddCommand(mstartCmd)
	machineCmd.AddCommand(mcrashCmd)
	machineCmd.AddCommand(mhaltCmd)
	machineCmd.AddCommand(minfoCmd)
	machineCmd.AddCommand(maddCmd)

	mstartCmd.Flags().StringVar(&machineName, "name", "", "Name to give machine")
	mstartCmd.MarkFlagRequired("name")
	mstartCmd.Flags().StringVar(&machineImage, "image", "localhost/netkit-deb-test", "Image to run machine with.")
	mstartCmd.Flags().StringArrayVar(&machineNetworks, "networks", []string{}, "Networks to attach to machine")

	maddCmd.Flags().StringVar(&machineName, "name", "", "Name for new machine.")
	maddCmd.MarkFlagRequired("name")
	maddCmd.Flags().StringVar(&machineImage, "image", "", "Image to use for new machine.")
	maddCmd.Flags().StringArrayVar(&machineNetworks, "networks", []string{}, "Networks to add to new machine.")
}
