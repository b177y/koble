package driver

import (
	"time"

	"github.com/cri-o/ocicni/pkg/ocicni"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

// Information about a machine available from the driver
type MachineInfo struct {
	Name      string               `json:"name"`
	Namespace string               `json:"namespace"`
	Lab       string               `json:"lab,omitempty"`
	Image     string               `json:"image"`
	Pid       int                  `json:"pid,omitempty"`
	State     string               `json:"state"`
	Status    string               `json:"status,omitempty"`
	Running   bool                 `json:"running"`
	StartedAt time.Time            `json:"started_at,omitempty"`
	Mounts    []string             `json:"mounts"`
	Networks  []string             `json:"networks"`
	Ports     []ocicni.PortMapping `json:"ports"`
	Uptime    string               `json:"uptime,omitempty"`
	ExitCode  int32                `json:"exit_code,omitempty"`
	ExitedAt  int64                `json:"exited_at,omitempty"`
}

type MachineState struct {
	Exists   bool
	State    *string
	Running  *bool
	ExitCode *int32
}

type MachineConfig struct {
	Image        string                 `default:"" yaml:"image,omitempty"`
	HostHome     bool                   `default:"false" yaml:"hosthome,omitempty"`
	Hostlab      string                 `default:"" yaml:"hostlab,omitempty"`
	Networks     []string               `default:"[]" yaml:"networks,omitempty"`
	Volumes      []spec.Mount           `default:"[]" yaml:"volumes,omitempty"`
	Dependencies []string               `default:"[]" yaml:"depends_on,omitempty"`
	DriverExtra  map[string]interface{} `default:"{}" yaml:"driver,omitempty"`
	Lab          string                 `default:"" yaml:"-"`
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
	State() (MachineState, error)
	Networks() ([]Network, error)
	Start(*MachineConfig) error
	Stop(force bool) error
	Remove() error
	Attach(opts *AttachOptions) error
	Exec(command string,
		opts *ExecOptions) error
	Shell(opts *ShellOptions) error
	Logs(opts *LogOptions) error
	WaitUntil(timeout time.Duration, target, failOn *MachineState) error
}
