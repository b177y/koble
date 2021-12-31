package uml

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/b177y/netkit/driver"
	log "github.com/sirupsen/logrus"
)

type mFound struct {
	name      string
	namespace string
}

func removeDuplicates(list []string) []string {
	keys := make(map[string]bool)
	filtered := []string{}
	for _, elem := range list {
		if _, ok := keys[elem]; !ok {
			keys[elem] = true
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

func (ud *UMLDriver) ListAllNamespaces() (namespaces []string, err error) {
	nsEntries, err := os.ReadDir(filepath.Join(ud.RunDir, "ns"))
	for _, n := range nsEntries {
		namespaces = append(namespaces, n.Name())
	}
	// TODO look for extra namespaces from pslist, filter by unique
	return namespaces, nil
}

func extractFromCmdline(cmdline, key string) (value string, err error) {
	parts := strings.Split(cmdline, "\x00")
	for _, p := range parts {
		if strings.HasPrefix(p, key+"=") {
			return strings.Replace(p, key+"=", "", 1), nil
		}
	}
	return "", fmt.Errorf("could not find %s within cmdline %s", key, cmdline)
}

func (ud *UMLDriver) ListMachinesForNamespace(namespace string) (machines []driver.MachineInfo,
	err error) {
	processList, err := getProcesses()
	if err != nil {
		return []driver.MachineInfo{}, err
	}
	var machinesFound []string
	// find machines from running processes
	for _, p := range processList {
		name, err := extractFromCmdline(p.cmdline, "name")
		if err != nil {
			// TODO WARN
			log.Warnf("Could not extract name from UML process (%d)", p.pid)
			continue
		}
		machinesFound = append(machinesFound, name)
	}
	// find machines from dir list
	// if these machines aren't running they wont have shown up in ps list
	dir := filepath.Join(ud.RunDir, "ns", namespace)
	entries, err := os.ReadDir(dir)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		// TODO WARN
		return []driver.MachineInfo{}, err
	} else if err == nil {
		for _, e := range entries {
			sym, err := filepath.EvalSymlinks(filepath.Join(dir, e.Name()))
			if err != nil {
				// TODO WARN
				os.RemoveAll(filepath.Join(dir, e.Name()))
				continue
			}
			fInfo, err := os.Stat(sym)
			if err != nil {
				// TODO WARN
				continue
			}
			if !fInfo.IsDir() {
				continue
			}
			machinesFound = append(machinesFound, e.Name())
		}
	}
	// filter out duplicates
	machinesFound = removeDuplicates(machinesFound)
	// for each entry of machinesFound get info
	for _, name := range machinesFound {
		m, err := ud.Machine(name, namespace)
		if err != nil {
			// TODO WARN
			continue
		}
		info, err := m.Info()
		if err != nil {
			// TODO WARN
			continue
		}
		machines = append(machines, info)
	}
	return machines, nil
}

func (ud *UMLDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	if all {
		namespaces, err := ud.ListAllNamespaces()
		if err != nil {
			return machines, err
		}
		for _, n := range namespaces {
			namespaceMachines, err := ud.ListMachines(n, false)
			if err != nil {
				return machines, err
			}
			machines = append(machines, namespaceMachines...)
		}
		return machines, nil
	} else {
		return ud.ListMachinesForNamespace(namespace)
	}
}