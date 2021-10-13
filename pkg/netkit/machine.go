package netkit

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/podman"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func StartMachine(name, image string, networks []string) error {
	exists, err := fileExists("lab.yml")
	if err != nil {
		return err
	}
	if !exists {
		log.Warn("You are not in a lab directory, starting new non-lab machine.")
	} else {
		f, err := ioutil.ReadFile("lab.yml")
		if err != nil {
			return err
		}
		lab := Lab{}
		err = yaml.Unmarshal(f, &lab)
		if err != nil {
			return err
		}
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	m := driver.Machine{
		Name:     name,
		Hostlab:  wd,
		Hosthome: home,
		Networks: []string{},
		Image:    image,
	}
	d := new(podman.PodmanDriver)
	err = d.SetupDriver()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting machine")
	_, err = d.StartMachine(m)
	return err
}

func CrashMachine() error {
	return nil
}

func HaltMachine() error {
	return nil
}

func DestroyMachine() error {
	return nil
}

func MachineInfo() error {
	return nil
}
