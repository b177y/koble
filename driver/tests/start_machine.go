package driver_test

import (
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
		var m = driver.Machine{
			Name:      "testmachine",
			Namespace: "testns",
		}
		AfterEach(func() {
			err := d.HaltMachine(m, true)
			if err != nil {
				panic(err)
			}
			err = d.RemoveMachine(m)
			if err != nil {
				panic(err)
			}
		})
		It("start a machine", func() {
			err := d.StartMachine(m)
			Expect(err).Should(BeNil())
		})
	})
}

// Test HALT
// for non force check machine still exists, state is not running,
// exitcode should be 0
// Test Force Halt

// Test remove machine

// Test MachineExists

// Test getmachinestate

// Test attach
// Test shell
// Test exec

// Test Logs

// Test list machines

// Test machine info

// Test list namespaces
// Test
