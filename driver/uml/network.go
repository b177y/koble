package uml

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/b177y/netkit/driver"
)

func (ud *UMLDriver) CreateNetwork(n driver.Network) (err error) {
	exists, err := ud.NetworkExists(n)
	if err != nil {
		return err
	}
	if exists {
		return driver.ErrExists
	}
	nHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(n.Name+"-"+n.Namespace)))
	netPath := filepath.Join(ud.RunDir, "network", nHash)
	err = os.MkdirAll(netPath, 0744)
	if err != nil && err != os.ErrExist {
		return err
	}
	cmd := exec.Command("uml_switch", "-hub", "-unix", filepath.Join(netPath, "hub.cnct"))
	return cmd.Start()
}

func (ud *UMLDriver) StartNetwork(net driver.Network) (err error) {
	return nil
}

func (ud *UMLDriver) RemoveNetwork(net driver.Network) (err error) {
	return err
}

func (ud *UMLDriver) StopNetwork(net driver.Network) (err error) {
	return nil
}

func (ud *UMLDriver) GetNetworkState(net driver.Network) (state driver.NetworkState,
	err error) {
	state.Running, err = ud.NetworkExists(net)
	return state, err
}

func (ud *UMLDriver) ListNetworks(lab string, all bool) (networks []driver.NetInfo, err error) {
	return networks, nil
}

func (ud *UMLDriver) NetworkExists(n driver.Network) (bool, error) {
	nHash := fmt.Sprintf("%x",
		sha256.Sum256([]byte(n.Name+"-"+n.Namespace)))
	hubPath := filepath.Join(ud.RunDir, "network", nHash, "hub.cnct")
	_, err := os.Stat(hubPath)
	if err != nil && err != os.ErrNotExist {
		return false, err
	} else if err == os.ErrNotExist {
		return false, nil
	} else {
		return true, nil
	}
}

func (ud *UMLDriver) NetInfo(net driver.Network) (nInfo driver.NetInfo, err error) {
	exists, err := ud.NetworkExists(net)
	if err != nil {
		return nInfo, err
	}
	if !exists {
		return nInfo, driver.ErrNotExists
	}
	return nInfo, err
}
