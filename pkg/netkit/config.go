package netkit

type DriverConfig struct {
	Name      string      `yaml:"name"`
	ExtraConf interface{} `yaml:"extra"`
}

type Config struct {
	Driver    DriverConfig `yaml:"driver"`
	Terminal  string       `yaml:"terminal"`
	OpenTerms bool         `yaml:"open_terms"`
}
