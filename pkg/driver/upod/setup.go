package upod

import "github.com/b177y/koble/pkg/driver/podman"

func (ud *UMLDriver) SetupDriver(conf map[string]interface{}) (err error) {
	ud.Name = "User Mode Linux"
	err = ud.loadConfig(conf)
	if err != nil {
		return err
	}
	ud.Podman = podman.PodmanDriver{}
	err = ud.Podman.SetupDriver(conf)
	if err != nil {
		return err
	}
	ud.Podman.DriverName = "uml"
	return
}
