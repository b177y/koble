package cmd

import (
	"fmt"
	"log"

	"github.com/b177y/netkit/driver/podman"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "The 'logs' subcommand is used to get logs from netkit machines",
	Run: func(cmd *cobra.Command, args []string) {
		d := new(podman.PodmanDriver)
		err := d.SetupDriver()
		if err != nil {
			log.Fatal(err)
		}
		stdoutChan := make(chan string)
		stderrChan := make(chan string)
		go func() {
			for recv := range stdoutChan {
				fmt.Println(recv)
			}
		}()
		go func() {
			for recv := range stderrChan {
				fmt.Println(recv)
			}
		}()
		err = d.GetMachineLogs(machine, stdoutChan, stderrChan)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// TODO change this to positional arg
	logsCmd.Flags().StringVarP(&machine, "machine", "m", "", "Machine to get logs from.")
}
