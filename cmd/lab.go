package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var labDescription string
var labAuthors []string
var labEmails []string
var labWeb []string

var labHaltForce bool
var labAllMachines bool

var lstartCmd = &cobra.Command{
	Use:   "start [options] MACHINE [MACHINE...]",
	Short: "start a koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.LabStart(args)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

var ldestroyCmd = &cobra.Command{
	Use:   "destroy [options] MACHINE [MACHINE...]",
	Short: "crash and remove all machines in a koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.LabDestroy(args, labAllMachines)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

var lhaltCmd = &cobra.Command{
	Use:   "stop [options] MACHINE [MACHINE...]",
	Short: "stop machines in a koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.LabHalt(args, labHaltForce, labAllMachines)
		if err != nil {
			log.Fatal(err)
		}
	},
	DisableFlagsInUseLine: true,
}

var linfoCmd = &cobra.Command{
	Use:   "info",
	Short: "view lab info",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.LabInfo()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var linitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialise a new koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := koble.InitLab(labName, labDescription, labAuthors, labEmails, labWeb)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var lvalidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate a koble lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.Validate()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "control koble labs",
}

func init() {
	labCmd.AddCommand(lstartCmd)
	labCmd.AddCommand(ldestroyCmd)
	labCmd.AddCommand(lhaltCmd)
	labCmd.AddCommand(linfoCmd)
	labCmd.AddCommand(linitCmd)
	labCmd.AddCommand(lvalidateCmd)

	linitCmd.Flags().StringVar(&labName, "name", "", "Name to give the lab. This will create a new directory with the specified name. If no name is given, the lab will be initialised in the current directory.")
	linitCmd.Flags().StringVar(&labDescription, "description", "", "Description of the new lab.")
	linitCmd.Flags().StringArrayVar(&labAuthors, "authors", []string{}, "Comma separated list of lab author(s)")
	linitCmd.Flags().StringArrayVar(&labEmails, "emails", []string{}, "Comma separated list of lab author emails.")
	linitCmd.Flags().StringArrayVar(&labWeb, "web", []string{}, "Comma separated list of lab web resource URLs.")

	ldestroyCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Destroy all koble machines, including those not in the current lab.")
	lhaltCmd.Flags().BoolVarP(&labHaltForce, "force", "f", false, "Force halt machines.")
	lhaltCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Halt all koble machines, including those not in the current lab.")
}
