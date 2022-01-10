package koble

type DriverConfig struct {
	// Name of driver to use
	Name      string                 `mapstructure:"name"`
	ExtraConf map[string]interface{} `mapstructure:"extra,remain"`
}

type Config struct {
	Driver DriverConfig `mapstructure:"driver"`
	// Name of which terminal to use
	// this terminal must be one of the default terminals or in
	// the user defined terms list
	// default is gnome
	Terminal string `mapstructure:"terminal"`
	// Whether to launch a terminal for start, attach and shell commands
	// default is true
	LaunchTerms bool `mapstructure:"launch_terms"`
	// Whether to launch a shell over tty attach on lab start
	// this only takes effect is LaunchTerms is true
	// default is false
	LaunchShell bool `mapstructure:"launch_shell"`
	// List of additional terminals to be chosen from by user
	// these will override the default terminals if the same
	// names are used
	Terms []Terminal `mapstructure:"terminals"`
	// Use plain output, e.g. no spinners, no prompts
	// default is false
	NonInteractive bool `mapstructure:"noninteractive"`
	// Do not use colour in output
	// default is false
	NoColor bool `mapstructure:"nocolor"`
	// namespace to use when not in a lab
	// default is "GLOBAL"
	Namespace string `mapstructure:"namespace"`
	// Amount of memory in MB to use for each machine
	// default is 128
	MachineMemory int `mapstructure:"machine_memory"`
}

type Terminal struct {
	Name    string   `mapstructure:"name"`
	Command []string `mapstructure:"command"`
}

var defaultTerms = []Terminal{
	{
		Name:    "alacritty",
		Command: []string{"alacritty", "-e"},
	},
	{
		Name:    "konsole",
		Command: []string{"konsole", "-e"},
	},
	{
		Name:    "gnome",
		Command: []string{"gnome-terminal", "--"},
	},
	{
		Name:    "kitty",
		Command: []string{"kitty"},
	},
	{
		Name:    "xterm",
		Command: []string{"xterm", "-e"},
	},
}
