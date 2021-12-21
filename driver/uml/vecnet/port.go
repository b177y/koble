package vecnet

import (
	"encoding/json"
	"fmt"
	"net"
)

type PortMapping struct {
	Proto     string `json:"proto"`
	HostAddr  string `json:"host_addr"`
	GuestAddr string `json:"guest_addr"`
	HostPort  uint16 `json:"host_port"`
	GuestPort uint16 `json:"guest_port"`
}

type ForwardRequest struct {
	Execute   string      `json:"execute"`
	Arguments PortMapping `json:"arguments"`
}

func ForwardPort(port PortMapping, sockpath string) error {
	fr := ForwardRequest{
		Execute:   "add_hostfwd",
		Arguments: port,
	}
	conn, err := net.Dial("unix", sockpath)
	if err != nil {
		return err
	}
	defer conn.Close()
	data, err := json.Marshal(fr)
	if err != nil {
		return err
	}
	fmt.Println("Sending data to slirp", string(data))
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	err = conn.(*net.UnixConn).CloseWrite()
	if err != nil {
		return err
	}
	buf := make([]byte, 2048)
	lenRead, err := conn.Read(buf)
	if err != nil {
		return err
	}
	var response map[string]interface{}
	err = json.Unmarshal(buf[0:lenRead], &response)
	if err != nil {
		return err
	}
	if e, found := response["error"]; found {
		return fmt.Errorf("error from slirp4netns setting up port forwarding: %s", e)
	}
	return nil
}
