package driver_test

import (
	"errors"

	"github.com/b177y/koble/pkg/driver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DeclareStartMachineTests(d driver.Driver) bool {
	return Describe("start machine", func() {
		m, err := d.Machine("testmachine", "teststartns")
		Expect(err).ShouldNot(HaveOccurred())
		AfterEach(func() {
			err := m.WaitUntil(60, driver.BootedState(), driver.ExitedState())
			Expect(err).ShouldNot(HaveOccurred())
			err = m.Stop(true)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.Remove()
			if !errors.Is(err, driver.ErrNotExists) {
				Expect(err).ShouldNot(HaveOccurred())
			}
		})
		It("start a machine", func() {
			err := m.Start(nil)
			Expect(err).ShouldNot(HaveOccurred())
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
		m, err := d.Machine("testmachine", "testhaltns")
		Expect(err).ShouldNot(HaveOccurred())
		BeforeEach(func() {
			err := m.Start(nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.BootedState(), driver.ExitedState())
			Expect(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			err := m.Remove()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("halt a machine gracefully", func() {
			err := m.Stop(false)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m.Running()).To(BeFalse())
			info, err := m.Info()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(info.ExitCode).To(BeZero())
		})
		It("halt a machine forcefully", func() {
			err := m.Stop(true)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m.Running()).To(BeFalse())
			info, err := m.Info()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(info.ExitCode).ToNot(BeZero())
		})
	})
}

func DeclareRemoveMachineTests(d driver.Driver) bool {
	return Describe("remove machine", func() {
		m, err := d.Machine("testmachine", "testremovens")
		Expect(err).ShouldNot(HaveOccurred())
		BeforeEach(func() {
			err := m.Start(nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.BootedState(), driver.ExitedState())
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("remove a stopped machine", func() {
			err := m.Stop(false)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.Remove()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("remove a force stopped machine", func() {
			err := m.Stop(true) // force is true here
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.Remove()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("remove a running machine", func() {
			err := m.Remove()
			Expect(err).Should(HaveOccurred())
			// cleanup
			err = m.Stop(true)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.Remove()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
}

// Test MachineExists
func DeclareExistsMachineTests(d driver.Driver) bool {
	return Describe("check machine exists", func() {
		m, err := d.Machine("testmachine", "testexistsns")
		Expect(err).ShouldNot(HaveOccurred())
		BeforeEach(func() {
			err := m.Start(nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.BootedState(), driver.ExitedState())
			Expect(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			exists, err := m.Exists()
			Expect(err).ShouldNot(HaveOccurred())
			if exists {
				m.Stop(true)
				err = m.WaitUntil(60, driver.ExitedState(), nil)
				Expect(err).ShouldNot(HaveOccurred())
				err = m.Remove()
				Expect(err).ShouldNot(HaveOccurred())
			}
		})
		It("running machine exists", func() {
			exists, err := m.Exists()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
		It("stopped machine exists", func() {
			err := m.Stop(false)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			exists, err := m.Exists()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
		It("force stopped machine exists", func() {
			err := m.Stop(true)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			exists, err := m.Exists()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})
		It("removed machine does not exist", func() {
			err := m.Stop(true)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.Remove()
			Expect(err).ShouldNot(HaveOccurred())
			exists, err := m.Exists()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})
		It("non existent machine does not exist", func() {
			// TODO move to another 'It' as it wastes time
			// starting and stopping a test machine
			m, err := d.Machine("nonexistent", "testexistsns")
			Expect(err).ShouldNot(HaveOccurred())
			exists, err := m.Exists()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})
	})
}

// Test getmachinestate
func DeclareGetStateMachineTests(d driver.Driver) bool {
	return Describe("get machine state", func() {
		m, err := d.Machine("testmachine", "teststatens")
		Expect(err).ShouldNot(HaveOccurred())
		BeforeEach(func() {
			err := m.Start(nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.BootedState(), driver.ExitedState())
			Expect(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			m.Stop(true)
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			err := m.Remove()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("running machine check state", func() {
			Expect(m.Running()).To(BeTrue())
		})
		It("stopped machine check state", func() {
			err := m.Stop(false)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.ExitedState(), nil)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(m.Running()).To(BeFalse())
		})
	})
}

// Test attach https://github.com/containers/podman/blob/c234c20a70304d526952f167c7c00122e5d54267/pkg/bindings/test/attach_test.go
func DeclareAttachMachineTests(d driver.Driver) bool {
	return Describe("attach to machine", func() {
		m, err := d.Machine("testmachine", "testattachns")
		Expect(err).ShouldNot(HaveOccurred())
		BeforeEach(func() {
			err := m.Start(nil)
			Expect(err).ShouldNot(HaveOccurred())
			err = m.WaitUntil(60, driver.BootedState(), driver.ExitedState())
			Expect(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			m.Stop(true)
			err := m.Remove()
			Expect(err).ShouldNot(HaveOccurred())
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
