package driver

// Information about a network available from the driver
type NetInfo struct {
	Name      string
	Namespace string
	External  bool
	Gateway   string
	IpRange   string
	Subnet    string
	IPv6      string
}

type NetConfig struct {
	External bool   `default:"false" mapstructure:"external,omitempty"`
	Gateway  string `default:"" mapstructure:"gateway,omitempty"`
	IpRange  string `default:"" mapstructure:"iprange,omitempty"`
	Subnet   string `default:"" mapstructure:"subnet,omitempty"`
	IPv6     string `default:"" mapstructure:"ipv6,omitempty"`
}

type Network interface {
	Name() string
	Id() string
	Create(opts *NetConfig) error
	Start() error
	Stop() error
	Remove() error
	Exists() (bool, error)
	Running() (bool, error)
	Info() (NetInfo, error)
}
