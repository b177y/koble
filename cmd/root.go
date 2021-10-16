package cmd

import (
	"fmt"

	"github.com/b177y/netkit/pkg/netkit"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var NetkitCLI = &cobra.Command{
	Use:     "netkit",
	Short:   "Netkit is a network emulation tool",
	Version: netkit.VERSION,
}

var verbose bool
var useTerm bool
var noTerm bool

var config netkit.Config

// Shared flag variables
var machine string
var labName string

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/netkit")
	viper.AddConfigPath("./examples/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
	NetkitCLI.AddCommand(labCmd)
	NetkitCLI.AddCommand(shellCmd)
	NetkitCLI.AddCommand(attachCmd)
	NetkitCLI.AddCommand(logsCmd)
	NetkitCLI.AddCommand(machineCmd)
	NetkitCLI.AddCommand(netCmd)
	NetkitCLI.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
