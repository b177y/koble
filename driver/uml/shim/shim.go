package shim

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/creack/pty"
	"github.com/docker/docker/pkg/broadcaster"
	"github.com/docker/docker/pkg/reexec"
)

var IMPORT = ""

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
	stdOutBr := new(broadcaster.Unbuffered)
	// TODO add logger chan, to send to MACHINE.log
	// if "Welcome To Netkit" seen, change machine status to running
	go io.Copy(stdOutBr, ptmx)
	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[INFO] New connection from %s.\n", c.LocalAddr())
			shimLog(fmt.Sprintf("[INFO] New connection from %s", c.LocalAddr()), dir, nil)
			go func() {
				stdOutBr.Add(c)
				go io.Copy(ptmx, c)
			}()
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
	reexec.Register("umlShim", runShim)
	if reexec.Init() {
		os.Exit(0)
	}
}
