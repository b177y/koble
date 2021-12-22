package netkit

import (
	"crypto/sha256"
	"fmt"

	"github.com/b177y/netkit/driver"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Netkit struct {
	Lab       Lab
	Config    Config
	Namespace string
	Driver    driver.Driver
}

func NewNetkit(namespace string) (*Netkit, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/netkit")
	viper.AddConfigPath("./examples/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	lab := Lab{
		Name: "",
	}
	labExists, err := GetLab(&lab)
	if err != nil {
		return nil, err
	}
	var d driver.Driver
	if initialiser, ok := availableDrivers[config.Driver.Name]; ok {
		d = initialiser()
		err = d.SetupDriver(config.Driver.ExtraConf)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Driver %s is not currently supported.", config.Driver.Name)
	}
	nk := &Netkit{
		Lab:    lab,
		Driver: d,
		Config: config,
	}
	if namespace != "" {
		nk.Namespace = namespace
	} else if labExists {
		nk.Namespace = fmt.Sprintf("%x",
			sha256.Sum256([]byte(lab.Directory)))
	} else {
		nk.Namespace = "GLOBAL"
	}
	err = validator.New().Var(nk.Namespace, "alphanum,max=64")
	if err != nil {
		return nil, err
	}
	return nk, nil
}
