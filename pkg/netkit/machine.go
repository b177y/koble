package netkit

import (
	"fmt"
	"os"

	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/podman"
	log "github.com/sirupsen/logrus"
)

func StartMachine(name, image string, networks []string) error {
	lab := Lab{
		Name: "",
	}
	exists, err := getLab(&lab)
	if err != nil {
		return err
	}
	if !exists {
		log.Warn("You are not in a lab directory, starting new non-lab machine.")
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
	_, err = d.StartMachine(m, lab.Name)
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

func MachineLogs(machine string, follow bool, tail int) error {
	lab := Lab{
		Name: "",
	}
	_, err := getLab(&lab)
	if err != nil {
		return err
	}
	d := new(podman.PodmanDriver)
	err = d.SetupDriver()
	if err != nil {
		return err
	}
	stdoutChan := make(chan string)
	stderrChan := make(chan string)
	go func() {
		for recv := range stdoutChan {
			fmt.Println(recv)
		}
	}()
	go func() {
		for recv := range stderrChan {
			fmt.Println(recv)
		}
	}()
	err = d.GetMachineLogs(machine, lab.Name, stdoutChan, stderrChan, follow, tail)
	return err
}

func ListMachines(all bool) error {
	lab := Lab{
		Name: "",
	}
	_, err := getLab(&lab)
	if err != nil {
		return err
	}
	d := new(podman.PodmanDriver)
	err = d.SetupDriver()
	if err != nil {
		return err
	}
	if !all {
		if lab.Name == "" {
			fmt.Println("Listing all machines which are not associated with a lab.")
			fmt.Printf("To see all machines use `netkit machine list --all`\n\n")
		} else {
			fmt.Printf("Listing all machines within this lab (%s).\n", lab.Name)
			fmt.Printf("To see all machines use `netkit machine list --all`\n\n")
		}
	}
	machines, err := d.ListMachines(lab.Name, all)
	// TODO - only show lab if using --all
	headers := []string{"Name", "Lab", "Image", "Networks", "State"}
	mlist, err := MachineInfoToStringArr(machines)
	if err != nil {
		return err
	}
	RenderTable(headers, mlist)
	return err
}
