package uml_test

import (
	"testing"

	driver_test "github.com/b177y/netkit/driver/tests"
	"github.com/b177y/netkit/driver/uml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ud = new(uml.UMLDriver)

func TestUML(t *testing.T) {
	err := driver_test.DeclareAllDriverTests(ud)
	if err != nil {
		t.Fatal(err)
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "UML Driver Suite")
}
