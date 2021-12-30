package uml

import (
	"os"
	"path/filepath"

	"github.com/b177y/netkit/driver"
)

func (ud *UMLDriver) ListAllNamespaces() (namespaces []string, err error) {
	nsEntries, err := os.ReadDir(filepath.Join(ud.RunDir, "ns"))
	for _, n := range nsEntries {
		namespaces = append(namespaces, n.Name())
	}
	// TODO look for extra namespaces from pslist, filter by unique
	return namespaces, nil
}

func (ud *UMLDriver) ListMachinesForNamespace(namespace string) ([]driver.MachineInfo, error) {
	// TODO : 1 : PS list (with ns filter)
	// TODO : 2 : List Dir (for stopped machines)
	// TODO : 3 : for each entry append info to array
	// dir := filepath.Join(ud.RunDir, "ns", namespace)
	// entries, err := os.ReadDir(dir)
	// if err != nil {
	// 	return machines, err
	// }
	// for _, e := range entries {
	// 	sym, err := filepath.EvalSymlinks(filepath.Join(dir, e.Name()))
	// 	if err != nil {
	// 		os.RemoveAll(filepath.Join(dir, e.Name()))
	// 		continue
	// 	}
	// 	fInfo, err := os.Stat(sym)
	// 	if err != nil {
	// 		return machines, err
	// 	}
	// 	if !fInfo.IsDir() {
	// 		continue
	// 	}
	// 	info, err := ud.MachineInfo(driver.Machine{
	// 		Name:      e.Name(),
	// 		Namespace: namespace,
	// 	})
	// 	if err != nil && err != driver.ErrNotExists {
	// 		return machines, err
	// 	}
	// 	machines = append(machines, info)
	// }

}

func (ud *UMLDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	// TODO if all, look in all subdirs
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
