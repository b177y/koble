package netkit

import (
	"github.com/b177y/netkit/driver"
	"github.com/b177y/netkit/driver/podman"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Netkit struct {
	Lab    Lab
	Config Config
	Driver driver.Driver
}

func NewNetkit() (*Netkit, error) {
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
	_, err = getLab(&lab)
	if err != nil {
		return nil, err
	}
	d := new(podman.PodmanDriver)
	err = d.SetupDriver()
	if err != nil {
		return nil, err
	}
	nk := &Netkit{
		Lab:    lab,
		Driver: d,
		Config: config,
	}
	return nk, nil
}
