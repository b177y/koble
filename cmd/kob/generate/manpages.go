package main

import (
	"os"

	"github.com/b177y/koble/cmd/kob/cli"
	_ "github.com/b177y/koble/cmd/kob/labs"
	_ "github.com/b177y/koble/cmd/kob/machines"
	_ "github.com/b177y/koble/cmd/kob/networks"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
)

func main() {
	header := &doc.GenManHeader{
		Title:   "Koble",
		Section: "1",
	}
	err := os.MkdirAll("out", 0700)
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenManTree(cli.RootCmd, header, "out")
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenMarkdownTree(cli.RootCmd, "out")
	if err != nil {
		log.Fatal(err)
	}
}
