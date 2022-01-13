package koble

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
)

var (
	Koanf  = koanf.New(".")
	kFlags = []koanf.Provider{}
)

func BindFlag(confKey string, f *flag.Flag) {
	flagSet := flag.NewFlagSet(f.Name, flag.ContinueOnError)
	flagSet.AddFlag(f)
	kFlags = append(kFlags, posflag.ProviderWithFlag(flagSet, ".", Koanf,
		func(f *flag.Flag) (string, interface{}) {
			return confKey, posflag.FlagVal(flagSet, f)
		}))
}

func (nk *Koble) LoadLab() (err error) {
	if nk.LabRoot == "" {
		log.Debug("not in lab directory so not loading lab config")
		return nil
	}
	vpl := koanf.New(".")
	labConfPath := filepath.Join(nk.LabRoot, "lab.yml")
	if err := vpl.Load(file.Provider(labConfPath), yaml.Parser()); err != nil {
		return fmt.Errorf("error reading lab.yml: %w", err)
	}

	// if lab does not set namespace, set it to lab path hash
	// if vpl.String("namespace") == "" {
	// 	vpl.Set("namespace", fmt.Sprintf("%x", md5.Sum([]byte(nk.LabRoot))))
	// }
	err = vpl.Unmarshal("", &nk.Lab)
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	nk.Lab.Name = filepath.Base(nk.LabRoot)
	nk.Lab.Directory = nk.LabRoot

	err = validator.New().Struct(nk.Lab)
	if err != nil {
		return fmt.Errorf("error validating lab.yml: %w", err)
	}
	nk.Lab.Machines, err = orderMachines(nk.Lab.Machines)
	if err != nil {
		return fmt.Errorf("could not order lab machines by dependency: %w", err)
	}
	err = Koanf.Merge(vpl)
	if err != nil {
		return fmt.Errorf("Could not merge lab driver config to default driver config")
	}
	return nil
}

func Load() (*Koble, error) {
	var nk Koble
	Koanf.Load(confmap.Provider(defaultConfig, "."), nil)
	// cursed uml support
	var confPath string
	if home := os.Getenv("UML_ORIG_HOME"); home != "" {
		confPath = filepath.Join(home, "/.config/koble/config.yml")
	} else {
		home = os.Getenv("HOME")
		confPath = filepath.Join(home, "/.config/koble/config.yml")
	}
	if err := Koanf.Load(file.Provider(confPath), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("error reading config.yml: %w", err)
	}
	var err error
	nk.LabRoot, err = getLabRoot()
	if err != nil {
		return nil, err
	}
	// set up lab if in one
	if nk.LabRoot != "" {
		err = nk.LoadLab()
		if err != nil {
			return nil, err
		}
	}
	// process flags
	for _, pf := range kFlags {
		if err := Koanf.Load(pf, nil); err != nil {
			return nil, fmt.Errorf("loading config from flag: %w", err)
		}
	}
	err = Koanf.Unmarshal("", &nk.Config)
	if err != nil {
		return nil, fmt.Errorf("error loading config.yml: %w", err)
	}
	err = validator.New().Struct(nk.Config)
	if err != nil {
		return nil, fmt.Errorf("error validating config.yml: %w", err)
	}
	if initialiser, ok := AvailableDrivers[nk.Config.Driver.Name]; ok {
		nk.Driver = initialiser()
		err = nk.Driver.SetupDriver(nk.Config.Driver.ExtraConf)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Driver %s is not currently supported.", nk.Config.Driver.Name)
	}
	nk.InitialWorkDir, err = os.Getwd()
	if err != nil {
		return nil, err
	}
	if nk.LabRoot != "" && nk.LabRoot != nk.InitialWorkDir {
		err = os.Chdir(nk.LabRoot)
		if err != nil {
			return nil, err
		}
	}
	color.NoColor = nk.Config.NoColor
	if nk.Config.Verbosity != 0 && nk.Config.Quiet {
		return nil, fmt.Errorf("verbose and quiet options cannot be used together")
	}
	switch nk.Config.Verbosity {
	case 0:
		log.SetLevel(log.WarnLevel)
	case 1:
		log.SetLevel(log.InfoLevel)
		nk.Config.NonInteractive = true
	case 2:
		log.SetLevel(log.DebugLevel)
		nk.Config.NonInteractive = true
	case 3:
		log.SetLevel(log.TraceLevel)
		nk.Config.NonInteractive = true
	default:
		return nil, fmt.Errorf("verbosity level %d is not valid",
			nk.Config.Verbosity)
	}
	if nk.Config.Quiet {
		log.SetLevel(log.ErrorLevel)
	}

	return &nk, nil
}

// Return to initial working directory from labroot
func (nk *Koble) Cleanup() (err error) {
	if nk.LabRoot == "" {
		return nil
	}
	if nk.LabRoot != nk.InitialWorkDir {
		return os.Chdir(nk.InitialWorkDir)
	}
	return nil
}
