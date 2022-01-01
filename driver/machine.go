package driver

import (
	"time"

	"github.com/cri-o/ocicni/pkg/ocicni"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

// Information about a machine available from the driver
type MachineInfo struct {
	Name      string
	Namespace string
	Lab       string
	Image     string
	Pid       int
	State     string
	Status    string
	Running   bool
	StartedAt time.Time
	Mounts    []string
	Networks  []string
	Ports     []ocicni.PortMapping
	Uptime    string
	ExitCode  int32
	ExitedAt  int64
}

type StartOptions struct {
	Image       string                 `default:""`
	HostHome    bool                   `default:"false"`
	Hostlab     string                 `default:""`
	Networks    []string               `default:"[]"`
	Volumes     []spec.Mount           `default:"[]"`
	DriverExtra map[string]interface{} `default:"{}"`
	Lab         string                 `default:""`
}

type ExecOptions struct {
	User    string
	Detach  bool
	Workdir string
}

type ShellOptions struct {
	User    string
	Workdir string
}

type LogOptions struct {
	Follow bool
	Tail   int
}

type AttachOptions struct{}

type Machine interface {
	Name() string
	Id() string
	Exists() (bool, error)
	Running() bool
	Info() (MachineInfo, error)
	Networks() ([]Network, error)
	Start(*StartOptions) error
	Stop(force bool) error
	Remove() error
	Attach(opts *AttachOptions) error
	Exec(command string,
		opts *ExecOptions) error
	Shell(opts *ShellOptions) error
	Logs(opts *LogOptions) error
	WaitUntil(state string, timeout time.Duration) error
}
