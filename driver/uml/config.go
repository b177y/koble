package uml

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DefaultImage string `mapstructure:"default_image"`
	Kernel       string `mapstructure:"kernel"`
	RunDir       string `mapstructure:"run_dir" validate:"dir,max=24"`
	StorageDir   string `mapstructure:"storage_dir" validate:"dir"`
	Testing      bool   `mapstructure:"testing"`
}

func (ud *UMLDriver) loadConfig(conf map[string]interface{}) error {
	var err error
	vpl := viper.New()
	home := os.Getenv("UML_ORIG_HOME")
	if home == "" {
		home, err = os.UserHomeDir()
		if err != nil {
			return err
		}
	}
	uid, err := strconv.Atoi(os.Getenv("UML_ORIG_UID"))
	if err != nil {
		uid = os.Getuid()
	}
	vpl.SetDefault("default_image", "koble-fs")
	vpl.SetDefault("kernel", "koble-kernel")
	vpl.SetDefault("storage_dir", fmt.Sprintf("%s/.local/share/uml", home))
	vpl.SetDefault("run_dir", fmt.Sprintf("/run/user/%d/uml", uid))
	vpl.SetDefault("testing", false)
	err = vpl.MergeConfigMap(conf)
	if err != nil {
		return err
	}
	err = vpl.Unmarshal(&ud.Config)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(ud.Config.StorageDir, "overlay"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on overlay dir")
	}
	err = os.MkdirAll(filepath.Join(ud.Config.StorageDir, "images"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on imagesdir")
	}
	err = os.MkdirAll(filepath.Join(ud.Config.StorageDir, "kernel"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on imagesdir")
	}
	err = os.MkdirAll(filepath.Join(ud.Config.RunDir, "ns", "GLOBAL"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on ud.RunDir")
	}
	if ud.Config.Testing {
		imagesDir := filepath.Join(ud.Config.StorageDir, "images")
		kernelDir := filepath.Join(ud.Config.StorageDir, "kernel")
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		plzFs := filepath.Join(wd, "driver/uml/koble-fs/koble-fs")
		err = os.Symlink(plzFs, filepath.Join(imagesDir, "koble-fs"))
		if err != nil && !errors.Is(err, os.ErrExist) {
			return err
		}
		plzKern := filepath.Join(wd, "driver/uml/koble-kernel/linux")
		err = os.Symlink(plzKern, filepath.Join(kernelDir, "koble-kernel"))
		if err != nil && !errors.Is(err, os.ErrExist) {
			return err
		}
	}
	return validator.New().Struct(ud.Config)
}
