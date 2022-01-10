package uml

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DefaultImage string `mapstructure:"default_image"`
	Kernel       string `mapstructure:"kernel" validate:"file"`
	RunDir       string `mapstructure:"run_dir" validate:"dir"`
	StorageDir   string `mapstructure:"storage_dir" validate:"dir"`
	Testing      bool   `mapstructure:"testing"`
}

func (pd *UMLDriver) loadConfig(conf map[string]interface{}) error {
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
	vpl.SetDefault("default_image", "/home/billy/repos/koble-fs/build/koble-fs")
	vpl.SetDefault("kernel", fmt.Sprintf("%s/netkit-jh/kernel/netkit-kernel", home))
	vpl.SetDefault("run_dir", fmt.Sprintf("/run/user/%d/uml", uid))
	vpl.SetDefault("storage_dir", fmt.Sprintf("%s/.local/share/uml", home))
	vpl.SetDefault("testing", false)
	err = vpl.MergeConfigMap(conf)
	if err != nil {
		return err
	}
	err = vpl.Unmarshal(&pd.Config)
	if err != nil {
		return err
	}
	return validator.New().Struct(pd.Config)
}
