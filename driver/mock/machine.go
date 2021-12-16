package mock

import (
	"github.com/b177y/netkit/driver"
)

type MockDriver struct {
	Name         string
	DefaultImage string
}

func (md *MockDriver) SetupDriver(conf map[string]interface{}) (err error) {
	md.Name = "Mock Driver"
	md.DefaultImage = "MockImage"
	return nil
}

func (md *MockDriver) GetDefaultImage() string {
	return "mock-image"
}

func (md *MockDriver) MachineExists(m driver.Machine) (exists bool,
	err error) {
	return true, nil
}

func (md *MockDriver) StartMachine(m driver.Machine) (err error) {
	return nil
}

func (md *MockDriver) HaltMachine(m driver.Machine, force bool) error {
	return nil
}

func (md *MockDriver) RemoveMachine(m driver.Machine) error {
	return nil
}

func (md *MockDriver) GetMachineState(m driver.Machine) (state driver.MachineState, err error) {
	return state, nil
}

func (md *MockDriver) AttachToMachine(m driver.Machine) (err error) {
	return nil
}

func (md *MockDriver) MachineExecShell(m driver.Machine, command,
	user string, detach bool, workdir string) (err error) {
	return driver.ErrNotImplemented
}

func (md *MockDriver) GetMachineLogs(m driver.Machine,
	follow bool, tail int) (err error) {
	return nil
}

func (md *MockDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	return []driver.MachineInfo{}, nil
}

func (md *MockDriver) MachineInfo(m driver.Machine) (info driver.MachineInfo, err error) {
	return info, nil
}

func (md *MockDriver) ListAllNamespaces() (namespaces []string, err error) {
	return namespaces, nil
}
