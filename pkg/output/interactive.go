package output

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/gosuri/uilive"
	"golang.org/x/crypto/ssh/terminal"
)

type InteractiveContainer struct {
	Out             io.Writer
	Spinners        []*Spinner
	RefreshInterval time.Duration
	lw              *uilive.Writer
	ticker          *time.Ticker
	tdone           chan bool
	mtx             *sync.RWMutex
	headerFunc      func(string) string
	titlePrefix     string
}

func (c *InteractiveContainer) AddOutput(name string) Output {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	spinner := NewSpinner(name)
	c.Spinners = append(c.Spinners, spinner)
	return spinner
}

func (c *InteractiveContainer) Listen() {
	for {
		c.mtx.Lock()
		interval := c.RefreshInterval
		c.mtx.Unlock()

		select {
		case <-time.After(interval):
			c.print()
		case <-c.tdone:
			c.print()
			close(c.tdone)
			return
		}
	}
}

func (c *InteractiveContainer) print() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.headerFunc != nil {
		fmt.Fprint(c.lw, c.headerFunc(c.titlePrefix))
	}
	width, _, err := terminal.GetSize(0)
	if err != nil {
		fmt.Fprintf(c.lw, "error getting term size: %v\n", err)
		return
	}
	for _, spinner := range c.Spinners {
		spinLine := strings.ReplaceAll(spinner.String(), "\n", "")
		if len(spinLine) > width {
			spinLine = spinLine[:width-4] + "..."
		}
		fmt.Fprintln(c.lw, spinLine)
		// fmt.Fprintln(c.lw, spinLine, []byte(spinLine))
	}
	c.lw.Flush()
}

func (c *InteractiveContainer) Start() {
	go c.Listen()
}

func (c *InteractiveContainer) Stop() {
	c.tdone <- true
	<-c.tdone
}
