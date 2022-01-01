package vecnet

import (
	"fmt"

	"github.com/containernetworking/plugins/pkg/ns"
)

// ip addresses
// https://github.com/containers/podman/blob/375ff223f430301edf25ef5a5f03a1ae1e029bef/libpod/network/internal/util/util.go
// https://github.com/containers/podman/blob/375ff223f430301edf25ef5a5f03a1ae1e029bef/libpod/network/internal/util/ip.go

func NetExists(name, namespace string) (exists bool, err error) {
	err = WithNetNS(namespace, func(ns.NetNS) error {
		exists, err = IfaceExistsByAlias("nbr_" + name)
		return err
	})
	return exists, err
}

// Create a new bridge within specified net namespace
func NewNet(name, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		return NewBridge("nbr_" + name)
	})
}

func RemoveNet(name, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		return DelBridge("nbr_" + name)
	})
}

func AddHostToNet(machine, network,
	namespace string) (ifaceName string, err error) {
	err = WithNetNS(namespace, func(ns.NetNS) error {
		tapAlias := fmt.Sprintf("mtap_%s_net_%s", machine, network)
		ifaceName, err = NewTap(tapAlias)
		if err != nil {
			return err
		}
		return AddTapToBridge(tapAlias, "nbr_"+network)
	})
	return ifaceName, err
}

func RemoveHostTap(machine, network, namespace string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		tapAlias := fmt.Sprintf("mtap_%s_net_%s", machine, network)
		return DelTap(tapAlias)
	})
}

func MakeNetExternal(network, namespace, subnet string) error {
	return WithNetNS(namespace, func(ns.NetNS) error {
		bridge := "nbr_" + network
		return AddSlirpIface("ext_"+network, bridge, namespace, subnet, "")
	})
}
