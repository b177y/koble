package machine

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var infoJson bool

var infoCmd = &cobra.Command{
	Use:                   "info [options] MACHINE",
	Short:                 "get info about a koble machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompMachine,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.MachineInfo(args[0], infoJson)
	},
}

func init() {
	infoCmd.Flags().BoolVar(&infoJson, "json", false, "Print machine info as json object to stdout")
	machineCmd.AddCommand(infoCmd)
	cli.Commands = append(cli.Commands, infoCmd)
}
