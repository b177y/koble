package koble

import (
	"errors"
	"fmt"
	"os"

	"github.com/b177y/koble/driver"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// redo with new viper config
func (nk *Koble) AddMachineToLab(name string, conf driver.MachineConfig) error {
	if nk.LabRoot == "" {
		return errors.New("lab.yml does not exist, are you in a lab directory?")
	}
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

	vpl := viper.New()
	vpl.SetConfigName("lab")
	vpl.SetConfigType("yaml")
	vpl.AddConfigPath(nk.LabRoot)
	err = vpl.ReadInConfig()
	if err != nil {
		return fmt.Errorf("could not read lab.yml: %w", err)
	}
	// write networks even if empty, so that machine gets added to lab.yml
	vpl.Set("machines."+name+".networks", conf.Networks)
	if conf.Image != "" {
		vpl.Set("machines."+name+".image", conf.Image)
	}
	err = vpl.WriteConfigAs("lab.yml")
	if err != nil {
		return err
	}
	fmt.Printf("Created new machine %s, with directory for machine files and %s.startup as the machine startup script.\n", name, name)
	return nil
}
