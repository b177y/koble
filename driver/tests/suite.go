package driver_test

import "github.com/b177y/netkit/driver"

func DeclareAllDriverTests(d driver.Driver) bool {
	var (
		_ = DeclareStartMachineTests(d)
	)
	return true
}
