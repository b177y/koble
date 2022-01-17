//go:build !exclude_podman
// +build !exclude_podman

package koble

import (
	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/driver/podman"
)

func init() {
	registerDriver("podman", func() driver.Driver {
		return new(podman.PodmanDriver)
	})
}
