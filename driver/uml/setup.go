package uml

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b177y/netkit/driver/uml/vecnet"
)

func (ud *UMLDriver) SetupDriver(conf map[string]interface{}) (err error) {
	if val, ok := conf["testing"]; ok {
		if b, ok := val.(bool); ok {
			ud.Testing = b
		} else {
			return fmt.Errorf("Driver 'testing' in config must be a bool.")
		}
	}
	if !ud.Testing {
		err = vecnet.CreateAndEnterUserNS("netkit")
	} else if os.Getuid() != 0 {
		return errors.New("Testing needs to be run within a new user/mount namespace: `unshare -mUr go test ...`")
	}
	if err != nil {
		return fmt.Errorf("Cannot create / enter user ns: %w", err)
	}
	ud.Name = "UserMode Linux"
	homedir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Could not get user home dir: %w", err)
	}
	// ud.Kernel = "/home/billy/repos/netkit-jh-build/tmpbuild/linux-5.14.9/linux"
	ud.Kernel = fmt.Sprintf("%s/netkit-jh/kernel/netkit-kernel", homedir)
	// ud.DefaultImage = fmt.Sprintf("%s/netkit-jh/fs/custom-fs", homedir)
	ud.DefaultImage = "/home/billy/repos/netkit-fs/build/netkit-fs"
	ud.RunDir = fmt.Sprintf("/run/user/%s/uml", os.Getenv("UML_ORIG_UID"))
	ud.StorageDir = fmt.Sprintf("%s/.local/share/uml", homedir)
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
