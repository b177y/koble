package output

import (
	"fmt"
	"io"
)

type LogOutput struct {
	out      io.Writer
	main     string
	finished string
}

func (lo *LogOutput) Write(p []byte) (int, error) {
	return fmt.Fprintf(lo.out, "[%s] %s\n", lo.main, string(p))
}

func (lo *LogOutput) Start() {
	fmt.Fprintf(lo.out, "%s\n", lo.main)
}
func (lo *LogOutput) Finished() {
	fmt.Fprintf(lo.out, "[%s] %s\n", lo.main, lo.finished)
}
func (lo *LogOutput) Error(err error) {
	fmt.Fprintf(lo.out, "[%s] %v\n", lo.main, err)
}

type LogContainer struct {
	Out        io.Writer
	Outputs    []*LogOutput
	headerFunc func() string
}

func (lc *LogContainer) AddOutput(main, finished string) Output {
	out := &LogOutput{
		out:      lc.Out,
		main:     main,
		finished: finished,
	}
	return out
}

func (lc *LogContainer) Start() {
	if lc.headerFunc != nil {
		fmt.Fprint(lc.Out, lc.headerFunc())
	}
}

func (lc *LogContainer) Stop() {}
