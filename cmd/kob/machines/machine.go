package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var networks []string
var image string

// whether to wait for operation to finish
// used by flags for start and stop subcommands
var wait bool

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "manage machines",
}

func init() {
	cli.RootCmd.AddCommand(machineCmd)
}
