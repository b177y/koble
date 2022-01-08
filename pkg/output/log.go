package output

import (
	"fmt"
	"io"
)

type LogOutput struct {
	out      io.Writer
	name     string
	finished string
}

func (lo *LogOutput) Write(p []byte) (int, error) {
	return fmt.Fprintf(lo.out, "[%s] %s\n", lo.name, string(p))
}

func (lo *LogOutput) Start() {
	fmt.Fprintf(lo.out, "%s\n", lo.name)
}

func (lo *LogOutput) Finished(msg string) {
	fmt.Fprintf(lo.out, "[%s] finished: %s\n", lo.name, msg)
}

func (lo *LogOutput) Success(msg string) {
	fmt.Fprintf(lo.out, "[%s] success: %s\n", lo.name, msg)
}

func (lo *LogOutput) Error(err error) {
	fmt.Fprintf(lo.out, "[%s] error: %v\n", lo.name, err)
}

type LogContainer struct {
	Out        io.Writer
	Outputs    []*LogOutput
	headerFunc func() string
}

func (lc *LogContainer) AddOutput(name string) Output {
	out := &LogOutput{
		out:  lc.Out,
		name: name,
	}
	return out
}

func (lc *LogContainer) Start() {
	if lc.headerFunc != nil {
		fmt.Fprint(lc.Out, lc.headerFunc())
	}
}

func (lc *LogContainer) Stop() {}
