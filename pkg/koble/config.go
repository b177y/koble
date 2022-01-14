package koble

import "time"

type DriverConfig struct {
	// Name of driver to use
	Name string `koanf:"name"`
	// extra config, top level key should be driver name
	ExtraConf map[string]interface{} `koanf:"extra,remain"`
}

type MachineConfig struct {
	// Amount of memory in MB to use for each machine
	// default is 128
	MachineMemory int `koanf:"memory"`
}

type LaunchConfig struct {
	MachineStart bool `koanf:"machine_start"`
	LabStart     bool `koanf:"lab_start"`
}

type Config struct {
	// Driver options
	Driver DriverConfig `koanf:"driver"`
	// Verbose (loglevel = Debug)
	Verbosity int `koanf:"verbose"`
	// Quiet (loglevel = error)
	Quiet bool `koanf:"quiet"`
	// Terminal to use, additional terminals and options
	Terminal TermConfig `koanf:"terminal"`
	// Term option overrides
	TermOpts map[string]string `koanf:"term_opts"`
	// whether to launch machines for machine / lab start
	Launch LaunchConfig `koanf:"launch"`
	// Use plain output, e.g. no spinners, no prompts
	// default is false
	NonInteractive bool `koanf:"noninteractive"`
	// Do not use colour in output
	// default is false
	NoColor bool `koanf:"nocolor"`
	// namespace to use when not in a lab
	// default is "GLOBAL"
	Namespace string `koanf:"namespace" validate:"alphanum,max=32"`
	// Wait (if -1 no wait, run in background) is how long in seconds to wait
	// for machines to startup / exit before giving timeout error
	// default is 300
	Wait time.Duration `koanf:"wait"`
	// Amount of memory in MB to use for each machine
	// default is 128
	Machine MachineConfig `koanf:"machine"`
}

var defaultConfig = map[string]interface{}{
	"driver.name":          "podman",
	"terminal.default":     "gnome",
	"terminal.exec":        "this",
	"launch.lab_start":     true,
	"launch.machine_start": true,
	"noninteractive":       false,
	"nocolor":              false,
	"namespace":            "GLOBAL",
	"wait":                 300,
	"machine.memory":       128,
}
