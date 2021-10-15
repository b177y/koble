package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/b177y/netkit/pkg/netkit"
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
		err := netkit.ExecMachineShell(machine, command, user, detachMode, workDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	shellCmd.Flags().StringVarP(&command, "command", "c", "/bin/bash", "Command to execute in shell.")
	shellCmd.Flags().StringVarP(&user, "user", "u", "", "User to execute shell as.")
	shellCmd.Flags().BoolVarP(&detachMode, "detach", "d", false, "Run the exec session in detached mode (backgrounded)")
	shellCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "Working directory to execute from.")
	// TODO --terminal mode
}
