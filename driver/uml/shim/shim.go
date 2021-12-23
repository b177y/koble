package shim

// https://programmerall.com/article/88412125350/
// https://github.com/moby/moby/blob/master/container/stream/streams.go

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/creack/pty"
	"github.com/docker/docker/pkg/broadcaster"
)

// Custom io.WriteCloser to check for netkit ready status
type ReadyChecker struct {
	Dir string
}

func (rc *ReadyChecker) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("Welcome to Netkit")) {
		err := updateState(rc.Dir, "running", 0)
		if err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

func (rc *ReadyChecker) Close() error {
	return nil
}

func updateState(dir, state string, exitCode int) error {
	err := os.WriteFile(filepath.Join(dir, "state"), []byte(state), 0600)
	if err != nil {
		return err
	}
	if state == "exited" {
		ec := []byte(fmt.Sprint(exitCode))
		err := os.WriteFile(filepath.Join(dir, "exitcode"), ec, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func shimLog(msg, dir string, err error) error {
	fn := filepath.Join(dir, "umlshim.log")
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(msg + "\n")
	return err
}

func RunShim() {
	dir := os.Args[1]
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		shimLog("Failed to make dir: ", dir, err)
		log.Fatal(err)
	}
	kern := os.Args[2]
	kernArgs := os.Args[3:]
	cmd := exec.Command(kern, kernArgs...)
	fmt.Println("Starting exec of netkit kernel")
	err = shimLog(fmt.Sprintf("starting kernel with %s %s\n", kern, kernArgs), dir, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write to shim log: %s", err.Error())
	}
	ptmx, err := pty.Start(cmd)
	if err != nil {
		shimLog("Failed to start pty command: ", dir, err)
		log.Fatal(err)
	}
	updateState(dir, "booting", 0)
	sockpath := filepath.Join(dir, "attach.sock")
	l, err := net.Listen("unix", sockpath)
	if err != nil {
		shimLog(fmt.Sprintf("Failed to start listen on sock (%s) ", sockpath), dir, err)
		log.Fatal(err)
	}
	logFile, err := os.OpenFile(filepath.Join(dir, "machine.log"),
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	stdOutBr := new(broadcaster.Unbuffered)
	stdOutBr.Add(logFile)
	rc := &ReadyChecker{
		Dir: dir,
	}
	stdOutBr.Add(rc)
	ctx, cancel := context.WithCancel(context.Background())
	go io.Copy(stdOutBr, ptmx)
	go func(ctx context.Context) {
		defer stdOutBr.Clean()
		defer l.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				c, err := l.Accept()
				if err != nil {
					c.Close()
					continue
				}
				fmt.Printf("[INFO] New connection from %s.\n", c.LocalAddr())
				shimLog(fmt.Sprintf("[INFO] New connection from %s", c.LocalAddr()), dir, nil)
				go func() {
					defer func() {
						c.Close()
					}()
					stdOutBr.Add(c)
					io.Copy(ptmx, c)
				}()
			}
		}
	}(ctx)
	err = cmd.Wait()
	cancel()
	ptmx.Close()
	if err != nil {
		log.Println("Error running uml", err)
		shimLog("Command wait() error", dir, err)
	}
	updateState(dir, "exited", cmd.ProcessState.ExitCode())
}
