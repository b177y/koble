package cmd

import (
	"fmt"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/pkg/koble"
	"github.com/spf13/cobra"
)

var driverCmd = &cobra.Command{
	Use:   "driver",
	Short: "manage a koble driver",
}

func init() {
	for _, d := range koble.AvailableDrivers {
		if dCmd, err := d().GetCLICommand(); err == driver.ErrNotImplemented {
			continue
		} else if err != nil {
			fmt.Println("Error: %w", err)
		} else {
			driverCmd.AddCommand(dCmd)
		}
	}
	KobleCLI.AddCommand(driverCmd)
}
