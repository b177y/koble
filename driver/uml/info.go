package uml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/b177y/koble/driver"
)

func (m *Machine) State() (state driver.MachineState, err error) {
	exists, err := m.Exists()
	if err != nil {
		return state, err
	} else if !exists {
		return driver.MachineState{Exists: false}, driver.ErrNotExists
	}
	state.Exists = true
	running := m.Running()
	state.Running = &running
	p, err := os.ReadFile(filepath.Join(m.mDir(), "state"))
	if err != nil {
		return state, err
	}
	stateString := string(p)
	state.State = &stateString
	if stateString == "exited" {
		ecBytes, err := os.ReadFile(filepath.Join(m.mDir(), "exitcode"))
		if err != nil {
			return state, err
		}
		ec, err := strconv.ParseInt(string(ecBytes), 10, 32)
		if err != nil {
			return state, err
		}
		exit := int32(ec)
		state.ExitCode = &exit
	}
	return state, nil
}

func (m *Machine) StartedAt() time.Time {
	stat, err := os.Stat(fmt.Sprintf("/proc/%d", m.Pid()))
	if err != nil {
		return time.Time{}
	}
	return stat.ModTime()
}

func saveInfo(dir string, info interface{}) error {
	configBytes, err := json.Marshal(info)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(dir, "config.json"),
		configBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func loadInfo(dir string, info interface{}) (err error) {
	content, err := ioutil.ReadFile(filepath.Join(dir, "config.json"))
	err = json.Unmarshal(content, &info)
	return err
}

func (m *Machine) Info() (info driver.MachineInfo, err error) {
	exists, err := m.Exists()
	if !exists {
		return driver.MachineInfo{}, driver.ErrNotExists
	}
	state, err := m.State()
	if err != nil {
		return info, err
	}
	if state.State != nil {
		info.State = *state.State
	} else {
		info.State = ""
	}
	if state.ExitCode != nil {
		info.ExitCode = *state.ExitCode
	} else {
		info.ExitCode = 0
	}
	if state.Running != nil {
		info.Running = *state.Running
	} else {
		info.Running = false
	}
	info.Pid = m.Pid()
	info.Name = m.Name()
	info.Namespace = m.namespace
	info.StartedAt = m.StartedAt()
	if info.State == "exited" {
		info.Status = fmt.Sprintf("exited (%d)", info.ExitCode)
	} else {
		info.Status = info.State
	}

	var saveInfo driver.MachineConfig
	err = loadInfo(m.mDir(), &saveInfo)
	if err != nil {
		return info, err
	}
	info.Networks = saveInfo.Networks
	info.Image = saveInfo.Image
	info.Lab = saveInfo.Lab

	return info, nil
}
