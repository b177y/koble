package main

import (
	"fmt"

	"github.com/b177y/netkit/cmd"
)

func main() {
	err := cmd.KobleCLI.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
