package koble

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/alessio/shellescape"
	"github.com/knadh/koanf/providers/confmap"
	log "github.com/sirupsen/logrus"
)

type Terminal struct {
	Command []string          `koanf:"command"`
	Options map[string]string `koanf:"options"`
}

type TermConfig struct {
	// name of default terminal to open
	// by default this is gnome
	Default string `koanf:"default"`
	// name of terminal to use for attach commands
	// by default this uses the terminal set for 'default'
	Attach string `koanf:"attach"`
	// name of terminal to use for shell commands
	// by default this uses the terminal set for 'default'
	Shell string `koanf:"shell"`
	// name of terminal to use for shell commands
	// by default this is set to 'this' (no terminal)
	Exec string `koanf:"exec"`
	// name of terminal to use for attaching on machine start
	// by default this uses the terminal set for 'default'
	MachineStart string `koanf:"machine_start"`
	// name of terminal to use for attaching on lab start
	// by default this uses the terminal set for 'default'
	LabStart string `koanf:"lab_start"`
	// extra terminal command and option definitions
	Terminals map[string]Terminal `koanf:"terminals,remain"`
}

func setTermDefaults() error {
	overrideCmds := []string{"attach", "shell", "machine_start", "lab_start"}
	overrideMap := make(map[string]interface{}, 0)
	for _, c := range overrideCmds {
		keyTerm := fmt.Sprintf("terminal.%s", c)
		if !Koanf.Exists(keyTerm) {
			overrideMap[keyTerm] = Koanf.String("terminal.default")
		}
	}
	return Koanf.Load(confmap.Provider(overrideMap, "."), nil)
}

func (t *Terminal) getArgs(opts LaunchOptions) (string, error) {
	var args []string
	for _, val := range t.Command {
		templ, err := template.New("term").
			Option("missingkey=error").
			Funcs(template.FuncMap{"ShellEscape": shellescape.Quote}).
			Parse(val)
		if err != nil {
			return "", err
		}
		var tpl bytes.Buffer
		err = templ.Execute(&tpl, opts)
		if err != nil {
			return "", err
		} else {
			args = append(args, tpl.String())
		}
	}
	return strings.Join(args, " "), nil
}

type LaunchOptions struct {
	Command   string
	Machine   string
	Lab       string
	Namespace string
	Options   map[string]string
}

var defaultTerms = map[string]Terminal{
	"alacritty": {
		Command: []string{"alacritty", "-e", "{{ .Command }}"},
	},
	"tmux": {
		// create session if not exists, then attach in new window
		Command: []string{"tmux", "has-session", `-t={{ index .Options "session" | ShellEscape }}`,
			"||", "tmux", "new", "-d", "-s", `{{ index .Options "session" | ShellEscape }}`,
			";", "tmux", "new-window", "-n", `{{ .Machine }}`, `-t={{ index .Options "session" | ShellEscape }}`, `{{ .Command }}`},
		Options: map[string]string{"session": "koble"},
	},
	"konsole": {
		Command: []string{"konsole", "--title", "{{ ShellEscape .Machine }}", "-e", "{{ .Command }}"},
	},
	"gnome": {
		Command: []string{"gnome-terminal", "--", "{{ .Command }}"},
	},
	"kitty": {
		Command: []string{"kitty", "{{ .Command }}"},
	},
	"xterm": {
		Command: []string{"xterm", "-e", "{{ .Command }}"},
	},
}

func (nk *Koble) getTerm(terminal string) (term Terminal, err error) {
	// Check default terminal list
	dTerm, dTermExists := defaultTerms[terminal]
	// Check custom terms first
	// This allows users to override default ones to add custom flags
	if t, ok := nk.Config.Terminal.Terminals[terminal]; ok {
		// if dterm merge
		if dTermExists {
			if len(t.Command) == 0 {
				t.Command = dTerm.Command
			}
			// merge options, priority to config over defaults
			optionsMap := dTerm.Options
			for key, val := range t.Options {
				optionsMap[key] = val
			}
			t.Options = optionsMap
		}
		return t, nil
	} else if dTermExists {
		return dTerm, nil
	}
	return term, fmt.Errorf("Terminal %s not found in config or default terminals.", terminal)
}

func (nk *Koble) LaunchInTerm(machine, terminal, command string) error {
	term, err := nk.getTerm(terminal)
	if err != nil {
		return err
	}
	if len(term.Command) == 0 {
		return errors.New("Terminal command must not be empty")
	}

	opts := LaunchOptions{}
	opts.Lab = nk.LabRoot
	opts.Namespace = nk.Config.Namespace
	opts.Machine = machine
	opts.Command = command
	if err != nil {
		return err
	}
	// use config default options with override from terminal
	opts.Options = term.Options
	if opts.Options == nil {
		opts.Options = make(map[string]string, 0)
	}
	for key, val := range nk.Config.TermOpts {
		opts.Options[key] = val
	}
	termArgs, err := term.getArgs(opts)
	if err != nil {
		return err
	}
	// check terminal exists
	_, err = exec.LookPath(term.Command[0])
	if err != nil {
		return fmt.Errorf("cannot find terminal %s in PATH: %w",
			term.Command[0], err)
	}
	log.Info("Relaunching current command in terminal with:", termArgs)
	cmd := exec.Command("/bin/bash", "-c", termArgs)
	cmd.Env = append(os.Environ(), "_KOBLE_IN_TERM=true")
	err = cmd.Start()
	return err
}
