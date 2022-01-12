package uml

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/b177y/koble/driver/uml/vecnet"
)

func (ud *UMLDriver) SetupDriver(conf map[string]interface{}) (err error) {
	err = ud.loadConfig(conf)
	if err != nil {
		return err
	}
	err = vecnet.CreateAndEnterUserNS("koble")
	if err != nil {
		return fmt.Errorf("Cannot create / enter user ns: %w", err)
	}
	ud.Name = "UserMode Linux"
	err = os.MkdirAll(filepath.Join(ud.Config.StorageDir, "overlay"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on overlay dir")
	}
	err = os.MkdirAll(filepath.Join(ud.Config.StorageDir, "overlay"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on ud.StorageDir")
	}
	err = os.MkdirAll(filepath.Join(ud.Config.RunDir, "ns", "GLOBAL"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on ud.RunDir")
	}
	return nil
}
