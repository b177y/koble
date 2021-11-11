package uml_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/b177y/netkit/driver"
	driver_test "github.com/b177y/netkit/driver/test"
	"github.com/b177y/netkit/driver/uml"
)

var ud = new(uml.UMLDriver)

func TestUMLImplementsDriver(t *testing.T) {
	var d interface{} = ud
	v, ok := d.(driver.Driver)
	if !ok {
		fmt.Println(v, ok)
		t.Error("UMLDriver does not satisify driver interface.")
	}

}

func TestHaltMachine(t *testing.T)     { driver_test.HaltMachine(t, ud) }
func TestForceHalt(t *testing.T)       { driver_test.ForceHalt(t, ud) }
func TestGetMachineState(t *testing.T) { driver_test.GetMachineState(t, ud) }
func TestAttachToMachine(t *testing.T) { driver_test.AttachToMachine(t, ud) }
func TestMachineShell(t *testing.T)    { driver_test.MachineShell(t, ud) }
func TestMachineLogs(t *testing.T)     { driver_test.MachineLogs(t, ud) }
func TestListMachines(t *testing.T)    { driver_test.ListMachines(t, ud) }
func TestMachineInfo(t *testing.T)     { driver_test.MachineInfo(t, ud) }

func TestMain(tm *testing.M) {
	conf := make(map[string]interface{})
	err := ud.SetupDriver(conf)
	if err != nil {
		// need to work out how to fail here TODO
		fmt.Println(err)
		os.Exit(1)
	}
	err = ud.StartMachine(driver_test.TestMachine)
	if err != nil {
		// need to work out how to fail here TODO
		fmt.Println(err)
		os.Exit(1)
	}
	c := tm.Run()
	err = ud.HaltMachine(driver_test.TestMachine, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = ud.RemoveMachine(driver_test.TestMachine)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(c)
}
