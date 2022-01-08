package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:                   "add [options] MACHINENAME",
	Short:                 "add a new machine to a lab",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return koble.AddMachineToLab(args[0], machineNetworks, machineImage)
	},
}

func init() {
	addCmd.Flags().StringVar(&addMachineImage, "image", "", "Image to use for new machine.")
	addCmd.Flags().StringArrayVar(&addMachineNetworks, "networks", []string{}, "Networks to add to new machine.")
	machineCmd.AddCommand(addCmd)
	cli.Commands = append(cli.Commands, addCmd)
}
