package driver

import (
	"time"

	"github.com/cri-o/ocicni/pkg/ocicni"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

// type Machine struct {
// 	Name         string `yaml:"name" validate:"alphanum,max=30"`
// 	Lab          string
// 	Namespace    string
// 	Hostlab      string
// 	Dependencies []string               `yaml:"depends_on,omitempty"`
// 	HostHome     bool                   `yaml:"hosthome,omitempty"`
// 	Networks     []string               `yaml:"networks,omitempty" validate:"alphanum,max=30"`
// 	Volumes      []spec.Mount           `yaml:"volumes,omitempty"`
// 	Image        string                 `yaml:"image,omitempty"`
// 	DriverExtra  map[string]interface{} `yaml:"driver_extra,omitempty"`
// }

type MachineInfo struct {
	Name      string
	Namespace string
	Image     string
	Pid       int
	State     string
	Status    string
	Running   bool
	StartedAt time.Time
	Mounts    []string
	Ports     []ocicni.PortMapping
	Uptime    string
	ExitCode  int32
	ExitedAt  int64
}

type StartOptions struct {
	Image       string
	HostHome    bool
	Hostlab     string
	Networks    []Network
	Volumes     []spec.Mount
	DriverExtra map[string]interface{}
	Lab         string
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
