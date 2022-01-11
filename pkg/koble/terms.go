package koble

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

type Terminal struct {
	Name    string            `mapstructure:"name"`
	Command []string          `mapstructure:"command"`
	Options map[string]string `mapstructure:"options,remain"`
}

func (t *Terminal) getArgs(opts LaunchOptions) ([]string, error) {
	var args []string
	for _, val := range t.Command {
		templ, err := template.New("term." + t.Name + "." + val).Option("missingkey=error").Parse(val)
		if err != nil {
			fmt.Println("error making template")
			return args, err
		}
		var tpl bytes.Buffer
		err = templ.Execute(&tpl, opts)
		if err != nil {
			fmt.Println("error executing template")
			return args, err
		} else {
			args = append(args, tpl.String())
		}
	}
	return args, nil
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
		Command: []string{"alacritty", "-e", "kob", "{{ .Command }}", "{{ .Machine }}", "--console"},
	},
	{
		Name:    "tmux",
		Command: []string{"tmux", "new-window", "-t", `{{ index .Options "session" }}`, `{{ .Command }}`},
		Options: map[string]string{"session": "koble"},
	},
	{
		Name:    "konsole",
		Command: []string{"konsole", "-e"},
	},
	{
		Name:    "gnome",
		Command: []string{"gnome-terminal", "--"},
	},
	{
		Name:    "kitty",
		Command: []string{"kitty"},
	},
	{
		Name:    "xterm",
		Command: []string{"xterm", "-e"},
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

	origCmd := append([]string{os.Args[0]}, os.Args[1:]...)
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
	fmt.Println("options are", opts.Options)
	termArgs, err := term.getArgs(opts)
	if err != nil {
		return err
	}
	fmt.Println("got term args", termArgs)
	log.Info("Relaunching current command in terminal with:", term.Name, termArgs)
	cmd := exec.Command(termArgs[0], termArgs[1:]...)
	cmd.Env = os.Environ()
	err = cmd.Start()
	return err
}
