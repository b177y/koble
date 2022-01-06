package cmd

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/b177y/netkit/pkg/koble"
	"github.com/spf13/cobra"
)

var networkName string
var networkExternal bool
var networkGateway net.IP
var networkSubnet net.IPNet
var networkIpv6 bool

var nListAll bool

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "view and manage koble networks",
}

var ninfoCmd = &cobra.Command{
	Use:   "info",
	Short: "get info about a koble network",
	Run: func(cmd *cobra.Command, args []string) {
		err := nk.NetworkInfo(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

var nlistCmd = &cobra.Command{
	Use:   "list",
	Short: "list koble networks",
	Run: func(cmd *cobra.Command, args []string) {
		if !nListAll {
			if nk.Lab.Name == "" {
				fmt.Fprintln(os.Stderr, "Listing all networks which are not associated with a lab.")
				fmt.Fprintf(os.Stderr, "To see all machines use `koble net list --all`\n\n")
			} else {
				fmt.Fprintf(os.Stderr, "Listing all networks within this lab (%s).\n", nk.Lab.Name)
				fmt.Fprintf(os.Stderr, "To see all machines use `koble net list --all`\n\n")
			}
		}
		err := nk.ListNetworks(nListAll)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var naddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new network to a lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := koble.AddNetworkToLab(networkName, networkExternal, networkGateway, networkSubnet, networkIpv6)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	netCmd.AddCommand(naddCmd)
	netCmd.AddCommand(nlistCmd)
	netCmd.AddCommand(ninfoCmd)

	naddCmd.Flags().StringVar(&networkName, "name", "", "Name for new network")
	naddCmd.MarkFlagRequired("name")
	naddCmd.Flags().BoolVar(&networkExternal, "external", false, "Allow access to external networks")
	naddCmd.Flags().IPVar(&networkGateway, "gateway", net.IP(""), "IPv4 or IPv6 gateway for the subnet")
	var ipNet net.IPNet
	naddCmd.Flags().IPNetVar(&networkSubnet, "subnet", ipNet, "subnet in CIDR format")
	naddCmd.Flags().BoolVar(&networkIpv6, "ipv6", false, "enable ipv6 networking")
	nlistCmd.Flags().BoolVarP(&nListAll, "all", "a", false, "List all networks (from all labs / non-labs)")
}
