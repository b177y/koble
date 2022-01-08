package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var machineNetworks []string
var machineImage string
var machineWait bool

var addMachineNetworks []string
var addMachineImage string

var mInfoJson bool

var machineCmd = &cobra.Command{
	Use:   "machine",
	Short: "manage machines",
}

func init() {
	cli.Commands = append(cli.Commands, machineCmd)
}
