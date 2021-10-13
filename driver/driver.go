// defines driver interface for netkit compatible drivers
package driver

type Machine struct {
	Name     string
	Hostlab  string
	Hosthome string
	Networks []string
	Image    string
}

type Network struct {
	Name     string
	External bool
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

	MachineExists(name string) (exists bool, err error)
	GetMachineState(name string) (state struct{}, err error)
	AttachToMachine(name string) (err error)
	MachineExecShell(name, command, user string,
		detach bool, workdir string) (err error)
	GetMachineLogs(name string, stdoutChan, stderrChan chan string,
		follow bool, tail int) (err error)

	NetworkExists(name string) (exists bool, err error)
	CreateNetwork(net Network) (id string, err error)
	StartNetwork(name string) (err error)
	StopNetwork(name string) (err error)
	RemoveNetwork(name string) (err error)
	GetNetworkState(name string) (state string, err error)
}
