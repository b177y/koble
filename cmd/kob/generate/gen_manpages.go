package generate

import (
	"os"

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
	err = doc.GenManTree(kob.rootCmd, header, "out")
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenMarkdownTree(cmd.KobleCLI, "out")
	if err != nil {
		log.Fatal(err)
	}
}
