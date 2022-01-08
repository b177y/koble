package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var machineNetworks []string
var machineImage string

var addMachineNetworks []string
var addMachineImage string

var mListAll bool
var mListJson bool

var mInfoJson bool

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "start and manage koble machines",
}

var mstartCmd = &cobra.Command{
	Use:                   "start [options] MACHINENAME",
	Short:                 "start a koble machine",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     autocompNonRunningMachine,
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.StartMachineWithStatus(args[0], machineImage, machineNetworks, true)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mstopCmd = &cobra.Command{
	Use:                   "stop [options] MACHINE",
	Short:                 "stop a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     autocompRunningMachine,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.HaltMachine(args[0], false)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mdestroyCmd = &cobra.Command{
	Use:                   "destroy [options] MACHINE",
	Short:                 "force stop and remove a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     autocompMachine,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.DestroyMachine(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mremoveCmd = &cobra.Command{
	Use:                   "remove [options] MACHINE",
	Short:                 "remove a koble machine",
	Aliases:               []string{"rm"},
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     autocompMachine,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.RemoveMachine(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

var minfoCmd = &cobra.Command{
	Use:                   "info [options] MACHINE",
	Short:                 "get info about a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     autocompMachine,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.MachineInfo(args[0], mInfoJson)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var maddCmd = &cobra.Command{
	Use:                   "add [options] MACHINENAME",
	Short:                 "add a new machine to a lab",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := koble.AddMachineToLab(args[0], machineNetworks, machineImage)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var mlistCmd = &cobra.Command{
	Use:     "list",
	Short:   "List koble machines",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if !mListAll {
			if nk.Namespace == "" {
				fmt.Fprintln(os.Stderr, "Listing all machines in the GLOBAL namespace.")
				fmt.Fprintf(os.Stderr, "To see all machines use `koble machine list --all`\n\n")
			} else {
				fmt.Fprintf(os.Stderr, "Listing all machines within the namespace (%s).\n", nk.Namespace)
				fmt.Fprintf(os.Stderr, "To see all machines use `koble machine list --all`\n\n")
			}
		}
		err := nk.ListMachines(mListAll, mListJson)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// add subcommands to koble machine ...
	machineCmd.AddCommand(mstartCmd)
	machineCmd.AddCommand(mstopCmd)
	machineCmd.AddCommand(mdestroyCmd)
	machineCmd.AddCommand(mremoveCmd)
	machineCmd.AddCommand(minfoCmd)
	machineCmd.AddCommand(maddCmd)
	machineCmd.AddCommand(mlistCmd)
	// add subcommands to koble ...
	KobleCLI.AddCommand(mstartCmd)
	KobleCLI.AddCommand(mstopCmd)
	KobleCLI.AddCommand(mdestroyCmd)
	KobleCLI.AddCommand(mremoveCmd)
	KobleCLI.AddCommand(minfoCmd)
	KobleCLI.AddCommand(maddCmd)
	KobleCLI.AddCommand(mlistCmd)

	mstartCmd.Flags().StringVar(&machineImage, "image", "", "Image to run machine with.")
	mstartCmd.Flags().StringArrayVar(&machineNetworks, "networks", []string{}, "Networks to attach to machine")

	maddCmd.Flags().StringVar(&addMachineImage, "image", "", "Image to use for new machine.")
	maddCmd.Flags().StringArrayVar(&addMachineNetworks, "networks", []string{}, "Networks to add to new machine.")

	mlistCmd.Flags().BoolVarP(&mListAll, "all", "a", false, "List all machines (from all labs / non-labs)")
	mlistCmd.Flags().BoolVar(&mListJson, "json", false, "Print machine list as json array to stdout")

	minfoCmd.Flags().BoolVar(&mInfoJson, "json", false, "Print machine info as json object to stdout")
}
