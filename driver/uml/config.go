package uml

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/confmap"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DefaultImage string `koanf:"default_image"`
	Kernel       string `koanf:"kernel"`
	RunDir       string `koanf:"run_dir" validate:"max=24"`
	StorageDir   string `koanf:"storage_dir"`
	Testing      bool   `koanf:"testing"`
}

func (ud *UMLDriver) loadConfig(conf map[string]interface{}) error {
	var err error
	vpl := koanf.New(".")
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
	err = vpl.Load(confmap.Provider(map[string]interface{}{
		"uml.default_image": "koble-fs",
		"uml.kernel":        "koble-kernel",
		"uml.storage_dir":   fmt.Sprintf("%s/.local/share/uml", home),
		"uml.run_dir":       fmt.Sprintf("/run/user/%d/uml", uid),
		"uml.testing":       false,
	}, ""), nil)
	if err != nil {
		return err
	}
	err = vpl.Load(confmap.Provider(conf, ""), nil)
	if err != nil {
		return err
	}
	err = vpl.Unmarshal("uml", &ud.Config)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"driver": "uml",
		"config": fmt.Sprintf("%+v", ud.Config)}).Debug("loaded driver config")
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
