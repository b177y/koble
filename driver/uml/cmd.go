package uml

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/b177y/netkit/driver/uml/vecnet"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/spf13/cobra"
)

var cmd = cobra.Command{
	Use: "uml [NAMESPACE]",
}

var unshareCmd = cobra.Command{
	Use: "unshare",
	Run: func(cmd *cobra.Command, args []string) {
		var namespace string
		if len(args) == 0 {
			namespace = "GLOBAL"
		} else {
			namespace = args[0]
		}
		fmt.Println("Entering namespace for", namespace)
		err := vecnet.CreateAndEnterUserNS("koble")
		if err != nil {
			log.Fatal(err)
		}
		err = vecnet.WithNetNS(namespace, func(ns.NetNS) error {
			bashCmd := exec.Command("/bin/bash")
			bashCmd.Stdin = os.Stdin
			bashCmd.Stdout = os.Stdout
			bashCmd.Stderr = os.Stderr
			return bashCmd.Run()
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func (ud *UMLDriver) GetCLICommand() (command *cobra.Command, err error) {
	cmd.AddCommand(&unshareCmd)
	return &cmd, nil
}
