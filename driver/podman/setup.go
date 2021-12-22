package podman

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v3/pkg/bindings"
	log "github.com/sirupsen/logrus"
)

type PodmanDriver struct {
	conn         context.Context
	Name         string
	DefaultImage string
	URI          string
}

func (pd *PodmanDriver) GetDefaultImage() string {
	return pd.DefaultImage
}

func (pd *PodmanDriver) SetupDriver(conf map[string]interface{}) (err error) {
	pd.Name = "Podman"
	pd.URI = fmt.Sprintf("unix://run/user/%s/podman/podman.sock",
		fmt.Sprint(os.Getuid()))
	pd.DefaultImage = "localhost/netkit-deb-test"
	// override uri with config option
	if val, ok := conf["uri"]; ok {
		if str, ok := val.(string); ok {
			pd.URI = str
		} else {
			return fmt.Errorf("Driver 'uri' in config must be a string.")
		}
	}
	if val, ok := conf["default_image"]; ok {
		if str, ok := val.(string); ok {
			pd.DefaultImage = str
		} else {
			return fmt.Errorf("Driver 'default_image' in config must be a string.")
		}
	}
	log.Debug("Attempting to connect to podman socket.")
	pd.conn, err = bindings.NewConnection(context.Background(), pd.URI)
	if err != nil {
		return err
	}
	return nil
}
