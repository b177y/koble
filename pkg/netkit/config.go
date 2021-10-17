package netkit

type DriverConfig struct {
	Name      string                 `mapstructure:"name"`
	ExtraConf map[string]interface{} `mapstructure:"extra,remain"`
}

type Config struct {
	Driver    DriverConfig `mapstructure:"driver"`
	Terminal  string       `mapstructure:"terminal"`
	OpenTerms bool         `mapstructure:"open_terms"`
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
