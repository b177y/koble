package driver_test

import (
	"fmt"

	"github.com/b177y/netkit/driver"
	// log "github.com/sirupsen/logrus"
)

var testconf = map[string]interface{}{"testing": true}

func DeclareAllDriverTests(dt interface{}) error {
	// log.SetLevel(log.DebugLevel)
	d, ok := dt.(driver.Driver)
	if !ok {
		return fmt.Errorf("not a valid Driver interface")
	}

	err := d.SetupDriver(testconf)
	if err != nil {
		return fmt.Errorf("couldn't set up driver for tests: %w", err)
	}
	var (
		_ = DeclareStartMachineTests(d)
		_ = DeclareHaltMachineTests(d)
		_ = DeclareRemoveMachineTests(d)
		_ = DeclareExistsMachineTests(d)
		_ = DeclareGetStateMachineTests(d)
	)
	return nil
}
