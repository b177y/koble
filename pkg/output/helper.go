package output

import (
	"io"
	"os"
)

func WithSimpleContainer(title string, header func() string, plain bool,
	toRun func(o Output) error) (err error) {
	oc := NewContainer(header, plain)
	oc.Start()
	defer oc.Stop()
	return WithStdout(title, oc, toRun)
}

func WithStdout(title string, c Container, toRun func(o Output) error) (err error) {
	out := c.AddOutput(title)
	out.Start()
	origStdout := os.Stdout
	origStderr := os.Stderr
	r, w, _ := os.Pipe()
	rE, wE, _ := os.Pipe()
	defer func() {
		if err != nil {
			out.Error(err)
		}
		r.Close()
		w.Close()
		rE.Close()
		wE.Close()
		os.Stdout = origStdout
		os.Stderr = origStderr
	}()
	os.Stdout = w
	os.Stderr = wE
	go func() {
		io.Copy(out, r)
	}()
	go func() {
		io.Copy(out, rE)
	}()
	return toRun(out)
}
