package uml

import (
	"fmt"
	"os"

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
	vpl := viper.New()
	vpl.SetDefault("default_image", "/home/billy/repos/koble-fs/build/koble-fs")
	vpl.SetDefault("kernel", fmt.Sprintf("%s/netkit-jh/kernel/netkit-kernel", os.Getenv("UML_ORIG_HOME")))
	vpl.SetDefault("run_dir", fmt.Sprintf("/run/user/%s/uml", os.Getenv("UML_ORIG_UID")))
	vpl.SetDefault("storage_dir", fmt.Sprintf("%s/.local/share/uml", os.Getenv("UML_ORIG_HOME")))
	vpl.SetDefault("testing", false)
	err := vpl.MergeConfigMap(conf)
	if err != nil {
		return err
	}
	err = vpl.Unmarshal(&pd.Config)
	if err != nil {
		return err
	}
	validate := validator.New()
	return validate.Struct(pd.Config)
}
