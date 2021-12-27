package driver_test

import (
	"fmt"

	"github.com/b177y/netkit/driver"
)

func DeclareAllDriverTests(dt interface{}) bool {
	d, ok := dt.(driver.Driver)
	if !ok {
		panic("not a valid Driver interface")
	}
	err := d.SetupDriver(nil)
	if err != nil {
		panic(fmt.Sprint("couldn't set up driver for tests: ", err))
	}
	var (
		_ = DeclareStartMachineTests(d)
		_ = DeclareHaltMachineTests(d)
		_ = DeclareRemoveMachineTests(d)
		_ = DeclareExistsMachineTests(d)
		_ = DeclareGetStateMachineTests(d)
	)
	return true
}
