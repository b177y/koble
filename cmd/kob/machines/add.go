package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/driver"
	"github.com/spf13/cobra"
)

var machineConfig driver.MachineConfig

var addCmd = &cobra.Command{
	Use:                   "add [options] NAME",
	Short:                 "add a new machine to a lab",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Example: `koble machine add a0 --image custom-image --network a1
koble machine add web --network extNet --network lan`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.AddMachineToLab(args[0], machineConfig)
	},
}

func init() {
	addCmd.Flags().StringVar(&machineConfig.Image, "image", "", "Image to use for new machine")
	addCmd.Flags().StringArrayVar(&machineConfig.Networks, "network", []string{}, "networks to add to new machine")
	machineCmd.AddCommand(addCmd)
	cli.RootCmd.AddCommand(addCmd)
}
