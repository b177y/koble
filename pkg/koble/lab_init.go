package koble

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/b177y/koble/util/validator"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	log "github.com/sirupsen/logrus"
)

type InitOpts struct {
	Name        string
	Description string
	Authors     []string
	Emails      []string
	Webs        []string
}

func (nk *Koble) InitLab(options InitOpts) error {
	labPath := filepath.Join(nk.InitialWorkDir, options.Name)
	log.WithFields(log.Fields{"LabRoot": nk.LabRoot, "labPath": labPath,
		"Options": fmt.Sprintf("%+v", options)}).Info("initialising new lab directory")
	if fileExists(filepath.Join(labPath, "lab.yml")) {
		return fmt.Errorf("lab.yml already exists in this directory.")
	} else if nk.LabRoot != "" {
		log.Warnf("There is already a lab in parent directory %s, creating a new lab in %s\n", nk.LabRoot, labPath)
	}
	err := os.MkdirAll(labPath, 0755)
	if err != nil {
		return err
	}
	err = os.Chdir(labPath)
	if err != nil {
		return err
	}
	options.Name = filepath.Base(labPath)
	if !validator.IsValidName(options.Name) {
		return fmt.Errorf("lab name '%s' must be alphanumeric and no more than 32 chars", options.Name)
	}
	// TODO check if in script mode
	// ask for name, description etc
	vpl := koanf.New(".")
	labMap := make(map[string]interface{}, 0)
	labMap["koble_version"] = VERSION
	labMap["created_at"] = time.Now().Format("02-01-2006")
	if options.Description != "" {
		labMap["description"] = options.Description
	}
	if len(options.Authors) != 0 {
		labMap["authors"] = options.Authors
	}
	if len(options.Emails) != 0 {
		labMap["emails"] = options.Emails
	}
	if len(options.Webs) != 0 {
		labMap["web"] = options.Webs
	}
	vpl.Load(confmap.Provider(labMap, "."), nil)
	labBytes, err := vpl.Marshal(yaml.Parser())
	if err != nil {
		return fmt.Errorf("could not convert lab config to yaml: %w", err)
	}
	labConfPath := filepath.Join(labPath, "lab.yml")
	err = os.WriteFile(labConfPath, labBytes, 0644)
	if err != nil {
		return fmt.Errorf("could not write lab.yml to file: %w", err)
	}
	err = os.Mkdir("shared", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	if !fileExists("shared.startup") {
		err = os.WriteFile("shared.startup", []byte(SHARED_STARTUP), 0644)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Initialised lab %s\n", options.Name)
	return nil
}
