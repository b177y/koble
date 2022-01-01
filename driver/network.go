package driver

type NetInfo struct {
	Name      string
	Namespace string
	External  bool
	Gateway   string
	IpRange   string
	Subnet    string
	IPv6      string
}

type NetCreateOptions struct {
	External bool   `default:"false"`
	Gateway  string `default:""`
	IpRange  string `default:""`
	Subnet   string `default:""`
	IPv6     string `default:""`
}

type Network interface {
	Name() string
	Id() string
	Create(opts *NetCreateOptions) error
	Start() error
	Stop() error
	Remove() error
	Exists() (bool, error)
	Running() (bool, error)
	Info() (NetInfo, error)
}
