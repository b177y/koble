//go:build !exclude_podman
// +build !exclude_podman

package netkit

import (
	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/podman"
)

func init() {
	registerDriver("podman", func() driver.Driver {
		return new(podman.PodmanDriver)
	})
}
