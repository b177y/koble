package output

import (
	"io"
	"os"
)

func WithSimpleContainer(title string, header func() string, plain bool,
	toRun func(c Container, o Output) error) (err error) {
	oc := NewContainer(header, plain)
	oc.Start()
	out := oc.AddOutput(title)
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
		oc.Stop()
	}()
	os.Stdout = w
	os.Stderr = wE
	go func() {
		io.Copy(out, r)
	}()
	go func() {
		io.Copy(out, rE)
	}()
	err = toRun(oc, out)
	if err != nil {
		return err
	}
	return nil
}
