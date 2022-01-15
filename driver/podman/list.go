package podman

import (
	"github.com/b177y/koble/driver"
	"github.com/containers/podman/v3/pkg/bindings/containers"
)

func (pd *PodmanDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", namespace, all) // TODO get namespace here
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.conn, opts)
	if err != nil {
		return machines, err
	}
	for _, c := range ctrs {
		name, _, _ := getInfoFromLabels(c.Labels)
		m, err := pd.Machine(name, namespace)
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
	filters := getFilters("", "", true)
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.conn, opts)
	if err != nil {
		return namespaces, err
	}
	for _, c := range ctrs {
		_, ns, _ := getInfoFromLabels(c.Labels)
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
