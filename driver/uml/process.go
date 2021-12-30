package uml

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"

	"github.com/b177y/netkit/driver"
)

type process struct {
	pid     int
	cmdline string
}

func processBySubstring(substring ...string) int {
	dirs, err := ioutil.ReadDir("/proc")
	if err != nil {
		return -1
	}
	var processes []process
	for _, entry := range dirs {
		if pid, err := strconv.Atoi(entry.Name()); err == nil {
			cmdline, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
			if err != nil {
				return -1
			}
			pgid, err := syscall.Getpgid(pid)
			if err != nil {
				return -1
			} else if pgid != pid {
				continue
			}
			processes = append(processes, process{
				pid:     pid,
				cmdline: strings.TrimSuffix(string(cmdline), "\n"),
			})
		}
	}
	for _, p := range processes {
		nonMatchFound := false
		for _, s := range substring {
			if !strings.Contains(p.cmdline, s) {
				nonMatchFound = true
				continue
			}
		}
		if !nonMatchFound {
			return p.pid
		}
	}
	return -1
}

func findMachineProcess(m driver.Machine) int {
	mHash := fmt.Sprintf("%x",
		md5.Sum([]byte(m.Name+"-"+m.Namespace)))
	return processBySubstring("umid="+mHash,
		"NETKITNAMESPACE="+m.Namespace)
}
