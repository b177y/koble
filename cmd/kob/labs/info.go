package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/spf13/cobra"
)

var linfoCmd = &cobra.Command{
	Use:   "info",
	Short: "view lab info",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.Lab.Info()
	},
}
