package netkit

import (
	"github.com/b177y/netkit/driver"
)

type DriverInitialiser func() driver.Driver

var availableDrivers = map[string]DriverInitialiser{}

func registerDriver(name string, d DriverInitialiser) {
	availableDrivers[name] = d
}
