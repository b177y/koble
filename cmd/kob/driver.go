package main

import (
	"github.com/b177y/koble/cmd/kob/cli"
	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/driver/podman"
	"github.com/b177y/koble/pkg/driver/uml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var driverCmd = &cobra.Command{
	Use:   "driver",
	Short: "manage a koble driver",
}

func init() {
	driver.RegisterDriver("podman", func() driver.Driver {
		return new(podman.PodmanDriver)
	})
	driver.RegisterDriver("uml", func() driver.Driver {
		return new(uml.UMLDriver)
	})
	err := driver.RegisterDriverCmds(driverCmd)
	if err != nil {
		log.Warnf("could not register driver subcommands: %v\n", err)
	}
	cli.RootCmd.AddCommand(driverCmd)
}
