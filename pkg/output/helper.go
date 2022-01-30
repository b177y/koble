package output

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
	return toRun(out)
}
