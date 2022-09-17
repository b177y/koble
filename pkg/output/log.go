package output

import (
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

type LogOutput struct {
	logger   *log.Logger
	name     string
	finished string
}

func (lo *LogOutput) Write(p []byte) (int, error) {
	lo.logger.Infof("%s: %s\n", lo.name, string(p))
	return len(p), nil
}

func (lo *LogOutput) Start() {
	lo.logger.Infoln(lo.name)
}

func (lo *LogOutput) Finished(msg string) {
	lo.logger.Infof("%s: finished: %s\n", lo.name, msg)
}

func (lo *LogOutput) Success(msg string) {
	lo.logger.Infof("%s: success: %s\n", lo.name, msg)
}

func (lo *LogOutput) Error(err error) {
	lo.logger.Errorf("%s: error: %v\n", lo.name, err)
}

type LogContainer struct {
	Out         io.Writer
	Outputs     []*LogOutput
	headerFunc  func(string) string
	titlePrefix string
}

func (lc *LogContainer) AddOutput(name string) Output {
	logger := log.New()
	logger.Out = lc.Out
	out := &LogOutput{
		name:   name,
		logger: logger,
	}
	return out
}

func (lc *LogContainer) Start() {
	if lc.headerFunc != nil {
		fmt.Fprint(lc.Out, lc.headerFunc(lc.titlePrefix))
	}
}

func (lc *LogContainer) Stop() {}
