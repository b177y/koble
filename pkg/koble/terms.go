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
	Name    string            `mapstructure:"name"`
	Command []string          `mapstructure:"command"`
	Options map[string]string `mapstructure:"options"`
}

func (t *Terminal) getArgs(opts LaunchOptions) (string, error) {
	var args []string
	for _, val := range t.Command {
		templ, err := template.New("term." + t.Name + "." + val).
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

var defaultTerms = []Terminal{
	{
		Name:    "alacritty",
		Command: []string{"alacritty", "-e", "{{ .Command }}"},
	},
	{
		// create session if not exists, then attach in new window
		Name: "tmux",
		Command: []string{"tmux", "has-session", `-t={{ index .Options "session" | ShellEscape }}`,
			"||", "tmux", "new", "-d", "-s", `{{ index .Options "session" | ShellEscape }}`,
			";", "tmux", "new-window", `-t={{ index .Options "session" | ShellEscape }}`, `{{ .Command }}`},
		Options: map[string]string{"session": "koble"},
	},
	{
		Name:    "konsole",
		Command: []string{"konsole", "--title", "{{ ShellEscape .Machine }}", "-e", "{{ .Command }}"},
	},
	{
		Name:    "gnome",
		Command: []string{"gnome-terminal", "--", "{{ .Command }}"},
	},
	{
		Name:    "kitty",
		Command: []string{"kitty", "{{ .Command }}"},
	},
	{
		Name:    "xterm",
		Command: []string{"xterm", "-e", "{{ .Command }}"},
	},
}

func (nk *Koble) getTerm() (term Terminal, err error) {
	// Check custom terms first
	// This allows users to override default ones to add custom flags
	for _, t := range nk.Config.Terms {
		if t.Name == nk.Config.Terminal {
			return t, nil
		}
	}
	// Check default terminal list
	for _, t := range defaultTerms {
		if t.Name == nk.Config.Terminal {
			return t, nil
		}
	}
	return term, fmt.Errorf("Terminal %s not found in config or default terminals.", nk.Config.Terminal)
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
		if a == "--terminal" {
			origCmd[i] = "--console"
			added = true
		}
	}
	if !added {
		origCmd = append(origCmd, "--console")
	}
	opts.Lab = nk.LabRoot
	opts.Namespace = nk.Config.Namespace
	opts.Machine = machine
	opts.Command = strings.Join(origCmd, " ")
	// use config default options with override from terminal
	opts.Options = term.Options
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
