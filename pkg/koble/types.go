package koble

import (
	"github.com/b177y/koble/driver"
)

type Koble struct {
	Lab            Lab
	LabRoot        string
	InitialWorkDir string
	Config         Config
	Driver         driver.Driver
}

type Lab struct {
	Name         string                          `mapstructure:"name,omitempty" validate:"alphanum,max=30"`
	Directory    string                          `mapstructure:"-"`
	CreatedAt    string                          `mapstructure:"created_at,omitempty" validate:"datetime"`
	KobleVersion string                          `mapstructure:"koble_version,omitempty"`
	Description  string                          `mapstructure:"description,omitempty"`
	Authors      []string                        `mapstructure:"authors,omitempty"`
	Emails       []string                        `mapstructure:"emails,omitempty" validate:"email"`
	Web          []string                        `mapstructure:"web,omitempty" validate:"url"`
	Machines     map[string]driver.MachineConfig `mapstructure:"machines,omitempty"`
	Networks     map[string]driver.NetConfig     `mapstructure:"networks,omitempty"`
}
