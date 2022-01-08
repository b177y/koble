package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "control koble labs",
}

func init() {
	cli.Commands = append(cli.Commands, labCmd)
}
