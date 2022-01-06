package podman

import (
	"strings"

	"github.com/b177y/netkit/driver"
	"github.com/containers/podman/v3/pkg/bindings/containers"
)

func (pd *PodmanDriver) ListMachines(namespace string, all bool) ([]driver.MachineInfo, error) {
	var machines []driver.MachineInfo
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", "", namespace, all) // TODO get namespace here
	opts.WithFilters(filters)
	ctrs, err := containers.List(pd.conn, opts)
	if err != nil {
		return machines, err
	}
	for _, c := range ctrs {
		name, _, _ := getInfoFromLabels(c.Labels)
		var mNetworks []string
		for _, n := range c.Networks {
			s := strings.Index(n, "koble_")
			if s == -1 {
				mNetworks = append(mNetworks, n)
			} else {
				s = s + 7 // s should be 0, + 7 accounts for koble_
				e := strings.Index(n[s:], "_")
				if e == -1 {
					mNetworks = append(mNetworks, n)
				} else {
					mNetworks = append(mNetworks, n[s:s+e])
				}
			}
		}
		machines = append(machines, driver.MachineInfo{
			Name:     name,
			Image:    c.Image[:strings.IndexByte(c.Image, ':')],
			State:    c.State,
			Uptime:   c.Status,
			ExitCode: c.ExitCode,
			ExitedAt: c.ExitedAt,
			Mounts:   c.Mounts,
			Pid:      c.Pid,
			Ports:    c.Ports,
		})
	}
	return machines, nil
}
func (pd *PodmanDriver) ListAllNamespaces() (namespaces []string, err error) {
	opts := new(containers.ListOptions)
	opts.WithAll(true)
	filters := getFilters("", "", "", true)
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
