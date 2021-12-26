package uml_test

import (
	"testing"

	driver_test "github.com/b177y/netkit/driver/tests"
	"github.com/b177y/netkit/driver/uml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ud = new(uml.UMLDriver)

var _ = driver_test.DeclareAllDriverTests(ud)

func TestPodman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UML Driver Suite")
}
