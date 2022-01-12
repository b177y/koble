package koble

type DriverConfig struct {
	// Name of driver to use
	Name      string                 `mapstructure:"name"`
	ExtraConf map[string]interface{} `mapstructure:"extra,remain"`
}

type Config struct {
	Driver DriverConfig `mapstructure:"driver"`
	// Whether to launch a terminal for start, attach and shell commands
	// default is true
	LaunchTerms bool `mapstructure:"launch_terms"`
	// Whether to launch a shell over tty attach on lab start
	// this only takes effect is LaunchTerms is true
	// default is false
	LaunchShell bool `mapstructure:"launch_shell"`
	// Terminal to use, additional terminals and options
	Terminal TermConfig `mapstructure:"terminal"`
	// Term option overrides
	TermOpts map[string]string `mapstructure:"term_opts"`
	// Use plain output, e.g. no spinners, no prompts
	// default is false
	NonInteractive bool `mapstructure:"noninteractive"`
	// Do not use colour in output
	// default is false
	NoColor bool `mapstructure:"nocolor"`
	// namespace to use when not in a lab
	// default is "GLOBAL"
	Namespace string `mapstructure:"namespace" validate:"alphanum,max=32"`
	// Amount of memory in MB to use for each machine
	// default is 128
	MachineMemory int `mapstructure:"machine_memory"`
}
