package cli

import (
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var Commands []*cobra.Command

var NK *koble.Koble

var Plain bool
var NoColor bool
