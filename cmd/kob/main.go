package kob

import (
	"fmt"
)

func main() {
	err := rootCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
