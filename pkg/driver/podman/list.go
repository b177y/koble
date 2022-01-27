package podman

import (
	"github.com/b177y/koble/pkg/driver"
	"github.com/containers/podman/v3/pkg/bindings/containers"
)

func (pd *PodmanDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", namespace, pd.DriverName, all)
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.Conn, opts)
	if err != nil {
		return machines, err
	}
	for _, c := range ctrs {
		name, ns := getInfoFromLabels(c.Labels)
		m, err := pd.Machine(name, ns)
		if err != nil {
			return machines, err
		}
		info, err := m.Info()
		if err != nil {
			return machines, err
		}
		machines = append(machines, info)
	}
	return machines, nil
}
func (pd *PodmanDriver) ListAllNamespaces() (namespaces []string, err error) {
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", "", pd.DriverName, true)
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.Conn, opts)
	if err != nil {
		return namespaces, err
	}
	for _, c := range ctrs {
		_, ns := getInfoFromLabels(c.Labels)
		found := false
		for _, n := range namespaces {
			if ns == n {
				found = true
			}
		}
		if !found {
			namespaces = append(namespaces, ns)
		}

	}
	return namespaces, err
}
