package driver

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func BootedState() *MachineState {
	state := "running"
	running := true
	return &MachineState{
		Exists:  true,
		State:   &state,
		Running: &running,
	}
}

func BootingState() *MachineState {
	state := "booting"
	running := true
	return &MachineState{
		Exists:  true,
		State:   &state,
		Running: &running,
	}
}

func ExitedState() *MachineState {
	state := "exited"
	running := false
	return &MachineState{
		Exists:  true,
		State:   &state,
		Running: &running,
	}
}

func statesEqual(s1, s2 MachineState) bool {
	if s1.Exists != s2.Exists {
		return false
	}
	if s1.State != nil && s2.State != nil {
		if *s1.State != *s2.State {
			return false
		}
	}
	if s1.Running != nil && s2.Running != nil {
		if *s1.Running != *s2.Running {
			return false
		}
	}
	if s1.ExitCode != nil && s2.ExitCode != nil {
		if *s1.ExitCode != *s2.ExitCode {
			return false
		}
	}
	// all fields match so states are equal
	return true
}

func WaitUntil(m Machine, timeout time.Duration, target *MachineState,
	failOn *MachineState) error {
	if target == nil {
		return fmt.Errorf("cannot wait for state to equal nil target")
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	for {
		time.Sleep(200 * time.Millisecond)
		// check timeout
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("timed out waiting for %s to be in specified state: %w",
				m.Name(), err)
		}
		// get state
		state, err := m.State()
		if err != nil {
			log.Tracef("error getting state for %s: %w\n", m.Name(), err)
			continue
		}
		if state.State != nil {
			if *state.State == "failed" {
				return fmt.Errorf("machine %s has reached failed state", m.Name())
			}
		}
		// compare state to target
		if statesEqual(state, *target) {
			// fmt.Println("REACHED TARGET CONDITION", state, target)
			return nil
		}
		// compare state to failOn (if not nil)
		if failOn != nil {
			if statesEqual(state, *failOn) {
				// fmt.Println("REACHED FAIL ON CONDITION")
				return fmt.Errorf("machine state in wait reached failOn state")
			}
		}
	}
}
