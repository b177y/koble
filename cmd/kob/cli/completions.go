package cli

import (
	"github.com/spf13/cobra"
)

func AutocompNamespace(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	namespaces, err := NK.Driver.ListAllNamespaces()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	return namespaces, cobra.ShellCompDirectiveNoFileComp
}

func AutocompMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := NK.Driver.ListMachines(NK.Namespace, true)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	var machines []string
	for _, m := range machineList {
		machines = append(machines, m.Name)
	}
	return machines, cobra.ShellCompDirectiveNoFileComp
}

func AutocompRunningMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := NK.Driver.ListMachines(NK.Namespace, true)
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

func AutocompNonRunningMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := NK.Driver.ListMachines(NK.Namespace, true)
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
