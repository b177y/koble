package netkit

type DriverConfig struct {
	Name      string      `yaml:"name"`
	ExtraConf interface{} `yaml:"extra"`
}

type Config struct {
	Driver    DriverConfig `yaml:"driver"`
	Terminal  string       `yaml:"terminal"`
	OpenTerms bool         `yaml:"open_terms"`
	Terms     []Terminal   `mapstructure:"terminals"`
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
