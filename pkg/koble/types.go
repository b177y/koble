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
	Name         string                          `koanf:"-" validate:"alphanum,max=30"`
	Directory    string                          `koanf:"-"`
	CreatedAt    string                          `koanf:"created_at,omitempty" validate:"datetime=02-01-2006"`
	KobleVersion string                          `koanf:"koble_version,omitempty"`
	Description  string                          `koanf:"description,omitempty"`
	Authors      []string                        `koanf:"authors,omitempty"`
	Emails       []string                        `koanf:"emails,omitempty" validate:"dive,email"`
	Web          []string                        `koanf:"web,omitempty" validate:"dive,url"`
	Machines     map[string]driver.MachineConfig `koanf:"machines,omitempty"`
	Networks     map[string]driver.NetConfig     `koanf:"networks,omitempty"`
}
