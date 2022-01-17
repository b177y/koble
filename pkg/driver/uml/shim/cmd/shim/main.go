package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/b177y/koble/driver/uml/shim"
	"github.com/docker/docker/pkg/reexec"
	"github.com/spf13/cobra"
)

var directory string

var UMLShimCLI = &cobra.Command{
	Use:                   "uml-shim [options] KERNELCMD",
	Short:                 "uml-shim is a tool for running and managing a UserMode Linux instance",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := reexec.Command("umlShim")
		c.Args = append(c.Args, directory)
		c.Args = append(c.Args, args...)
		c.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}
		// c.Stdout = os.Stdout
		// c.Stderr = os.Stderr
		fmt.Println("Running umlShim", c)
		if err := c.Start(); err != nil {
			log.Fatalf("failed to run command: %s", err)
		}
	},
}

func main() {
	UMLShimCLI.Execute()
}

func init() {
	reexec.Register("umlShim", shim.RunShim)
	if reexec.Init() {
		os.Exit(0)
	}
	UMLShimCLI.Flags().StringVarP(&directory, "directory", "d", "", "directory")
}
