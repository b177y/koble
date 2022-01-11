package lab

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var initOpts koble.InitOpts

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialise a new lab",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.InitLab(initOpts)
	},
}

func init() {
	initCmd.Flags().StringVar(&initOpts.Name, "name", "", "Name to give the lab. This will create a new directory with the specified name. If no name is given, the lab will be initialised in the current directory.")
	initCmd.Flags().StringVar(&initOpts.Description, "description", "", "Description of the new lab")
	initCmd.Flags().StringArrayVar(&initOpts.Authors, "author", []string{}, "lab author")
	initCmd.Flags().StringArrayVar(&initOpts.Emails, "email", []string{}, "email associated with lab")
	initCmd.Flags().StringArrayVar(&initOpts.Webs, "web", []string{}, "lab web resource URLs associated with lab")
	labCmd.AddCommand(initCmd)
}
