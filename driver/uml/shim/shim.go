package shim

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/creack/pty"
	"github.com/docker/docker/pkg/reexec"
)

var IMPORT = ""

func handleConnection(c net.Conn, cOut, cIn chan []byte) {
	go func() {
		for {
			buf := make([]byte, 2048)
			c.Read(buf)
			cIn <- buf
		}
	}()
	go func() {
		for {
			buf := <-cOut
			br := bytes.NewReader(buf)
			io.Copy(c, br)
		}
	}()
}

type Broadcaster struct {
	clients []chan []byte
}

func (b *Broadcaster) SendAll(msg []byte) {
	for _, c := range b.clients {
		c <- msg
	}
}

func (b *Broadcaster) AddClient(newChan chan []byte) {
	b.clients = append(b.clients, newChan)
}

func shimLog(msg, dir string, err error) error {
	fn := filepath.Join(dir, "umlshim.log")
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(msg + "\n")
	return err
}

func runShim() {
	dir := os.Args[1]
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		shimLog("Failed to make dir: ", dir, err)
		log.Fatal(err)
	}
	kern := os.Args[2]
	kernArgs := os.Args[3:]
	cmd := exec.Command(kern, kernArgs...)
	// cmd := exec.Command("/bin/bash")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		shimLog("Failed to start pty command: ", dir, err)
		log.Fatal(err)
	}
	// TODO set DIR/status to booting
	sockpath := filepath.Join(dir, "attach.sock")
	l, err := net.Listen("unix", sockpath)
	if err != nil {
		shimLog(fmt.Sprintf("Failed to start listen on sock (%s) ", sockpath), dir, err)
		log.Fatal(err)
	}
	cIn := make(chan []byte)
	bc := new(Broadcaster)
	// TODO add logger chan, to send to MACHINE.log
	// if "Welcome To Netkit" seen, change machine status to running
	go func() {
		for {
			buf := make([]byte, 1)
			io.ReadAtLeast(ptmx, buf, 1)
			bc.SendAll(buf)
			shimLog(fmt.Sprintf("[STDOUT] %s", string(buf)), dir, nil)
		}
	}()
	go func() {
		for {
			buf := <-cIn
			br := bytes.NewReader(buf)
			io.Copy(ptmx, br)
		}
	}()
	go func() {
		for {
			fd, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[INFO] New connection from %s.\n", fd.LocalAddr())
			shimLog(fmt.Sprintf("[INFO] New connection from %s", fd.LocalAddr()), dir, nil)
			newChan := make(chan []byte)
			bc.AddClient(newChan)
			go handleConnection(fd, newChan, cIn)
		}
	}()
	err = cmd.Wait()
	l.Close()
	if err != nil {
		log.Println("Error running uml", err)
		shimLog("Command wait() error", dir, err)
	}
	// ec := cmd.ProcessState.ExitCode()
	// TODO write exit code to a file
}

func init() {
	// log.Printf("init start, os.Args = %+v\n", os.Args)
	reexec.Register("umlShim", runShim)
	if reexec.Init() {
		os.Exit(0)
	}
}
