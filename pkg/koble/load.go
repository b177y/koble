package koble

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/b177y/koble/driver"
	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func Load(namespace string) (*Koble, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/koble")
	viper.SetDefault("driver", DriverConfig{Default: "podman"})
	viper.SetDefault("terminal", "gnome")
	viper.SetDefault("launch_terms", true)
	viper.SetDefault("launch_shell", false)
	viper.SetDefault("noninteractive", false)
	viper.SetDefault("nocolor", false)
	viper.SetDefault("default_namespace", "GLOBAL")
	viper.SetDefault("machine_memory", 128)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	var d driver.Driver
	if initialiser, ok := AvailableDrivers[config.Driver.Default]; ok {
		d = initialiser()
		err = d.SetupDriver(config.Driver.ExtraConf)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Driver %s is not currently supported.", config.Driver.Default)
	}
	nk := &Koble{
		Driver: d,
		Config: config,
	}
	nk.LabRoot, err = getLabRoot()
	if err != nil {
		return nil, err
	}
	if namespace != "" {
		nk.Namespace = namespace
	} else if nk.LabRoot != "" {
		nk.Namespace = fmt.Sprintf("%x",
			md5.Sum([]byte(nk.LabRoot)))
	} else {
		nk.Namespace = config.DefaultNamespace
	}
	err = validator.New().Var(nk.Namespace, "alphanum,max=32")
	if err != nil {
		return nil, err
	}
	color.NoColor = config.NoColor
	return nk, nil
}

func (nk *Koble) LoadLab() (err error) {
	if nk.LabRoot == "" {
		return nil
	}
	// change dir to labroot
	nk.InitialWorkDir, err = os.Getwd()
	if err != nil {
		return err
	}
	if nk.LabRoot != nk.InitialWorkDir {
		err = os.Chdir(nk.LabRoot)
		if err != nil {
			return err
		}
	}
	vpl := viper.New()
	vpl.SetConfigName("lab")
	vpl.SetConfigType("yaml")
	vpl.AddConfigPath(".")

	err = vpl.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read lab.yml: %w", err)
	}
	err = vpl.Unmarshal(&nk.Lab)
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}
	nk.Lab.Machines, err = orderMachines(nk.Lab.Machines)
	if err != nil {
		return fmt.Errorf("could not order lab machines by dependency: %w", err)
	}
	nk.Lab.Name = filepath.Base(nk.LabRoot)
	nk.Lab.Directory = nk.LabRoot
	return nil
}

// Return to initial working directory from labroot
func (nk *Koble) ExitLab() (err error) {
	if nk.LabRoot == "" {
		return nil
	}
	if nk.LabRoot != nk.InitialWorkDir {
		return os.Chdir(nk.InitialWorkDir)
	}
	return nil
}
