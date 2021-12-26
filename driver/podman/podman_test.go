package podman_test

import (
	"testing"

	"github.com/b177y/netkit/driver/podman"
	driver_test "github.com/b177y/netkit/driver/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var pd = new(podman.PodmanDriver)

var _ = driver_test.DeclareAllDriverTests(pd)

func TestPodman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Podman Driver Suite")
}
