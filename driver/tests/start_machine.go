package driver_test

import (
	"fmt"

	"github.com/b177y/netkit/driver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// var TestMachine = driver.Machine{
// 	Name:      "NETKITTESTING",
// 	Namespace: "TESTINGNS",
// }

func DeclareStartMachineTests(d driver.Driver) bool {
	return Describe("start machine", func() {
		BeforeEach(func() {
			// check machine isnt running
			fmt.Println("Running before each for")
		})
		AfterEach(func() {
			// teardown
			fmt.Println("Running after each for")
		})
		It("start a machine", func() {
			// start machine
			fmt.Println("Starting machine for")
			Expect(0).To(BeZero())
		})
	})
}

// func StartMachine(t *testing.T, d driver.Driver) {
// 	exists, err := d.MachineExists(TestMachine)
// 	if err != nil {
// 		t.Error(err)
// 	} else if !exists {
// 		t.Errorf("Machine %s does not exist after StartMachine call.\n",
// 			TestMachine.Name)
// 	}
// }

// func HaltMachine(t *testing.T, d driver.Driver) {
// 	m := driver.Machine{
// 		Name:      "NETKITTESTHALT",
// 		Namespace: "NETKITTESTING",
// 	}
// 	err := d.StartMachine(m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = d.HaltMachine(m, false)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	state, err := d.GetMachineState(m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if state.Running {
// 		t.Errorf("Machine %s is still running after Halt.", m.Name)
// 	}
// 	err = d.RemoveMachine(m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func ForceHalt(t *testing.T, d driver.Driver) {
// 	m := driver.Machine{
// 		Name:      "NETKITTESTHALT",
// 		Namespace: "NETKITTESTING",
// 	}
// 	err := d.StartMachine(m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = d.HaltMachine(m, true)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	state, err := d.GetMachineState(m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if state.Running {
// 		t.Errorf("Machine %s is still running after Halt.", m.Name)
// 	}
// 	err = d.RemoveMachine(m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func GetMachineState(t *testing.T, d driver.Driver) {
// 	state, err := d.GetMachineState(TestMachine)
// 	if err != nil {
// 		t.Error(err)
// 	} else if !state.Running {
// 		t.Errorf("Machine %s should be running", TestMachine.Name)
// 	}
// }

// func AttachToMachine(t *testing.T, d driver.Driver) {
// 	t.Errorf("Test not made")
// }

// func MachineShell(t *testing.T, d driver.Driver) {
// 	t.Errorf("Test not made")
// }

// func MachineLogs(t *testing.T, d driver.Driver) {
// 	t.Errorf("Test not made")
// }

// func ListMachines(t *testing.T, d driver.Driver) {
// 	t.Errorf("Test not made")
// }

// func MachineInfo(t *testing.T, d driver.Driver) {
// 	t.Errorf("Test not made")
// }
