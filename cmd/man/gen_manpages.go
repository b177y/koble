package main

import (
	"os"

	"github.com/b177y/netkit/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
)

const OUTPUTDIR = "./build/man"

func main() {
	header := &doc.GenManHeader{
		Title:   "Netkit",
		Section: "1",
	}
	err := os.MkdirAll(OUTPUTDIR, 0700)
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenManTree(cmd.NetkitCLI, header, OUTPUTDIR)
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenMarkdownTree(cmd.NetkitCLI, OUTPUTDIR)
	if err != nil {
		log.Fatal(err)
	}
}
