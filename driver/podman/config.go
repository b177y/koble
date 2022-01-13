package podman

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/confmap"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DefaultImage string `koanf:"default_image"`
	URI          string `koanf:"uri" validate:"uri"`
}

func (pd *PodmanDriver) loadConfig(conf map[string]interface{}) error {
	vpl := koanf.New(".")
	err := vpl.Load(confmap.Provider(map[string]interface{}{
		"podman.uri":           fmt.Sprintf("unix://run/user/%d/podman/podman.sock", os.Getuid()),
		"podman.default_image": "localhost/koble-deb-test",
	}, ""), nil)
	if err != nil {
		return err
	}
	err = vpl.Load(confmap.Provider(conf, ""), nil)
	if err != nil {
		return err
	}
	err = vpl.Unmarshal("podman", &pd.Config)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"driver": "podman",
		"config": fmt.Sprintf("%+v", pd.Config)}).Debug("loaded driver config")
	return validator.New().Struct(pd.Config)
}
