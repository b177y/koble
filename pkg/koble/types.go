package koble

import (
	"github.com/b177y/koble/driver"
	"github.com/knadh/koanf"
)

type Koble struct {
	Lab            Lab
	LabRoot        string
	InitialWorkDir string
	Config         Config
	Driver         driver.Driver
	Koanf          koanf.Koanf
}

type Lab struct {
	Name         string                          `mapstructure:"-" validate:"alphanum,max=30"`
	Directory    string                          `mapstructure:"-"`
	CreatedAt    string                          `mapstructure:"created_at,omitempty" validate:"datetime=02-01-2006"`
	KobleVersion string                          `mapstructure:"koble_version,omitempty"`
	Description  string                          `mapstructure:"description,omitempty"`
	Authors      []string                        `mapstructure:"authors,omitempty"`
	Emails       []string                        `mapstructure:"emails,omitempty" validate:"dive,email"`
	Web          []string                        `mapstructure:"web,omitempty" validate:"dive,url"`
	Machines     map[string]driver.MachineConfig `mapstructure:"machines,omitempty"`
	Networks     map[string]driver.NetConfig     `mapstructure:"networks,omitempty"`
}
