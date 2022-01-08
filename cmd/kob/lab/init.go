package lab

import (
	"github.com/b177y/koble/pkg/koble"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var initOpts koble.InitOpts

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialise a new koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := koble.InitLab(initOpts)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	initCmd.Flags().StringVar(&initOpts.Name, "name", "", "Name to give the lab. This will create a new directory with the specified name. If no name is given, the lab will be initialised in the current directory.")
	initCmd.Flags().StringVar(&initOpts.Description, "description", "", "Description of the new lab.")
	initCmd.Flags().StringArrayVar(&initOpts.Authors, "authors", []string{}, "Comma separated list of lab author(s)")
	initCmd.Flags().StringArrayVar(&initOpts.Emails, "emails", []string{}, "Comma separated list of lab author emails.")
	initCmd.Flags().StringArrayVar(&initOpts.Webs, "web", []string{}, "Comma separated list of lab web resource URLs.")
	labCmd.AddCommand(initCmd)
}
