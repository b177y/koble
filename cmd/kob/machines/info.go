package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:                   "info [options] MACHINE",
	Short:                 "get info about a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompMachine,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.MachineInfo(args[0], mInfoJson)
	},
}

func init() {
	infoCmd.Flags().BoolVar(&mInfoJson, "json", false, "Print machine info as json object to stdout")
	machineCmd.AddCommand(infoCmd)
	cli.Commands = append(cli.Commands, infoCmd)
}
