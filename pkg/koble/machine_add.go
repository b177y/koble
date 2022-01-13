package koble

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b177y/koble/driver"
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	log "github.com/sirupsen/logrus"
)

func (nk *Koble) AddMachineToLab(name string, conf driver.MachineConfig) error {
	if nk.LabRoot == "" {
		return errors.New("lab.yml does not exist, are you in a lab directory?")
	}
	log.WithFields(log.Fields{"name": name, "config": fmt.Sprintf("%+v", conf)}).
		Info("adding machine to lab")
	err := validator.New().Var(name, "alphanum,max=30")
	if err != nil {
		return err
	}

	if _, ok := nk.Lab.Machines[name]; ok {
		return fmt.Errorf("a machine named %s already exists", name)
	}

	err = os.Mkdir(name, 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	fn := name + ".startup"
	err = os.WriteFile(fn, []byte(DEFAULT_STARTUP), 0644)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	vpl := koanf.New(".")
	labConfPath := filepath.Join(nk.LabRoot, "lab.yml")
	if err := vpl.Load(file.Provider(labConfPath), yaml.Parser()); err != nil {
		return fmt.Errorf("error reading lab.yml: %w", err)
	}
	machineMap := make(map[string]interface{}, 0)
	if len(conf.Networks) != 0 {
		machineMap["networks"] = conf.Networks
	}
	if conf.Image != "" {
		machineMap["image"] = conf.Image
	}
	vpl.Load(confmap.Provider(map[string]interface{}{
		"machines." + name: machineMap,
	}, "."), nil)
	labBytes, err := vpl.Marshal(yaml.Parser())
	if err != nil {
		return fmt.Errorf("could not convert lab config to yaml: %w", err)
	}
	err = os.WriteFile(labConfPath, labBytes, 0644)
	if err != nil {
		return fmt.Errorf("could not write lab.yml to file: %w", err)
	}
	fmt.Printf("Created new machine %s, with directory for machine files and %s.startup as the machine startup script.\n", name, name)
	return nil
}
