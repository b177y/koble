package koble

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

func Reexec() {
	gob.Register(map[string]interface{}{})
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
	command := []string{os.Args[0], "launchTermReexec",
		b64nk, "attach", machine}
	return strings.Join(command, " "), nil
}

type reExecOpts struct {
	Command string
	User    string
	Detach  bool
	Workdir string
}

func (nk *Koble) reexecShell(machine, user, workdir string) (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&nk.Config)
	if err != nil {
		return "", err
	}
	b64nk := base64.StdEncoding.EncodeToString([]byte(buf.Bytes()))
	opts := reExecOpts{
		User:    user,
		Workdir: workdir,
	}
	var optsBuf bytes.Buffer
	optsEnc := gob.NewEncoder(&optsBuf)
	err = optsEnc.Encode(&opts)
	if err != nil {
		return "", err
	}
	optsB64 := base64.StdEncoding.EncodeToString([]byte(optsBuf.Bytes()))

	command := []string{os.Args[0], "launchTermReexec",
		b64nk, "shell", machine, optsB64}
	return strings.Join(command, " "), nil
}

func (nk *Koble) reexecExec(machine, command, user string,
	detach bool, workdir string) (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&nk.Config)
	if err != nil {
		return "", err
	}
	b64nk := base64.StdEncoding.EncodeToString([]byte(buf.Bytes()))
	opts := reExecOpts{
		Command: command,
		User:    user,
		Detach:  detach,
		Workdir: workdir,
	}
	var optsBuf bytes.Buffer
	optsEnc := gob.NewEncoder(&optsBuf)
	err = optsEnc.Encode(&opts)
	if err != nil {
		return "", err
	}
	optsB64 := base64.StdEncoding.EncodeToString([]byte(optsBuf.Bytes()))

	cmd := []string{os.Args[0], "launchTermReexec", b64nk, "exec",
		machine, optsB64}
	return strings.Join(cmd, " "), nil
}

func handleReexec() {
	// args[1] == "launchTermReexec"
	// args[2] == koble.Koble struct
	// args[3] == attach/shell/exec
	// args[4] == machine name
	// args[5] == user, args[6] == workdir (if shell)
	// args[5] == command, args[6] == user, args[7] == detach, args[8] == workdir (if exec)
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
	if cmd := os.Args[3]; cmd == "attach" {
		if len(os.Args) < 4 {
			fmt.Printf("args (%v) not formatted correctly\n", os.Args)
			os.Exit(1)
		}
		nk.AttachToMachine(os.Args[4], "this")
	} else if cmd == "shell" {
		if len(os.Args) < 6 {
			fmt.Printf("args (%v) not formatted correctly\n", os.Args)
			os.Exit(1)
		}
		decodedOpts, err := base64.StdEncoding.DecodeString(os.Args[5])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		buf := bytes.NewBuffer(decodedOpts)
		dec := gob.NewDecoder(buf)
		var opts reExecOpts
		err = dec.Decode(&opts)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		nk.Shell(os.Args[4], opts.User, opts.Workdir)
	} else if cmd == "exec" {
		if len(os.Args) < 6 {
			fmt.Printf("args (%v) not formatted correctly\n", os.Args)
			os.Exit(1)
		}
		decodedOpts, err := base64.StdEncoding.DecodeString(os.Args[5])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		buf := bytes.NewBuffer(decodedOpts)
		dec := gob.NewDecoder(buf)
		var opts reExecOpts
		err = dec.Decode(&opts)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		nk.Exec(os.Args[4], opts.Command, opts.User, opts.Detach, opts.Workdir)
	} else {
		fmt.Printf("Error: command %s doesn't exist for launchTermReexec\n", cmd)
		os.Exit(1)
	}
}
