package podman

import (
	"context"

	"github.com/containers/podman/v3/pkg/bindings"
	log "github.com/sirupsen/logrus"
)

func (pd *PodmanDriver) SetupDriver(conf map[string]interface{}) (err error) {
	pd.Name = "Podman"
	err = pd.loadConfig(conf)
	if err != nil {
		return err
	}
	log.Debug("Attempting to connect to podman socket.")
	pd.Conn, err = bindings.NewConnection(context.Background(), pd.Config.URI)
	if err != nil {
		return err
	}
	return nil
}
