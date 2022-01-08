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
	main     string
	finished string
	status   string
	charset  []string
	spinchar int
	done     bool
	err      error
}

func (s *Spinner) String() string {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.done {
		return green(" ✔") + " " + s.finished
	} else if s.err != nil {
		return red(" ✗") + " " + s.err.Error()
	}
	return fmt.Sprintf(" %s %s : %s", cyan(s.charset[s.spinchar]), s.main, s.status)
}

func (s *Spinner) Error(err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.err = err
	s.done = true
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

func (s *Spinner) Finished() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.done = true
}

func NewSpinner(main string, finished string) *Spinner {
	return &Spinner{
		mtx:      &sync.RWMutex{},
		main:     main,
		finished: finished,
		charset:  SPINCHARS,
		spinchar: 0,
	}
}
