package koble

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type InitOpts struct {
	Name        string
	Description string
	Authors     []string
	Emails      []string
	Webs        []string
}

func (nk *Koble) InitLab(options InitOpts) error {
	if nk.LabRoot == nk.InitialWorkDir {
		fmt.Println("nk labroot", nk.LabRoot, nk.InitialWorkDir)
		return fmt.Errorf("lab.yml already exists in this directory.")
	} else if nk.LabRoot != "" {
		log.Warnf("There is already a lab at %s, creating a new lab in %s\n", nk.LabRoot, nk.InitialWorkDir)
		err := os.Chdir(nk.InitialWorkDir)
		if err != nil {
			return err
		}
	}
	if options.Name == "" {
		log.Debug("Name not given, initialising lab in current directory.")
		options.Name = filepath.Base(nk.InitialWorkDir)
		err := validator.New().Var(options.Name, "alphanum,max=30")
		if err != nil {
			return err
		}
	} else {
		err := validator.New().Var(options.Name, "alphanum,max=30")
		if err != nil {
			return err
		}
		if fileExists(options.Name) {
			return fmt.Errorf("file or directory %s already exists", options.Name)
		}
		err = os.Mkdir(options.Name, 0755)
		if err != nil {
			return err
		}
		err = os.Chdir(options.Name)
		if err != nil {
			return err
		}
	}
	// TODO check if in script mode
	// ask for name, description etc
	vpl := viper.New()
	vpl.Set("koble_version", VERSION)
	vpl.Set("created_at", time.Now().Format("02-01-2006"))
	if options.Description != "" {
		vpl.Set("description", options.Description)
	}
	if len(options.Authors) != 0 {
		vpl.Set("authors", options.Authors)
	}
	if len(options.Emails) != 0 {
		vpl.Set("emails", options.Emails)
	}
	if len(options.Webs) != 0 {
		vpl.Set("webs", options.Webs)
	}
	err := vpl.SafeWriteConfigAs("lab.yml")
	if err != nil {
		return err
	}
	err = os.Mkdir("shared", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	err = os.WriteFile("shared.startup", []byte(SHARED_STARTUP), 0644)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	return nil
}
