package cli

import (
	"github.com/spf13/cobra"
)

// Autocompletion function to get list of active namespaces
func AutocompNamespace(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	namespaces, err := NK.Driver.ListAllNamespaces()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	return namespaces, cobra.ShellCompDirectiveNoFileComp
}

// Autocompletion function to get list of machines
func AutocompMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := NK.Driver.ListMachines(NK.Config.Namespace, true)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	var machines []string
	for _, m := range machineList {
		machines = append(machines, m.Name)
	}
	return machines, cobra.ShellCompDirectiveNoFileComp
}

// Autocompletion function to get list of running machines
func AutocompRunningMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := NK.Driver.ListMachines(NK.Config.Namespace, true)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	var machines []string
	for _, m := range machineList {
		if m.State == "running" {
			machines = append(machines, m.Name)
		}
	}
	return machines, cobra.ShellCompDirectiveNoFileComp
}

// Autocompletion function to get list of non-running machines
func AutocompNonRunningMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := NK.Driver.ListMachines(NK.Config.Namespace, true)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	var machines []string
	for _, m := range machineList {
		if m.State != "running" {
			machines = append(machines, m.Name)
		}
	}
	return machines, cobra.ShellCompDirectiveNoFileComp
}
