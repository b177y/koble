package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/b177y/netkit/pkg/netkit"
	"github.com/spf13/cobra"
)

var machineName string
var machineNetworks []string
var machineImage string

var addMachineName string
var addMachineNetworks []string
var addMachineImage string

var mListAll bool

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "The 'machine' subcommand is used to start and manage netkit machines",
}

var mstartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.StartMachine(machineName, machineImage, machineNetworks)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mhaltCmd = &cobra.Command{
	Use:                   "halt [options] MACHINE",
	Short:                 "Halt a netkit machine",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.HaltMachine(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mcleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Destroy a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crashing machine...")
	},
}

var minfoCmd = &cobra.Command{
	Use:   "info [options] MACHINE",
	Short: "Get info about a netkit machine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.MachineInfo(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

var maddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new machine to a lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := netkit.AddMachineToLab(addMachineName, machineNetworks, machineImage)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mlistCmd = &cobra.Command{
	Use:   "list",
	Short: "List netkit machines",
	Run: func(cmd *cobra.Command, args []string) {
		if !mListAll {
			if nk.Namespace == "" {
				fmt.Fprintln(os.Stderr, "Listing all machines in the GLOBAL namespace.")
				fmt.Fprintf(os.Stderr, "To see all machines use `netkit machine list --all`\n\n")
			} else {
				fmt.Fprintf(os.Stderr, "Listing all machines within the namespace (%s).\n", nk.Namespace)
				fmt.Fprintf(os.Stderr, "To see all machines use `netkit machine list --all`\n\n")
			}
		}
		err := nk.ListMachines(mListAll)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	machineCmd.AddCommand(mstartCmd)
	machineCmd.AddCommand(mcleanCmd)
	machineCmd.AddCommand(mhaltCmd)
	machineCmd.AddCommand(minfoCmd)
	machineCmd.AddCommand(maddCmd)
	machineCmd.AddCommand(mlistCmd)

	mstartCmd.Flags().StringVar(&machineName, "name", "", "Name to give machine")
	mstartCmd.MarkFlagRequired("name")
	mstartCmd.Flags().StringVar(&machineImage, "image", "", "Image to run machine with.")
	mstartCmd.Flags().StringArrayVar(&machineNetworks, "networks", []string{}, "Networks to attach to machine")

	maddCmd.Flags().StringVar(&addMachineName, "name", "", "Name for new machine.")
	maddCmd.MarkFlagRequired("name")
	maddCmd.Flags().StringVar(&addMachineImage, "image", "", "Image to use for new machine.")
	maddCmd.Flags().StringArrayVar(&addMachineNetworks, "networks", []string{}, "Networks to add to new machine.")

	mlistCmd.Flags().BoolVarP(&mListAll, "all", "a", false, "List all machines (from all labs / non-labs)")
}
