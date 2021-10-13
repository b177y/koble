package cmd

import (
	"log"

	"github.com/b177y/netkit/driver/podman"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "The 'attach' subcommand is used to attach to the main tty on a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO get driver from ?cmd?
		d := new(podman.PodmanDriver)
		err := d.SetupDriver()
		if err != nil {
			log.Fatal(err)
		}
		err = d.AttachToMachine(machine)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	attachCmd.Flags().StringVarP(&machine, "machine", "m", "", "Machine to attach to.")
	// TODO add ability to change detach sequence + add to config
}
