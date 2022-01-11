package uml_test

import (
	"testing"

	driver_test "github.com/b177y/koble/driver/tests"
	"github.com/b177y/koble/driver/uml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var ud = new(uml.UMLDriver)

func TestUML(t *testing.T) {
	// err := vecnet.CreateAndEnterUserNS("koble")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	err := driver_test.DeclareAllDriverTests(ud)
	if err != nil {
		t.Fatal(err)
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "UML Driver Suite")
}
