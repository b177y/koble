package cmd

import (
	"log"

	"github.com/b177y/netkit/driver/podman"
	"github.com/spf13/cobra"
)

var command string
var user string
var detachMode bool
var workDir string

var shellCmd = &cobra.Command{
	Use:   "shell [MACHINE]",
	Short: "The 'shell' subcommand is used to connect to a shell on a netkit machine",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO get driver from ?cmd?
		machine := args[0]
		d := new(podman.PodmanDriver)
		err := d.SetupDriver()
		if err != nil {
			log.Fatal(err)
		}
		err = d.MachineExecShell(machine, command, user, detachMode, workDir)
		if err != nil {
			log.Fatal(err)
		}
		// custom command
		// --user to execute as
	},
}

func init() {
	shellCmd.Flags().StringVarP(&command, "command", "c", "/bin/bash", "Command to execute in shell.")
	shellCmd.Flags().StringVarP(&command, "user", "u", "", "User to execute shell as.")
	shellCmd.Flags().BoolVarP(&detachMode, "detach", "d", false, "Run the exec session in detached mode (backgrounded)")
	shellCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	// TODO --terminal mode
}
