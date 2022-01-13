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
	log "github.com/sirupsen/logrus"
)

type Terminal struct {
	Command []string          `koanf:"command"`
	Options map[string]string `koanf:"options"`
}

type TermConfig struct {
	Name string `koanf:"name"`
	// Whether to launch a terminal for start, attach and shell commands
	// default is true
	Launch bool `koanf:"launch"`
	// Whether to launch a shell instead of tty attach on lab / machine start
	// this only takes effect is LaunchTerms is true
	// default is false
	LaunchShell bool                `koanf:"launch_shell"`
	Terminals   map[string]Terminal `koanf:"terminals,remain"`
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
		Command: []string{"alacritty", "--hold", "-e", "{{ .Command }}"},
	},
	"tmux": {
		// create session if not exists, then attach in new window
		Command: []string{"tmux", "has-session", `-t={{ index .Options "session" | ShellEscape }}`,
			"||", "tmux", "new", "-d", "-s", `{{ index .Options "session" | ShellEscape }}`,
			";", "tmux", "new-window", `-t={{ index .Options "session" | ShellEscape }}`, `{{ .Command }}`},
		Options: map[string]string{"session": "koble", "var2": "example"},
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

func (nk *Koble) getTerm() (term Terminal, err error) {
	// Check default terminal list
	dTerm, dTermExists := defaultTerms[nk.Config.Terminal.Name]
	// Check custom terms first
	// This allows users to override default ones to add custom flags
	if t, ok := nk.Config.Terminal.Terminals[nk.Config.Terminal.Name]; ok {
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
	return term, fmt.Errorf("Terminal %s not found in config or default terminals.", nk.Config.Terminal.Name)
}

func (nk *Koble) LaunchInTerm(machine string) error {
	term, err := nk.getTerm()
	if err != nil {
		return err
	}
	if len(term.Command) == 0 {
		return errors.New("Terminal command must not be empty")
	}

	opts := LaunchOptions{}
	origCmd := os.Args
	added := false
	for i, a := range origCmd {
		if a == "--launch" || a == "--launch=true" {
			origCmd[i] = "--launch=false"
			added = true
		}
	}
	if !added {
		origCmd = append(origCmd, "--launch=false")
	}
	opts.Lab = nk.LabRoot
	opts.Namespace = nk.Config.Namespace
	opts.Machine = machine
	opts.Command = strings.Join(origCmd, " ")
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
	log.Info("Relaunching current command in terminal with:", termArgs)
	cmd := exec.Command("/bin/bash", "-c", termArgs)
	cmd.Env = os.Environ()
	err = cmd.Start()
	return err
}
