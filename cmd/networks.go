package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/b177y/netkit/pkg/netkit"
	"github.com/spf13/cobra"
)

var networkName string
var networkExternal bool
var networkGateway net.IP
var networkSubnet net.IPNet
var networkIpv6 bool

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "The 'net' subcommand is used to view and manage netkit networks",
}

var ninfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a netkit network",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getting machine info...")
	},
}

var nlistCmd = &cobra.Command{
	Use:   "list",
	Short: "list netkit networks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getting machine info...")
	},
}

var naddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new network to a lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := netkit.AddNetworkToLab(networkName, networkExternal, networkGateway, networkSubnet, networkIpv6)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	netCmd.AddCommand(naddCmd)

	naddCmd.Flags().StringVar(&networkName, "name", "", "Name for new network")
	naddCmd.MarkFlagRequired("name")
	naddCmd.Flags().BoolVar(&networkExternal, "external", false, "Allow access to external networks")
	naddCmd.Flags().IPVar(&networkGateway, "gateway", net.IP(""), "IPv4 or IPv6 gateway for the subnet")
	var ipNet net.IPNet
	naddCmd.Flags().IPNetVar(&networkSubnet, "subnet", ipNet, "subnet in CIDR format")
	naddCmd.Flags().BoolVar(&networkIpv6, "ipv6", false, "enable ipv6 networking")
}
