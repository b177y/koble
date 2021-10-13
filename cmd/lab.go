package cmd

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/b177y/netkit/pkg/netkit"
	"github.com/spf13/cobra"
)

var labDescription string
var labAuthors []string
var labEmails []string
var labWeb []string

var machineName string
var machineNetworks []string
var machineImage string

var networkName string
var networkInternal bool
var networkGateway net.IP
var networkSubnet net.IPNet
var networkIpv6 bool

var lstartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting labbb...")
	},
}

var lcleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cleaning lab")
	},
}

var lcrashCmd = &cobra.Command{
	Use:   "crash",
	Short: "Crash a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crashing lab")
	},
}

var lhaltCmd = &cobra.Command{
	Use:   "halt",
	Short: "Halt a netkit lab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Halting lab")
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

var laddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new machine or network to a lab",
}

var machAddCmd = &cobra.Command{
	Use:   "machine",
	Short: "Add a new machine to a lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := netkit.AddMachineToLab(machineName, machineNetworks, machineImage)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var netAddCmd = &cobra.Command{
	Use:   "net",
	Short: "Add a new network to a lab",
	Run: func(cmd *cobra.Command, args []string) {
		err := netkit.AddNetworkToLab(networkName, networkInternal, networkGateway, networkSubnet, networkIpv6)
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
	labCmd.AddCommand(lcrashCmd)
	labCmd.AddCommand(lhaltCmd)
	labCmd.AddCommand(linfoCmd)
	labCmd.AddCommand(linitCmd)
	labCmd.AddCommand(laddCmd)

	linitCmd.Flags().StringVar(&labName, "name", "", "Name to give the lab. This will create a new directory with the specified name. If no name is given, the lab will be initialised in the current directory.")
	linitCmd.Flags().StringVar(&labDescription, "description", "", "Description of the new lab.")
	linitCmd.Flags().StringArrayVar(&labAuthors, "authors", []string{}, "Comma separated list of lab author(s)")
	linitCmd.Flags().StringArrayVar(&labEmails, "emails", []string{}, "Comma separated list of lab author emails.")
	linitCmd.Flags().StringArrayVar(&labWeb, "web", []string{}, "Comma separated list of lab web resource URLs.")

	laddCmd.AddCommand(machAddCmd)
	laddCmd.AddCommand(netAddCmd)

	machAddCmd.Flags().StringVar(&machineName, "name", "", "Name for new machine.")
	machAddCmd.MarkFlagRequired("name")
	machAddCmd.Flags().StringVar(&machineImage, "image", "", "Image to use for new machine.")
	machAddCmd.Flags().StringArrayVar(&machineNetworks, "networks", []string{}, "Networks to add to new machine.")

	netAddCmd.Flags().StringVar(&networkName, "name", "", "Name for new network")
	netAddCmd.MarkFlagRequired("name")
	netAddCmd.Flags().BoolVar(&networkInternal, "internal", true, "restrict external access from this network")
	netAddCmd.Flags().IPVar(&networkGateway, "gateway", net.IP(""), "IPv4 or IPv6 gateway for the subnet")
	var ipNet net.IPNet
	netAddCmd.Flags().IPNetVar(&networkSubnet, "subnet", ipNet, "subnet in CIDR format")
	netAddCmd.Flags().BoolVar(&networkIpv6, "ipv6", false, "enable ipv6 networking")
}
