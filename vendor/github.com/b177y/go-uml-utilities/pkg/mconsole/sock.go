package mconsole

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var (
	MCONSOLE_WRITETIMEOUT time.Duration = 3
	MCONSOLE_READTIMEOUT  time.Duration = 3
)

func min(x uint32, y uint32) uint32 {
	if x < y {
		return x
	}
	return y
}

func recvOutput(conn net.UnixConn) (output string, err error) {
	var reply mconsoleReply
	reply.More = 1
	for reply.More == 1 {
		respBytes := make([]byte, MCONSOLE_MAX_DATA+12)
		conn.SetReadDeadline(time.Now().Add(MCONSOLE_READTIMEOUT * time.Second))
		_, err := conn.Read(respBytes)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return "", fmt.Errorf("read socket timeout")
		} else if err != nil {
			return "", err
		}
		err = binary.Read(bytes.NewBuffer(respBytes), binary.LittleEndian, &reply)
		if err != nil {
			return "", err
		}
		if reply.Err != 0 {
			return "", fmt.Errorf("Error from mconsole: %d", reply.Err)
		}
		output += string(reply.Data[:])
	}
	return strings.Trim(output, "\n\r\x00"), err
}

// Sends a command to an open mconsole socket.
// Returns the output of the command.
func SendCommand(command string, conn net.UnixConn) (output string, err error) {
	req := mconsoleRequest{
		magic:   MCONSOLE_MAGIC,
		version: MCONSOLE_VERSION,
		length:  min(uint32(len(command)), MCONSOLE_MAX_DATA),
	}
	copy(req.data[:], []byte(command)[:req.length])
	req.data[req.length] = byte('\x00')
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, &req)
	if err != nil {
		return "", err
	}
	conn.SetWriteDeadline(time.Now().Add(MCONSOLE_WRITETIMEOUT * time.Second))
	_, err = conn.Write(buf.Bytes())
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return "", fmt.Errorf("write socket timeout")
	} else if err != nil {
		return "", err
	}
	return recvOutput(conn)
}

func openConn(sockpath string) (*net.UnixConn, error) {
	ra, err := net.ResolveUnixAddr("unixgram", sockpath)
	if err != nil {
		return nil, err
	}
	la, err := net.ResolveUnixAddr("unixgram", "@"+fmt.Sprint(os.Getpid())+"@@@@")
	if err != nil {
		return nil, err
	}
	return net.DialUnix("unixgram", la, ra)
}

// Opens an mconsole socket and sends a command.
// Returns the output of the command.
func CommandWithSock(command string,
	sockpath string) (output string, err error) {
	conn, err := openConn(sockpath)
	if err != nil {
		return "", err
	}
	defer func() {
		conn.Close()
	}()
	return SendCommand(command, *conn)
}
