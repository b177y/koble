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
	Finished(string)
	Success(string)
	Error(error)
}

type Container interface {
	Start()
	Stop()
	AddOutput(name string) Output
}

func NewContainer(headerFunc func(string) string, titlePrefix string, plain bool) (c Container) {
	if plain {
		c = &LogContainer{
			Out:         os.Stdout,
			Outputs:     make([]*LogOutput, 0),
			headerFunc:  headerFunc,
			titlePrefix: titlePrefix,
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
			titlePrefix:     titlePrefix,
		}
	}
	return c
}
