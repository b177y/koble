package upod

import (
	"github.com/b177y/koble/pkg/driver"
)

type Network struct {
	name      string
	namespace string
	ud        *UMLDriver
	p         driver.Network
}

func (n *Network) Id() string {
	return n.p.Id()
}

func (n *Network) Name() string {
	return n.p.Name()
}

func (n *Network) Create(opts *driver.NetConfig) (err error) {
	return n.p.Create(opts)
}

func (n *Network) Start() error {
	return n.p.Start()
}

func (n *Network) Remove() (err error) {
	return n.p.Remove()
}

func (n *Network) Stop() (err error) {
	return n.p.Stop()
}

func (n *Network) Running() (running bool, err error) {
	return n.p.Running()
}

func (ud *UMLDriver) ListNetworks(namespace string, all bool) (networks []driver.NetInfo, err error) {
	return ud.Podman.ListNetworks(namespace, all)
}

func (n *Network) Exists() (bool, error) {
	return n.p.Exists()
}

func (n *Network) Info() (nInfo driver.NetInfo, err error) {
	return n.p.Info()
}
