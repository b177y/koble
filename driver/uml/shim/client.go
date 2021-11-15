package shim

import (
	"io"
	"net"
	"os"

	"github.com/moby/term"
	"golang.org/x/sync/errgroup"
)

var defaultEscapeSequence = []byte{16, 17}

func Attach(sock string) error {
	saved, err := term.MakeRaw(os.Stdin.Fd())
	if err != nil {
		return err
	}
	defer func() {
		term.RestoreTerminal(os.Stdin.Fd(), saved)
	}()
	conn, err := net.Dial("unix", sock)
	if err != nil {
		return err
	}
	stdinWithEscape := term.NewEscapeProxy(os.Stdin, defaultEscapeSequence)
	eg := new(errgroup.Group)
	eg.Go(func() error {
		defer conn.Close()
		_, err := io.Copy(conn, stdinWithEscape)
		if err.Error() == "read escape sequence" {
			return err
		} else if err != nil {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		defer conn.Close()
		io.Copy(os.Stdout, conn)
		if err != nil {
			return err
		}
		return nil
	})
	return eg.Wait()
}
