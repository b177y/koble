// defines driver interface for netkit compatible drivers
package driver

type Machine struct {
	Name       string
	Hostlab    string
	Hosthome   string
	Networks   []string
	Filesystem string
}

type Network struct {
	Name     string
	Internal bool
	Gateway  string
	IpRange  string
	Subnet   string
	IPv6     string
}

type Driver interface {
	SetupDriver() (err error)

	StartMachine(m Machine) (id string, err error)
	StopMachine(name string) (err error)
	CrashMachine(name string) (err error)

	GetMachineState(name string) (state struct{}, err error)
	ConnectToMachine(name string) (err error)
	GetMachineLogs(name string, stdoutChan, stderrChan chan string) (err error)

	CreateNetwork(net Network) (id string, err error)
	StartNetwork(name string) (err error)
	StopNetwork(name string) (err error)
	RemoveNetwork(name string) (err error)
	GetNetworkState(name string) (state string, err error)
}
