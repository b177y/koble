package podman_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/podman"
	driver_test "github.com/b177y/netkit/driver/test"
)

var pd = new(podman.PodmanDriver)

func TestPodmanImplementsDriver(t *testing.T) {
	var d interface{} = pd
	v, ok := d.(driver.Driver)
	if !ok {
		fmt.Println(v, ok)
		t.Error("PodmanDriver does not satisify driver interface.")
	}

}

func TestHaltMachine(t *testing.T)     { driver_test.HaltMachine(t, pd) }
func TestForceHalt(t *testing.T)       { driver_test.ForceHalt(t, pd) }
func TestGetMachineState(t *testing.T) { driver_test.GetMachineState(t, pd) }
func TestAttachToMachine(t *testing.T) { driver_test.AttachToMachine(t, pd) }
func TestMachineShell(t *testing.T)    { driver_test.MachineShell(t, pd) }
func TestMachineLogs(t *testing.T)     { driver_test.MachineLogs(t, pd) }
func TestListMachines(t *testing.T)    { driver_test.ListMachines(t, pd) }
func TestMachineInfo(t *testing.T)     { driver_test.MachineInfo(t, pd) }

func TestMain(tm *testing.M) {
	conf := make(map[string]interface{})
	conf["uri"] = "unix://run/user/1000/podman/podman.sock"
	err := pd.SetupDriver(conf)
	if err != nil {
		// need to work out how to fail here TODO
		fmt.Println(err)
		os.Exit(1)
	}
	err = pd.StartMachine(driver_test.TestMachine)
	if err != nil {
		// need to work out how to fail here TODO
		fmt.Println(err)
		os.Exit(1)
	}
	c := tm.Run()
	err = pd.HaltMachine(driver_test.TestMachine, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = pd.RemoveMachine(driver_test.TestMachine)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(c)
}
