package output

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

type InteractiveContainer struct {
	Out             io.Writer
	Spinners        []*Spinner
	RefreshInterval time.Duration
	lw              *uilive.Writer
	ticker          *time.Ticker
	tdone           chan bool
	mtx             *sync.RWMutex
	headerFunc      func() string
}

func (c *InteractiveContainer) AddOutput(main, finished string) Output {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	spinner := NewSpinner(main, finished)
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
		fmt.Fprint(c.lw, c.headerFunc())
	}
	for _, spinner := range c.Spinners {
		fmt.Fprintln(c.lw, spinner.String())
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
