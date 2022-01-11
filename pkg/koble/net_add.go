package koble

import (
	"errors"
	"fmt"
	"net"

	"github.com/b177y/koble/driver"
	"github.com/go-playground/validator/v10"
)

// redo with new viper config
func AddNetworkToLab(name string, external bool, gateway net.IP, subnet net.IPNet, ipv6 bool) error {
	if gateway.String() != "<nil>" {
		if subnet.IP == nil {
			return errors.New("To use a specified gateway you need to also specify a subnet.")
		} else if !subnet.Contains(gateway) {
			return fmt.Errorf("Gateway %s is not in subnet %s.", gateway.String(), subnet.String())
		}
	}
	lab := Lab{}
	// exists, err := GetLab(&lab)
	// if err != nil {
	// 	return err
	// }
	// if !exists {
	// 	return errors.New("lab.yml does not exist, are you in a lab directory?")
	// }
	err := validator.New().Var(name, "alphanum,max=30")
	if err != nil {
		return err
	}
	for nn := range lab.Networks {
		if nn == name {
			return fmt.Errorf("A network with the name %s already exists.", name)
		}
	}
	net := driver.NetConfig{
		External: external,
		//Gateway:  gateway,
		Subnet: subnet.String(),
		//IPv6:     ipv6,
	}

	if net.Subnet == "<nil>" {
		net.Subnet = ""
	}
	lab.Networks[name] = net
	err = SaveLab(&lab)
	if err != nil {
		return err
	}
	fmt.Printf("Created new network %s.\n", name)
	return nil
}
