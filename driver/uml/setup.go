package uml

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b177y/koble/driver/uml/vecnet"
)

func (ud *UMLDriver) SetupDriver(conf map[string]interface{}) (err error) {
	if val, ok := conf["testing"]; ok {
		if b, ok := val.(bool); ok {
			ud.Testing = b
		} else {
			return fmt.Errorf("Driver 'testing' in config must be a bool.")
		}
	}
	if !ud.Testing && os.Getuid() != 0 {
		err = vecnet.CreateAndEnterUserNS("koble")
		if err != nil {
			return fmt.Errorf("Cannot create / enter user ns (%v): %w", os.Environ(), err)
		}
	} else if os.Getuid() != 0 {
		return errors.New("Testing needs to be run within a new user/mount namespace: `unshare -mUr go test ...`")
	}
	ud.Name = "UserMode Linux"
	// ud.Kernel = "/home/billy/repos/netkit-jh-build/tmpbuild/linux-5.14.9/linux"
	ud.Kernel = fmt.Sprintf("%s/netkit-jh/kernel/netkit-kernel", os.Getenv("UML_ORIG_HOME"))
	// ud.DefaultImage = fmt.Sprintf("%s/netkit-jh/fs/custom-fs", homedir)
	ud.DefaultImage = "/home/billy/repos/koble-fs/build/koble-fs"
	ud.RunDir = fmt.Sprintf("/run/user/%s/uml", os.Getenv("UML_ORIG_UID"))
	ud.StorageDir = fmt.Sprintf("%s/.local/share/uml", os.Getenv("UML_ORIG_HOME"))
	// override kernel with config option
	if val, ok := conf["kernel"]; ok {
		if str, ok := val.(string); ok {
			ud.Kernel = str
		} else {
			return fmt.Errorf("Driver 'kernel' in config must be a string.")
		}
	}
	if val, ok := conf["default_image"]; ok {
		if str, ok := val.(string); ok {
			ud.DefaultImage = str
		} else {
			return fmt.Errorf("Driver 'default_image' in config must be a string.")
		}
	}
	if val, ok := conf["run_dir"]; ok {
		if str, ok := val.(string); ok {
			ud.RunDir = str
		} else {
			return fmt.Errorf("Driver 'run_dir' in config must be a string.")
		}
	}
	if val, ok := conf["storage_dir"]; ok {
		if str, ok := val.(string); ok {
			ud.StorageDir = str
		} else {
			return fmt.Errorf("Driver 'storage_dir' in config must be a string.")
		}
	}
	err = os.MkdirAll(filepath.Join(ud.StorageDir, "overlay"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on ud.StorageDir")
	}
	err = os.MkdirAll(filepath.Join(ud.RunDir, "ns", "GLOBAL"), 0744)
	if err != nil && err != os.ErrExist {
		return fmt.Errorf("Could not mkdir on ud.RunDir")
	}
	return nil
}
