package uml

import (
	"errors"
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
	if !ud.Config.Testing && os.Getuid() != 0 {
		err = vecnet.CreateAndEnterUserNS("koble")
		if err != nil {
			return fmt.Errorf("Cannot create / enter user ns (%v): %w", os.Environ(), err)
		}
	} else if os.Getuid() != 0 {
		return errors.New("Testing needs to be run within a new user/mount namespace: `unshare -mUr go test ...`")
	}
	ud.Name = "UserMode Linux"
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
