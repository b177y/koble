package koble

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func (nk *Koble) LoadLab() (err error) {
	if nk.LabRoot == "" {
		return nil
	}
	vpl := viper.New()
	vpl.SetConfigName("lab")
	vpl.SetConfigType("yaml")
	vpl.AddConfigPath(nk.LabRoot)

	err = vpl.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read lab.yml: %w", err)
	}

	// if lab does not set namespace, set it to lab path hash
	if vpl.Get("namespace") == nil {
		vpl.Set("namespace", fmt.Sprintf("%x", md5.Sum([]byte(nk.LabRoot))))
	}
	err = vpl.Unmarshal(&nk.Lab)
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	nk.Lab.Name = filepath.Base(nk.LabRoot)
	nk.Lab.Directory = nk.LabRoot

	err = validator.New().Struct(nk.Lab)
	if err != nil {
		return fmt.Errorf("error validating lab.yml: %w", err)
	}

	nk.Lab.Machines, err = orderMachines(nk.Lab.Machines)
	if err != nil {
		return fmt.Errorf("could not order lab machines by dependency: %w", err)
	}

	// cm := make(map[string]interface{}, 0)
	// cm["driver"] = vpl.Get("driver")
	// viper.AllSettings()
	err = viper.MergeConfigMap(vpl.AllSettings())
	if err != nil {
		return fmt.Errorf("Could not merge lab driver config to default driver config")
	}
	return nil
}

func Load() (*Koble, error) {
	var nk Koble
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// cursed uml support
	if os.Getenv("UML_ORIG_HOME") != "" {
		viper.AddConfigPath("$UML_ORIG_HOME/.config/koble")
	} else {
		viper.AddConfigPath("$HOME/.config/koble")
	}
	viper.SetDefault("driver.name", "podman")
	viper.SetDefault("terminal.name", "gnome")
	viper.SetDefault("launch_terms", true)
	viper.SetDefault("launch_shell", false)
	viper.SetDefault("noninteractive", false)
	viper.SetDefault("nocolor", false)
	viper.SetDefault("namespace", "GLOBAL")
	viper.SetDefault("machine.memory", 128)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("error reading config.yml: %w", err)
	}
	nk.LabRoot, err = getLabRoot()
	if err != nil {
		return nil, err
	}
	// set up lab if in one
	if nk.LabRoot != "" {
		err = nk.LoadLab()
		if err != nil {
			return nil, err
		}
	}
	err = viper.Unmarshal(&nk.Config)
	if err != nil {
		return nil, fmt.Errorf("error loading config.yml: %w", err)
	}
	if initialiser, ok := AvailableDrivers[nk.Config.Driver.Name]; ok {
		nk.Driver = initialiser()
		err = nk.Driver.SetupDriver(nk.Config.Driver.ExtraConf)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Driver %s is not currently supported.", nk.Config.Driver.Name)
	}
	err = validator.New().Struct(nk.Config)
	if err != nil {
		return nil, fmt.Errorf("error validating config.yml: %w", err)
	}
	nk.InitialWorkDir, err = os.Getwd()
	if err != nil {
		return nil, err
	}
	if nk.LabRoot != "" && nk.LabRoot != nk.InitialWorkDir {
		err = os.Chdir(nk.LabRoot)
		if err != nil {
			return nil, err
		}
	}
	color.NoColor = nk.Config.NoColor
	if nk.Config.Verbosity != 0 && nk.Config.Quiet {
		return nil, fmt.Errorf("verbose and quiet options cannot be used together")
	}
	switch nk.Config.Verbosity {
	case 0:
		log.SetLevel(log.WarnLevel)
	case 1:
		log.SetLevel(log.InfoLevel)
		nk.Config.NonInteractive = true
	case 2:
		log.SetLevel(log.DebugLevel)
		nk.Config.NonInteractive = true
	case 3:
		log.SetLevel(log.TraceLevel)
		nk.Config.NonInteractive = true
	default:
		return nil, fmt.Errorf("verbosity level %d is not valid",
			nk.Config.Verbosity)
	}
	if nk.Config.Quiet {
		log.SetLevel(log.ErrorLevel)
	}

	return &nk, nil
}

// Return to initial working directory from labroot
func (nk *Koble) Cleanup() (err error) {
	if nk.LabRoot == "" {
		return nil
	}
	if nk.LabRoot != nk.InitialWorkDir {
		return os.Chdir(nk.InitialWorkDir)
	}
	return nil
}
