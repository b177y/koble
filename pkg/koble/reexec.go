package koble

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"os"
	"strings"

	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/driver/podman"
	"github.com/b177y/koble/pkg/driver/uml"
	"github.com/b177y/koble/pkg/driver/upod"
)

func init() {
	// TODO find better solution to avoid importing these drivers in pkg/koble
	driver.RegisterDriver("podman", func() driver.Driver {
		return new(podman.PodmanDriver)
	})
	driver.RegisterDriver("uml", func() driver.Driver {
		return new(uml.UMLDriver)
	})
	driver.RegisterDriver("upod", func() driver.Driver {
		return new(upod.UMLDriver)
	})
	gob.Register(map[string]interface{}{})
	// gob.Register(&podman.PodmanDriver{})
	// gob.Register(&uml.UMLDriver{})
	// handleReexecFuncs["launchTerm"] = attachReexec
	if len(os.Args) >= 2 {
		if os.Args[1] == "launchTermReexec" {
			handleReexec()
			os.Exit(0)
		}
	}
}

func (nk *Koble) reexecAttach(machine string) (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&nk.Config)
	if err != nil {
		return "", err
	}
	b64nk := base64.StdEncoding.EncodeToString([]byte(buf.Bytes()))
	command := []string{os.Args[0], "launchTermReexec", b64nk, machine}
	return strings.Join(command, " "), nil
}

func handleReexec() {
	// args [1] == name of reexec func
	// args [2] == koble.Koble struct
	decodedb64nk, err := base64.StdEncoding.DecodeString(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	buf := bytes.NewBuffer(decodedb64nk)
	dec := gob.NewDecoder(buf)
	var nk Koble
	err = dec.Decode(&nk.Config)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	err = nk.processConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	nk.AttachToMachine(os.Args[3], "this")
}
