package machine

import (
	"fmt"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/output"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:                   "destroy [options] MACHINE",
	Short:                 "force stop and remove a machine",
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompMachine,
	DisableFlagsInUseLine: true,
	RunE:                  destroy,
}

func init() {
	machineCmd.AddCommand(destroyCmd)
	cli.Commands = append(cli.Commands, destroyCmd)
}

var destroy = func(cmd *cobra.Command, args []string) error {
	return output.WithSimpleContainer(
		fmt.Sprintf("Destroying machine %s", args[0]),
		nil,
		cli.NK.Config.NonInteractive,
		func(out output.Output) (err error) {
			defer func() {
				if err == nil {
					out.Success("Destroyed machine " + args[0])
				}
			}()
			return cli.NK.DestroyMachine(args[0])
		})
}
