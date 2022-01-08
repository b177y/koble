package output

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

var SPINCHARS = []string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"}

var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()

type Spinner struct {
	mtx      *sync.RWMutex
	buf      *bytes.Buffer
	name     string
	finished string
	status   string
	charset  []string
	spinchar int
	done     bool
	success  bool
	err      error
}

func (s *Spinner) String() string {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.err != nil {
		return red(" ✗") + " " + s.err.Error()
	} else if s.done {
		return s.finished
	}
	return fmt.Sprintf(" %s %s : %s", cyan(s.charset[s.spinchar]), s.name, s.status)
}

func (s *Spinner) Error(err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.err = err
}

func (s *Spinner) Spin() {

	for {
		if s.spinchar == len(s.charset)-1 {
			s.spinchar = 0
		} else {
			s.spinchar = s.spinchar + 1
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func (s *Spinner) Write(p []byte) (n int, err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.status = string(p)
	return len(p), nil
}

func (s *Spinner) Start() {
	go s.Spin()
}

func (s *Spinner) Finished(msg string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.finished = cyan(" ✔") + " " + msg
	s.done = true
}

func (s *Spinner) Success(msg string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.finished = green(" ✔") + " " + msg
	s.done = true
}

func NewSpinner(name string) *Spinner {
	return &Spinner{
		mtx:      &sync.RWMutex{},
		name:     name,
		charset:  SPINCHARS,
		spinchar: 0,
	}
}
