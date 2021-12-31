package uml

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type process struct {
	pid     int
	cmdline string
}

func getProcesses() (pList []process, err error) {
	dirs, err := ioutil.ReadDir("/proc")
	if err != nil {
		// TODO change func signature to forget error and just use warns?
		return pList, err
	}
	for _, entry := range dirs {
		if pid, err := strconv.Atoi(entry.Name()); err == nil {
			cmdline, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
			if err != nil {
				log.Warnf("Could not read /proc/%d/cmdline: %v\n", pid, err)
				continue
			} else if strings.Contains(string(cmdline), "umlShim") {
				// we want to catch uml kernel processes not the shim
				continue
			} else if !strings.Contains(string(cmdline), "UMLNAMESPACE=") {
				continue
			}
			pgid, err := syscall.Getpgid(pid)
			if err != nil {
				log.Warnf("Could not get pgid for pid %d: %v\n", pid, err)
				return pList, nil
			} else if pgid != pid {
				continue
			}
			pList = append(pList, process{
				pid:     pid,
				cmdline: strings.TrimSuffix(string(cmdline), "\n"),
			})
		}
	}
	return pList, nil
}

func processBySubstring(substring ...string) int {
	processes, err := getProcesses()
	if err != nil {
		return -1
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
