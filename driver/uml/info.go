package uml

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/b177y/netkit/driver"
)

func (m *Machine) State() string {
	mDir := filepath.Join(m.ud.RunDir, "machine", m.Id())
	stateFile := filepath.Join(mDir, "state")
	p, err := os.ReadFile(stateFile)
	if err != nil {
		return ""
	}
	return string(p)
}

func (m *Machine) Status() (string, int32) {
	status := m.State()
	if status != "exited" {
		return status, 0
	}
	mDir := filepath.Join(m.ud.RunDir, "machine", m.Id())
	ecFile := filepath.Join(mDir, "exitcode")
	p, err := os.ReadFile(ecFile)
	if err == nil {
		ec, err := strconv.ParseInt(string(p), 10, 32)
		if err == nil {
			return fmt.Sprintf("%s (%d)", status, ec), int32(ec)
		}
	}
	return fmt.Sprintf("%s (?)", status), 0
}

func (m *Machine) StartedAt() time.Time {
	stat, err := os.Stat(fmt.Sprintf("/proc/%d", m.Pid()))
	if err != nil {
		return time.Time{}
	}
	return stat.ModTime()
}

func (m *Machine) Info() (info driver.MachineInfo, err error) {
	exists, err := m.Exists()
	if !exists {
		return driver.MachineInfo{}, driver.ErrNotExists
	}
	info.Name = m.Name()
	info.Running = m.Running()
	info.Namespace = m.namespace
	info.State = m.State()
	info.Status, info.ExitCode = m.Status()
	info.Pid = m.Pid()
	info.StartedAt = m.StartedAt()

	// mDir := filepath.Join(m.ud.RunDir, "machine", m.Id())
	// content, err := ioutil.ReadFile(filepath.Join(mDir, "config.json"))
	// var mConfig driver.Machine
	// err = json.Unmarshal(content, &mConfig)
	// if err != nil {
	// 	return info, err
	// }
	// info.Networks = mConfig.Networks
	// info.Image = mConfig.Image
	// info.Lab = mConfig.Lab
	return info, nil
}
