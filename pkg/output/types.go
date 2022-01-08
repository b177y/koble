package output

import (
	"os"
	"sync"
	"time"

	"github.com/gosuri/uilive"
)

type Output interface {
	Write(p []byte) (int, error)
	Start()
	Finished()
	Error(error)
}

type Container interface {
	Start()
	Stop()
	AddOutput(main, finished string) Output
}

func NewContainer(headerFunc func() string, plain bool) (c Container) {
	if plain {
		c = &LogContainer{
			Out:        os.Stdout,
			Outputs:    make([]*LogOutput, 0),
			headerFunc: headerFunc,
		}
	} else {
		lw := uilive.New()
		lw.Out = os.Stdout
		c = &InteractiveContainer{
			Out:             lw.Out,
			Spinners:        make([]*Spinner, 0),
			RefreshInterval: time.Millisecond * 10,
			tdone:           make(chan bool),
			lw:              lw,
			mtx:             &sync.RWMutex{},
			headerFunc:      headerFunc,
		}
	}
	return c
}
