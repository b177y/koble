package koble

import (
	"github.com/b177y/koble/pkg/driver"
)

type Koble struct {
	Lab            Lab
	LabRoot        string
	InitialWorkDir string
	Config         Config
	Driver         driver.Driver
}

type Lab struct {
	Name         string                          `koanf:"-" validate:"alphanum,max=30"`
	Directory    string                          `koanf:"-"`
	CreatedAt    string                          `koanf:"created_at" validate:"datetime=02-01-2006"`
	KobleVersion string                          `koanf:"koble_version"`
	Description  string                          `koanf:"description"`
	Authors      []string                        `koanf:"authors"`
	Emails       []string                        `koanf:"emails" validate:"dive,email"`
	Webs         []string                        `koanf:"web" validate:"dive,url"`
	Machines     map[string]driver.MachineConfig `koanf:"machines"`
	Networks     map[string]driver.NetConfig     `koanf:"networks" mapstructure:"networks"`
}
