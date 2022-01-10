package koble

import (
	"crypto/md5"
	"fmt"

	"github.com/b177y/koble/driver"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Koble struct {
	Lab       Lab
	Config    Config
	Namespace string
	Driver    driver.Driver
}

type Lab struct {
	nk           *Koble
	Name         string                          `yaml:"name,omitempty" validate:"alphanum,max=30"`
	Directory    string                          `yaml:"dir,omitempty"`
	CreatedAt    string                          `yaml:"created_at,omitempty" validate:"datetime"`
	KobleVersion string                          `yaml:"koble_version,omitempty"`
	Description  string                          `yaml:"description,omitempty"`
	Authors      []string                        `yaml:"authors,omitempty"`
	Emails       []string                        `yaml:"emails,omitempty" validate:"email"`
	Web          []string                        `yaml:"web,omitempty" validate:"url"`
	Machines     map[string]driver.MachineConfig `yaml:"machines,omitempty"`
	Networks     map[string]driver.NetConfig     `yaml:"networks,omitempty"`
	DefaultImage string                          `yaml:"default_image,omitempty"`
}

func NewKoble(namespace string) (*Koble, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/koble")
	viper.SetDefault("driver", DriverConfig{Name: "podman"})
	viper.SetDefault("terminal", "gnome")
	viper.SetDefault("launch_terms", true)
	viper.SetDefault("launch_shell", false)
	viper.SetDefault("noninteractive", false)
	viper.SetDefault("nocolor", false)
	viper.SetDefault("default_namespace", "GLOBAL")
	viper.SetDefault("machine_memory", 128)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	lab := Lab{Name: ""}
	labExists, err := GetLab(&lab)
	if err != nil {
		return nil, err
	}
	var d driver.Driver
	if initialiser, ok := AvailableDrivers[config.Driver.Name]; ok {
		d = initialiser()
		err = d.SetupDriver(config.Driver.ExtraConf)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Driver %s is not currently supported.", config.Driver.Name)
	}
	nk := &Koble{
		Lab:    lab,
		Driver: d,
		Config: config,
	}
	if namespace != "" {
		nk.Namespace = namespace
	} else if labExists {
		nk.Namespace = fmt.Sprintf("%x",
			md5.Sum([]byte(lab.Directory)))
	} else {
		nk.Namespace = config.DefaultNamespace
	}
	err = validator.New().Var(nk.Namespace, "alphanum,max=32")
	if err != nil {
		return nil, err
	}
	return nk, nil
}
