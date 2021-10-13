package cmd

import (
	"log"

	"github.com/b177y/netkit/driver/podman"
	"github.com/spf13/cobra"
)

var command string
var user string

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "The 'shell' subcommand is used to connect to a shell on a netkit machine",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO get driver from ?cmd?
		d := new(podman.PodmanDriver)
		err := d.SetupDriver()
		if err != nil {
			log.Fatal(err)
		}
		// TODO get machine name from ?args?
		err = d.MachineExecShell(machine)
		if err != nil {
			log.Fatal(err)
		}
		// custom command
		// --user to execute as
	},
}

func init() {
	shellCmd.Flags().StringVarP(&machine, "machine", "m", "", "Machine to connect shell to.")
	shellCmd.Flags().StringVarP(&command, "command", "c", "", "Command to execute in shell.")
	shellCmd.Flags().StringVarP(&command, "user", "u", "", "User to execute shell as.")
}
