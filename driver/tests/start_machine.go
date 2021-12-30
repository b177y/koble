package driver_test

import (
	"errors"

	"github.com/b177y/netkit/driver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DeclareStartMachineTests(d driver.Driver) bool {
	return Describe("start machine", func() {
		var m = driver.Machine{
			Name:      "testmachine",
			Namespace: "testns",
		}
		AfterEach(func() {
			err := d.WaitUntil(m, "running", 60)
			Expect(err).To(BeNil())
			err = d.HaltMachine(m, true)
			Expect(err).To(BeNil())
			err = d.RemoveMachine(m)
			if !errors.Is(err, driver.ErrNotExists) {
				Expect(err).To(BeNil())
			}
		})
		It("start a machine", func() {
			err := d.StartMachine(m)
			Expect(err).Should(BeNil())
		})
		// It("start a machine with invalid name", func() {
		// 	err := d.StartMachine(driver.Machine{
		// 		Name:      "machine@1234!",
		// 		Namespace: "testns",
		// 	})
		// 	Expect(err).Should(BeNil())
		// })
	})
}

func DeclareHaltMachineTests(d driver.Driver) bool {
	return Describe("halt machine", func() {
		var m = driver.Machine{
			Name:      "testmachine",
			Namespace: "testns",
		}
		BeforeEach(func() {
			err := d.StartMachine(m)
			Expect(err).To(BeNil())
			err = d.WaitUntil(m, "running", 60)
			Expect(err).To(BeNil())
		})
		AfterEach(func() {
			err := d.RemoveMachine(m)
			Expect(err).To(BeNil())
		})
		It("halt a machine gracefully", func() {
			err := d.HaltMachine(m, false)
			Expect(err).To(BeNil())
			err = d.WaitUntil(m, "exited", 60)
			Expect(err).To(BeNil())
			state, err := d.GetMachineState(m)
			Expect(err).Should(BeNil())
			Expect(state.Running).To(BeFalse())
			Expect(state.ExitCode).To(BeZero())
		})
		It("halt a machine forcefully", func() {
			err := d.HaltMachine(m, true)
			Expect(err).To(BeNil())
			err = d.WaitUntil(m, "exited", 60)
			Expect(err).To(BeNil())
			state, err := d.GetMachineState(m)
			Expect(err).Should(BeNil())
			Expect(state.Running).To(BeFalse())
			Expect(state.ExitCode).ToNot(BeZero())
		})
	})
}

func DeclareRemoveMachineTests(d driver.Driver) bool {
	return Describe("remove machine", func() {
		var m = driver.Machine{
			Name:      "testmachine",
			Namespace: "testns",
		}
		BeforeEach(func() {
			err := d.StartMachine(m)
			Expect(err).To(BeNil())
			err = d.WaitUntil(m, "running", 60)
			Expect(err).To(BeNil())
		})
		It("remove a stopped machine", func() {
			err := d.HaltMachine(m, false)
			Expect(err).To(BeNil())
			err = d.RemoveMachine(m)
			Expect(err).To(BeNil())
		})
		It("remove a force stopped machine", func() {
			err := d.HaltMachine(m, true) // force is true here
			Expect(err).To(BeNil())
			err = d.RemoveMachine(m)
			Expect(err).To(BeNil())
		})
		It("remove a running machine", func() {
			err := d.RemoveMachine(m)
			Expect(err).ToNot(BeNil())
			// cleanup
			err = d.HaltMachine(m, true)
			Expect(err).To(BeNil())
			err = d.RemoveMachine(m)
			Expect(err).To(BeNil())
		})
	})
}

// Test MachineExists
func DeclareExistsMachineTests(d driver.Driver) bool {
	return Describe("check machine exists", func() {
		var m = driver.Machine{
			Name:      "testmachine",
			Namespace: "testns",
		}
		BeforeEach(func() {
			err := d.StartMachine(m)
			Expect(err).ShouldNot(HaveOccurred())
			err = d.WaitUntil(m, "running", 60)
			Expect(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			d.HaltMachine(m, true)
			err := d.RemoveMachine(m)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("running machine exists", func() {
			exists, err := d.MachineExists(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
		It("stopped machine exists", func() {
			err := d.HaltMachine(m, false)
			Expect(err).ShouldNot(HaveOccurred())
			err = d.WaitUntil(m, "exited", 60)
			Expect(err).ShouldNot(HaveOccurred())
			exists, err := d.MachineExists(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
		It("force stopped machine exists", func() {
			err := d.HaltMachine(m, true)
			Expect(err).ShouldNot(HaveOccurred())
			err = d.WaitUntil(m, "exited", 60)
			Expect(err).ShouldNot(HaveOccurred())
			exists, err := d.MachineExists(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
		It("removed machine does not exist", func() {
			err := d.HaltMachine(m, true)
			Expect(err).ShouldNot(HaveOccurred())
			err = d.WaitUntil(m, "exited", 60)
			Expect(err).ShouldNot(HaveOccurred())
			err = d.RemoveMachine(m)
			Expect(err).ShouldNot(HaveOccurred())
			exists, err := d.MachineExists(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})
		It("non existent machine does not exist", func() {
			// TODO move to another 'It' as it wastes time
			// starting and stopping a test machine
			exists, err := d.MachineExists(driver.Machine{
				Name:      "nonexistent",
				Namespace: "testns",
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})
	})
}

// Test getmachinestate
func DeclareGetStateMachineTests(d driver.Driver) bool {
	return Describe("get machine state", func() {
		var m = driver.Machine{
			Name:      "testmachine",
			Namespace: "testns",
		}
		BeforeEach(func() {
			err := d.StartMachine(m)
			Expect(err).ShouldNot(HaveOccurred())
			err = d.WaitUntil(m, "running", 60)
			Expect(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			d.HaltMachine(m, true)
			err := d.RemoveMachine(m)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("running machine check state", func() {
			state, err := d.GetMachineState(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(state.Running).To(BeTrue())
		})
		It("running machine check state", func() {
			err := d.HaltMachine(m, false)
			Expect(err).ShouldNot(HaveOccurred())
			err = d.WaitUntil(m, "exited", 60)
			Expect(err).ShouldNot(HaveOccurred())
			state, err := d.GetMachineState(m)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(state.Running).To(BeFalse())
		})
	})
}

// Test attach https://github.com/containers/podman/blob/c234c20a70304d526952f167c7c00122e5d54267/pkg/bindings/test/attach_test.go
func DeclareAttachMachineTests(d driver.Driver) bool {
	return Describe("attach to machine", func() {
		var m = driver.Machine{
			Name:      "testmachine",
			Namespace: "testns",
		}
		BeforeEach(func() {
			err := d.StartMachine(m)
			Expect(err).To(BeNil())
			err = d.WaitUntil(m, "running", 60)
			Expect(err).To(BeNil())
		})
		AfterEach(func() {
			d.HaltMachine(m, true)
			err := d.RemoveMachine(m)
			Expect(err).ToNot(BeNil())
		})
		It("attach to running machine", func() {
		})
	})
}

// Test machine info

// Test list machines

// Test list namespaces

// Test shell
// Test exec

// Test Logs
