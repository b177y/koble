package cli

import (
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var Commands []*cobra.Command

var NK *koble.Koble

func AddTermFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("terminal", "t", "gnome", "terminal to launch")
	koble.BindFlag("terminal.name", cmd.Flags().Lookup("terminal"))
	cmd.Flags().Bool("launch", false, "launch terminal for attach session")
	koble.BindFlag("terminal.launch", cmd.Flags().Lookup("launch"))
	cmd.Flags().StringToString("term-opt", map[string]string{}, "option to pass to terminal")
	koble.BindFlag("term_opts", cmd.Flags().Lookup("term-opt"))
}

func AddWaitFlag(cmd *cobra.Command) {
	cmd.Flags().Int("wait", 300, "seconds to wait for machine to boot before timeout (default 300, -1 is don't wait)")
	koble.BindFlag("wait", cmd.Flags().Lookup("wait"))
}
