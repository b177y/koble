package lab

import (
	"github.com/b177y/koble/cmd/kob"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var linfoCmd = &cobra.Command{
	Use:   "info",
	Short: "view lab info",
	Run: func(cmd *cobra.Command, args []string) {
		err := kob.NK.LabInfo()
		if err != nil {
			log.Fatal(err)
		}
	},
}
