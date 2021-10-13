package cmd

import (
	"fmt"
	"log"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/podman"
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
		m := driver.Machine{
			Name:     "h12",
			Hostlab:  "/home/billy/repos/rootless-netkit/examples/lab04",
			Hosthome: "/home/billy",
			Networks: []string{},
			Image:    "localhost/netkit-deb-test",
		}
		d := new(podman.PodmanDriver)
		err := d.SetupDriver()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Starting machine")
		_, err = d.StartMachine(m)
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

func init() {
	machineCmd.AddCommand(mstartCmd)
	machineCmd.AddCommand(mcrashCmd)
	machineCmd.AddCommand(mhaltCmd)
	machineCmd.AddCommand(minfoCmd)
}
