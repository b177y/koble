package uml

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b177y/koble/driver"
	"github.com/b177y/koble/driver/uml/vecnet"
	"github.com/creasty/defaults"
)

type Network struct {
	name      string
	namespace string
	ud        *UMLDriver
}

func (n *Network) Id() string {
	return fmt.Sprintf("%x",
		md5.Sum([]byte(n.name+"."+n.namespace)))
}

func (n *Network) Name() string {
	return n.name
}

func (n *Network) Create(opts *driver.NetConfig) error {
	if opts == nil {
		opts = new(driver.NetConfig)
	}
	if err := defaults.Set(opts); err != nil {
		return err
	}
	exists, err := n.Exists()
	if err != nil {
		return fmt.Errorf("could not check if net %s exists: %w",
			n.name, err)
	}
	if exists {
		return driver.ErrExists
	}
	if err := vecnet.NewNet(n.name, n.namespace); err != nil {
		return err
	}
	if opts.External {
		return vecnet.MakeNetExternal(n.name, n.namespace, "")
	}
	netDir := filepath.Join(n.ud.Config.RunDir, "net", n.Id())
	if err := os.MkdirAll(netDir, 0744); err != nil {
		return err
	}
	return saveInfo(netDir, opts)
}

func (n *Network) Start() (err error) {
	return nil
}

func (n *Network) Remove() (err error) {
	return vecnet.RemoveNet(n.name, n.namespace)
}

func (n *Network) Stop() (err error) {
	return nil
}

func (n *Network) Running() (running bool, err error) {
	return n.Exists()
}

func (ud *UMLDriver) ListNetworks(namespace string,
	all bool) (networks []driver.NetInfo, err error) {
	return networks, nil
}

func (n *Network) Exists() (bool, error) {
	return vecnet.NetExists(n.name, n.namespace)
}

func (n *Network) Info() (nInfo driver.NetInfo, err error) {
	if exists, err := n.Exists(); err != nil {
		return nInfo, err
	} else if !exists {
		return nInfo, driver.ErrNotExists
	}
	var info driver.NetConfig
	err = loadInfo(filepath.Join(n.ud.Config.RunDir, "net", n.Id()), info)
	return nInfo, err
}
