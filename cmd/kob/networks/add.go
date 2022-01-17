package networks

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/driver"
	"github.com/spf13/cobra"
)

var netConf driver.NetConfig

var addCmd = &cobra.Command{
	Use:                   "add [options] NAME",
	Short:                 "add a new network to a lab",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NK.AddNetworkToLab(args[0], netConf)
	},
}

func init() {
	addCmd.Flags().BoolVar(&netConf.External, "external", false, "Allow access to external networks")
	addCmd.Flags().StringVar(&netConf.Gateway, "gateway", "", "IPv4 or IPv6 gateway for the subnet")
	addCmd.Flags().StringVar(&netConf.Subnet, "subnet", "", "subnet in CIDR format")
	addCmd.Flags().StringVar(&netConf.IPv6, "ipv6", "", "ipv6 networking subnet")
	netCmd.AddCommand(addCmd)
}
