package machine

import (
	"github.com/b177y/koble/cmd/kob"
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
	Short: "start and manage koble machines",
}

func init() {
	kob.RootCmd.AddCommand(machineCmd)
}
