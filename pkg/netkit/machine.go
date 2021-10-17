package netkit

import (
	"fmt"
	"os"

	"github.com/b177y/netkit/driver"
	log "github.com/sirupsen/logrus"
)

func (nk *Netkit) StartMachine(name, image string, networks []string) error {
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
	if err != nil {
		log.Fatal(err)
	}
	_, err = nk.Driver.StartMachine(m, nk.Lab.Name)
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

func (nk *Netkit) MachineLogs(machine string, follow bool, tail int) error {
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
	err := nk.Driver.GetMachineLogs(machine, nk.Lab.Name,
		stdoutChan, stderrChan, follow, tail)
	return err
}

func (nk *Netkit) ListMachines(all bool) error {
	if !all {
		if nk.Lab.Name == "" {
			fmt.Println("Listing all machines which are not associated with a lab.")
			fmt.Printf("To see all machines use `netkit machine list --all`\n\n")
		} else {
			fmt.Printf("Listing all machines within this lab (%s).\n", nk.Lab.Name)
			fmt.Printf("To see all machines use `netkit machine list --all`\n\n")
		}
	}
	machines, err := nk.Driver.ListMachines(nk.Lab.Name, all)
	var mlist [][]string
	var headers []string
	if all {
		mlist, headers = MachineInfoToStringArr(machines, true)
	} else {
		mlist, headers = MachineInfoToStringArr(machines, false)
	}
	RenderTable(headers, mlist)
	return err
}
