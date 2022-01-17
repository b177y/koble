package koble

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"

	"github.com/b177y/koble/pkg/driver"
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

func loadMachines(vpl *koanf.Koanf) (err error) {
	overrideMap := make(map[string]interface{}, 0)
	mNames := vpl.MapKeys("machines")
	for _, name := range mNames {
		keyHosthome := fmt.Sprintf("machines.%s.hosthome", name)
		if !vpl.Exists(keyHosthome) {
			overrideMap[keyHosthome] = false
		}
		keyHostlab := fmt.Sprintf("machines.%s.hostlab", name)
		if !vpl.Exists(keyHostlab) {
			overrideMap[keyHostlab] = true
		}
	}
	return vpl.Load(confmap.Provider(overrideMap, "."), nil)
}

func (nk *Koble) loadLab() (err error) {
	if nk.LabRoot == "" {
		log.Debug("not in lab directory so not loading lab config")
		return nil
	}
	vpl := koanf.New(".")
	labConfPath := filepath.Join(nk.LabRoot, "lab.yml")
	if err := vpl.Load(file.Provider(labConfPath), yaml.Parser()); err != nil {
		return fmt.Errorf("error reading lab.yml: %w", err)
	}

	if err = loadMachines(vpl); err != nil {
		return err
	}

	err = vpl.Unmarshal("", &nk.Lab)
	if err != nil {
		return fmt.Errorf("invalid lab config: %w", err)
	}
	// if lab does not set namespace, set it to lab path hash
	if vpl.String("namespace") == "" {
		vpl.Load(confmap.Provider(map[string]interface{}{
			"namespace": fmt.Sprintf("%x", md5.Sum([]byte(nk.LabRoot))),
		}, ""), nil)
	}

	nk.Lab.Machines, err = orderMachines(nk.Lab.Machines)
	if err != nil {
		return fmt.Errorf("could not order lab machines by dependency: %w", err)
	}

	nk.Lab.Name = filepath.Base(nk.LabRoot)
	nk.Lab.Directory = nk.LabRoot

	if err := validator.New().Struct(nk.Lab); err != nil {
		return fmt.Errorf("error validating lab.yml: %w", err)
	}
	if err := Koanf.Merge(vpl); err != nil {
		return fmt.Errorf("Could not merge lab driver config to default driver config")
	}
	return nil
}

func validateConfig(conf Config) error {
	var err error
	err = validator.New().Struct(conf)
	if err != nil {
		return fmt.Errorf("error validating config.yml: %w", err)
	}
	if conf.Verbosity != 0 && conf.Quiet {
		return fmt.Errorf("verbose and quiet options cannot be used together")
	}
	if conf.Launch.LabStart && conf.Terminal.LabStart == "this" {
		return fmt.Errorf("terminal 'this' when launch on lab start is enabled")
	}
	return nil
}

func (nk *Koble) processConfig() error {
	color.NoColor = nk.Config.NoColor
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
		return fmt.Errorf("verbosity level %d is not valid",
			nk.Config.Verbosity)
	}
	if nk.Config.Quiet {
		log.SetLevel(log.ErrorLevel)
	}
	var err error
	nk.Driver, err = driver.GetDriver(nk.Config.Driver.Name,
		nk.Config.Driver.ExtraConf)
	return err
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
		if err := nk.loadLab(); err != nil {
			return nil, err
		}
	}
	if err := setTermDefaults(); err != nil {
		return nil, err
	}
	// process flags
	for _, pf := range kFlags {
		if err := Koanf.Load(pf, nil); err != nil {
			return nil, fmt.Errorf("loading config from flag: %w", err)
		}
	}
	if err := Koanf.Unmarshal("", &nk.Config); err != nil {
		return nil, fmt.Errorf("error loading config.yml: %w", err)
	}

	// validate config
	if err := validateConfig(nk.Config); err != nil {
		return nil, err
	}

	// process config
	if err := nk.processConfig(); err != nil {
		return nil, err
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
