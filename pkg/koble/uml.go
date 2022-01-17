//go:build !exclude_uml
// +build !exclude_uml

package koble

import (
	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/pkg/driver/uml"
)

func init() {
	registerDriver("uml", func() driver.Driver {
		return new(uml.UMLDriver)
	})
}
