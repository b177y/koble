package machine

import (
	"fmt"

	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/output"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:                   "remove [options] MACHINE",
	Short:                 "remove a machine",
	Aliases:               []string{"rm"},
	Args:                  cobra.ExactArgs(1),
	ValidArgsFunction:     cli.AutocompNonRunningMachine,
	DisableFlagsInUseLine: true,
	Example:               `koble machine remove a0`,
	RunE:                  remove,
}

func init() {
	machineCmd.AddCommand(removeCmd)
	cli.RootCmd.AddCommand(removeCmd)
}

var remove = func(cmd *cobra.Command, args []string) error {
	return output.WithSimpleContainer(
		fmt.Sprintf("Removing machine %s", args[0]),
		nil,
		cli.NK.Config.NonInteractive,
		func(out output.Output) (err error) {
			defer func() {
				if err == nil {
					out.Success("Removed machine " + args[0])
				}
			}()
			return cli.NK.RemoveMachine(args[0], out)
		})
}
