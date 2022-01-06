package koble

import (
	"github.com/b177y/netkit/driver"
)

type DriverInitialiser func() driver.Driver

var AvailableDrivers = map[string]DriverInitialiser{}

func registerDriver(name string, d DriverInitialiser) {
	AvailableDrivers[name] = d
}
