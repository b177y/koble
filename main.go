package main

import (
	"fmt"
	"log"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/podman"
)

func main() {
	m := driver.Machine{
		Name:       "h12",
		Hostlab:    "/home/billy/repos/rootless-netkit/examples/lab04",
		Hosthome:   "/home/billy",
		Networks:   []string{},
		Filesystem: "localhost/netkit-deb-test",
	}

	fmt.Println("making new driver")
	d := new(podman.PodmanDriver)
	fmt.Println("setting up driver")
	err := d.SetupDriver()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting machine")
	_, err = d.StartMachine(m)
	err = d.ConnectToMachine("h12")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("We have a running container with id", id)
	// err := cmd.NetkitCLI.Execute()
	// if err != nil && err.Error() != "" {
	// 	fmt.Println(err)
	// }
}
