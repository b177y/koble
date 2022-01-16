package cli

import (
	"github.com/b177y/koble/pkg/koble"
	"github.com/spf13/cobra"
)

var Commands []*cobra.Command

var NK *koble.Koble

func AddTermFlags(cmd *cobra.Command, launchOpt string) {
	cmd.Flags().StringP("terminal", "t", "gnome", "terminal to launch")
	koble.BindFlag("terminal."+launchOpt, cmd.Flags().Lookup("terminal"))
	cmd.Flags().StringToString("term-opt", map[string]string{}, "option to pass to terminal")
	koble.BindFlag("term_opts", cmd.Flags().Lookup("term-opt"))
}

func AddWaitFlag(cmd *cobra.Command) {
	cmd.Flags().Int("wait", 300, "seconds to wait for machine to boot before timeout, negative value will disable wait")
	koble.BindFlag("wait", cmd.Flags().Lookup("wait"))
}
