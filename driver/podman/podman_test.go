package podman_test

import (
	"testing"

	"github.com/b177y/koble/driver/podman"

	driver_test "github.com/b177y/koble/driver/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var pd = new(podman.PodmanDriver)

func TestPodman(t *testing.T) {
	err := driver_test.DeclareAllDriverTests(pd)
	if err != nil {
		t.Fatal(err)
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Podman Driver Suite")
}
