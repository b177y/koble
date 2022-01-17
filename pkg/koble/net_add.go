package koble

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b177y/koble/pkg/driver"
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	log "github.com/sirupsen/logrus"
)

// redo with new viper config
func (nk *Koble) AddNetworkToLab(name string, conf driver.NetConfig) error {
	if nk.LabRoot == "" {
		return errors.New("lab.yml does not exist, are you in a lab directory?")
	}
	log.WithFields(log.Fields{"name": name, "config": fmt.Sprintf("%+v", conf)}).
		Info("adding network to lab")
	err := validator.New().Var(name, "alphanum,max=30")
	if err != nil {
		return err
	}
	if _, ok := nk.Lab.Networks[name]; ok {
		return fmt.Errorf("a network named %s already exists", name)
	}
	vpl := koanf.New(".")
	labConfPath := filepath.Join(nk.LabRoot, "lab.yml")
	if err := vpl.Load(file.Provider(labConfPath), yaml.Parser()); err != nil {
		return fmt.Errorf("error reading lab.yml: %w", err)
	}
	netMap := make(map[string]interface{}, 0)
	netMap["external"] = conf.External
	if conf.Gateway != "" {
		netMap["gateway"] = conf.Gateway
	}
	if conf.IpRange != "" {
		netMap["iprange"] = conf.IpRange
	}
	if conf.Subnet != "" {
		netMap["subnet"] = conf.Subnet
	}
	if conf.IPv6 != "" {
		netMap["ipv6"] = conf.IPv6
	}
	vpl.Load(confmap.Provider(map[string]interface{}{
		"networks." + name: netMap,
	}, "."), nil)
	labBytes, err := vpl.Marshal(yaml.Parser())
	if err != nil {
		return fmt.Errorf("could not convert lab config to yaml: %w", err)
	}
	err = os.WriteFile(labConfPath, labBytes, 0644)
	if err != nil {
		return fmt.Errorf("could not write lab.yml to file: %w", err)
	}
	fmt.Printf("Created new network %s.\n", name)
	return nil
}
