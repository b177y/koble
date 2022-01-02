package vecnet

// Copied from podman's https://github.com/containers/podman/blob/375ff223f430301edf25ef5a5f03a1ae1e029bef/libpod/network/internal/util/ip.go
// and https://github.com/containers/podman/blob/375ff223f430301edf25ef5a5f03a1ae1e029bef/libpod/network/internal/util/util.go

import (
	"errors"
	"fmt"
	"math/rand"
	"net"

	"github.com/sirupsen/logrus"
)

func networkIntersectsWithNetworks(n *net.IPNet, networklist []*net.IPNet) bool {
	for _, nw := range networklist {
		if networkIntersect(n, nw) {
			return true
		}
	}
	return false
}

func networkIntersect(n1, n2 *net.IPNet) bool {
	return n2.Contains(n1.IP) || n1.Contains(n2.IP)
}

func incByte(subnet *net.IPNet, idx int, shift uint) error {
	if idx < 0 {
		return errors.New("no more subnets left")
	}
	if subnet.IP[idx] == 255 {
		subnet.IP[idx] = 0
		return incByte(subnet, idx-1, 0)
	}
	subnet.IP[idx] += 1 << shift
	return nil
}

// nextSubnet returns subnet incremented by 1
func nextSubnet(subnet *net.IPNet) (*net.IPNet, error) {
	newSubnet := &net.IPNet{
		IP:   subnet.IP,
		Mask: subnet.Mask,
	}
	ones, bits := newSubnet.Mask.Size()
	if ones == 0 {
		return nil, fmt.Errorf("%s has only one subnet", subnet.String())
	}
	zeroes := uint(bits - ones)
	shift := zeroes % 8
	idx := ones/8 - 1
	if idx < 0 {
		idx = 0
	}
	if err := incByte(newSubnet, idx, shift); err != nil {
		return nil, err
	}
	return newSubnet, nil
}

// getFreeIPv4NetworkSubnet returns a unused ipv4 subnet
func getFreeIPv4NetworkSubnet(usedNetworks []*net.IPNet, startIp net.IP) (*net.IPNet, error) {
	network := &net.IPNet{
		IP:   startIp,
		Mask: net.IPMask{255, 255, 255, 0},
	}

	// TODO: make sure to not use public subnets
	for {
		if intersectsConfig := networkIntersectsWithNetworks(network, usedNetworks); !intersectsConfig {
			logrus.Debugf("found free ipv4 network subnet %s", network.String())
			return network, nil
		}
		var err error
		network, err = nextSubnet(network)
		if err != nil {
			return nil, err
		}
	}
}

// getRandomIPv6Subnet returns a random internal ipv6 subnet as described in RFC3879.
func getRandomIPv6Subnet() (net.IPNet, error) {
	ip := make(net.IP, 8, net.IPv6len)
	// read 8 random bytes
	_, err := rand.Read(ip)
	if err != nil {
		return net.IPNet{}, nil
	}
	// first byte must be FD as per RFC3879
	ip[0] = 0xfd
	// add 8 zero bytes
	ip = append(ip, make([]byte, 8)...)
	return net.IPNet{IP: ip, Mask: net.CIDRMask(64, 128)}, nil
}

// getFreeIPv6NetworkSubnet returns a unused ipv6 subnet
func getFreeIPv6NetworkSubnet(usedNetworks []*net.IPNet) (*net.IPNet, error) {
	// FIXME: Is 10000 fine as limit? We should prevent an endless loop.
	for i := 0; i < 10000; i++ {
		// RFC4193: Choose the ipv6 subnet random and NOT sequentially.
		network, err := getRandomIPv6Subnet()
		if err != nil {
			return nil, err
		}
		if intersectsConfig := networkIntersectsWithNetworks(&network, usedNetworks); !intersectsConfig {
			logrus.Debugf("found free ipv6 network subnet %s", network.String())
			return &network, nil
		}
	}
	return nil, errors.New("failed to get random ipv6 subnet")
}
