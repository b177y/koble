package cmd

import (
	"log"

	"github.com/b177y/netkit/driver/podman"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "The 'connect' subcommand is used to connect to netkit machines",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO get driver from ?cmd?
		d := new(podman.PodmanDriver)
		err := d.SetupDriver()
		if err != nil {
			log.Fatal(err)
		}
		// TODO get machine name from ?args?
		err = d.ConnectToMachine("h12")
		if err != nil {
			log.Fatal(err)
		}
	},
}
