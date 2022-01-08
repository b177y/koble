package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

// whether to wait for lab operation to finish
// used by flags for start and stop subcommands
var wait bool

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "manage labs",
}

func init() {
	cli.Commands = append(cli.Commands, labCmd)
}
