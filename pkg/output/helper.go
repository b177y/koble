package output

func WithSimpleContainer(title string, header func() string, plain bool,
	toRun func(c Container, o Output) error) (err error) {
	oc := NewContainer(header, plain)
	oc.Start()
	out := oc.AddOutput(title)
	out.Start()
	defer func() {
		if err != nil {
			out.Error(err)
		}
		oc.Stop()
	}()
	return toRun(oc, out)
}
