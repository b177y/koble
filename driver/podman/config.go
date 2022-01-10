package podman

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DefaultImage string `mapstructure:"default_image"`
	URI          string `mapstructure:"uri" validate:"uri"`
}

func (pd *PodmanDriver) loadConfig(conf map[string]interface{}) error {
	vpl := viper.New()
	vpl.SetDefault("URI", fmt.Sprintf("unix://run/user/%d/podman/podman.sock",
		os.Getuid()))
	vpl.SetDefault("default_image", "localhost/koble-deb-test")
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
