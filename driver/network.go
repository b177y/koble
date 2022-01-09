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
	External bool   `default:"false" json:"external,omitempty"`
	Gateway  string `default:"" json:"gateway,omitempty"`
	IpRange  string `default:"" json:"iprange,omitempty"`
	Subnet   string `default:"" json:"subnet,omitempty"`
	IPv6     string `default:"" json:"ipv6,omitempty"`
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
