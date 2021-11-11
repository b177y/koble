package driver_test

import (
	"testing"

	"github.com/b177y/netkit/driver"
)

var TestMachine = driver.Machine{
	Name:      "NETKITTESTING",
	Namespace: "TESTINGNS",
}

func StartMachine(t *testing.T, d driver.Driver) {
	exists, err := d.MachineExists(TestMachine)
	if err != nil {
		t.Error(err)
	} else if !exists {
		t.Errorf("Machine %s does not exist after StartMachine call.\n",
			TestMachine.Name)
	}
}

func HaltMachine(t *testing.T, d driver.Driver) {
	m := driver.Machine{
		Name:      "NETKITTESTHALT",
		Namespace: "NETKITTESTING",
	}
	err := d.StartMachine(m)
	if err != nil {
		t.Error(err)
	}
	err = d.HaltMachine(m, false)
	if err != nil {
		t.Error(err)
	}
	state, err := d.GetMachineState(m)
	if err != nil {
		t.Error(err)
	}
	if state.Running {
		t.Errorf("Machine %s is still running after Halt.", m.Name)
	}
	err = d.RemoveMachine(m)
	if err != nil {
		t.Error(err)
	}
}

func ForceHalt(t *testing.T, d driver.Driver) {
	m := driver.Machine{
		Name:      "NETKITTESTHALT",
		Namespace: "NETKITTESTING",
	}
	err := d.StartMachine(m)
	if err != nil {
		t.Error(err)
	}
	err = d.HaltMachine(m, true)
	if err != nil {
		t.Error(err)
	}
	state, err := d.GetMachineState(m)
	if err != nil {
		t.Error(err)
	}
	if state.Running {
		t.Errorf("Machine %s is still running after Halt.", m.Name)
	}
	err = d.RemoveMachine(m)
	if err != nil {
		t.Error(err)
	}
}

func GetMachineState(t *testing.T, d driver.Driver) {
	state, err := d.GetMachineState(TestMachine)
	if err != nil {
		t.Error(err)
	} else if !state.Running {
		t.Errorf("Machine %s should be running", TestMachine.Name)
	}
}

func AttachToMachine(t *testing.T, d driver.Driver) {}

func MachineShell(t *testing.T, d driver.Driver) {}

func MachineLogs(t *testing.T, d driver.Driver) {}

func ListMachines(t *testing.T, d driver.Driver) {}

func MachineInfo(t *testing.T, d driver.Driver) {}
