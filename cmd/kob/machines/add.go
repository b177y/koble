package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var addNetworks []string
var addImage string

var addCmd = &cobra.Command{
	Use:                   "add [options] MACHINENAME",
	Short:                 "add a new machine to a lab",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return koble.AddMachineToLab(args[0], addNetworks, addImage)
	},
}

func init() {
	addCmd.Flags().StringVar(&addImage, "image", "", "Image to use for new machine.")
	addCmd.Flags().StringArrayVar(&addNetworks, "networks", []string{}, "Networks to add to new machine.")
	machineCmd.AddCommand(addCmd)
	cli.Commands = append(cli.Commands, addCmd)
}
