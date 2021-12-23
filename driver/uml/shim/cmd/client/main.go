package main

import (
	"log"
	"path/filepath"

	"github.com/b177y/netkit/driver/uml/shim"
	"github.com/spf13/cobra"
)

var ShimClient = &cobra.Command{
	Use:                   "shim-client DIR",
	Short:                 "shim-client is a tool for connecting to a uml shim socket",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := shim.Attach(filepath.Join(args[0], "attach.sock"))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	ShimClient.Execute()
}
