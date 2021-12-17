package cmd

import (
	"github.com/spf13/cobra"
)

func autocompNamespace(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	namespaces, err := nk.Driver.ListAllNamespaces()
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	return namespaces, cobra.ShellCompDirectiveNoFileComp
}

func autocompMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := nk.Driver.ListMachines(namespace, true)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveError
	}
	var machines []string
	for _, m := range machineList {
		machines = append(machines, m.Name)
	}
	return machines, cobra.ShellCompDirectiveNoFileComp
}

func autocompRunningMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := nk.Driver.ListMachines(namespace, true)
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

func autocompNonRunningMachine(cmd *cobra.Command, args []string,
	toComplete string) ([]string, cobra.ShellCompDirective) {
	machineList, err := nk.Driver.ListMachines(namespace, true)
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
