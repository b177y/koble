package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/b177y/netkit/pkg/netkit"
	"github.com/spf13/cobra"
)

var labDescription string
var labAuthors []string
var labEmails []string
var labWeb []string

var labHaltForce bool
var labMachines []string
var labAllMachines bool

var lstartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.LabStart()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var lcleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.LabClean(labMachines, labAllMachines)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var lhaltCmd = &cobra.Command{
	Use:   "halt",
	Short: "Halt a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.LabHalt(labMachines, labHaltForce, labAllMachines)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var linfoCmd = &cobra.Command{
	Use:   "info",
	Short: "View lab info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lab info")
	},
}

var linitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := netkit.InitLab(labName, labDescription, labAuthors, labEmails, labWeb)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var lvalidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.Validate()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "The 'lab' subcommand is used to control netkit labs",
}

func init() {
	labCmd.AddCommand(lstartCmd)
	labCmd.AddCommand(lcleanCmd)
	labCmd.AddCommand(lhaltCmd)
	labCmd.AddCommand(linfoCmd)
	labCmd.AddCommand(linitCmd)
	labCmd.AddCommand(lvalidateCmd)

	linitCmd.Flags().StringVar(&labName, "name", "", "Name to give the lab. This will create a new directory with the specified name. If no name is given, the lab will be initialised in the current directory.")
	linitCmd.Flags().StringVar(&labDescription, "description", "", "Description of the new lab.")
	linitCmd.Flags().StringArrayVar(&labAuthors, "authors", []string{}, "Comma separated list of lab author(s)")
	linitCmd.Flags().StringArrayVar(&labEmails, "emails", []string{}, "Comma separated list of lab author emails.")
	linitCmd.Flags().StringArrayVar(&labWeb, "web", []string{}, "Comma separated list of lab web resource URLs.")

	lcleanCmd.Flags().StringArrayVar(&labMachines, "machines", []string{}, "Comma separated list of lab machines to clean.")
	lcleanCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Clean all Netkit machines, including those not in the current lab.")
	lhaltCmd.Flags().StringArrayVar(&labMachines, "machines", []string{}, "Comma separated list of lab machines to halt.")
	lhaltCmd.Flags().BoolVarP(&labHaltForce, "force", "f", false, "Force halt machines.")
	lhaltCmd.Flags().BoolVarP(&labAllMachines, "all", "a", false, "Halt all Netkit machines, including those not in the current lab.")
}
