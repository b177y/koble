package netkit

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/b177y/netkit/driver"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func fileExists(name string) (exists bool, err error) {
	if _, err := os.Stat(name); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func getLab(lab *Lab) (exists bool, err error) {
	exists, err = fileExists("lab.yml")
	if err != nil {
		// not necessarily false, so check err before exists
		return false, err
	}
	if !exists {
		return false, nil
	}
	f, err := ioutil.ReadFile("lab.yml")
	if err != nil {
		return true, err
	}
	err = yaml.Unmarshal(f, &lab)
	if err != nil {
		return true, err
	}
	dir, err := os.Getwd()
	if err != nil {
		return true, err
	}
	lab.Name = filepath.Base(dir)
	return true, nil
}

func saveLab(lab *Lab) error {
	labYaml, err := yaml.Marshal(lab)
	if err != nil {
		return err
	}
	err = os.WriteFile("lab.yml", labYaml, 0644)
	return err
}

type NetkitError struct {
	Err   error
	From  string
	Doing string
	Extra string
}

func (ne NetkitError) Error() string {
	errString := ""
	if ne.From != "" {
		errString += fmt.Sprintf("[%s] :", ne.From)
	}
	errString += ne.Err.Error()
	if ne.Doing != "" {
		errString += fmt.Sprintf("\nWhile doing: %s\n", ne.Doing)
	}
	errString += ne.Extra
	return errString
}

func NewError(err error, from, doing, extra string) NetkitError {
	return NetkitError{
		Err:   err,
		From:  from,
		Doing: doing,
		Extra: extra,
	}
}

func RenderTable(headers []string, mlist [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(mlist)
	table.Render()
}

func MachineInfoToStringArr(machines []driver.MachineInfo, showLab bool) (mlist [][]string, headers []string) {
	headers = append(headers, "name")
	if showLab {
		headers = append(headers, "lab")
	}
	headers = append(headers, "image")
	headers = append(headers, "networks")
	headers = append(headers, "state")

	for _, m := range machines {
		var minfo []string
		minfo = append(minfo, m.Name)
		if showLab {
			minfo = append(minfo, m.Lab)
		}
		minfo = append(minfo, m.Image)
		minfo = append(minfo, strings.Join(m.Networks, ","))
		minfo = append(minfo, m.State)
		// Add machine info to list of machines
		mlist = append(mlist, minfo)
	}
	return mlist, headers
}

func LaunchInTerm() error {
	args := []string{"-e"}
	args = append(args, os.Args...)
	added := false
	for i, a := range args {
		if a == "--terminal" {
			args[i] = "--console"
			added = true
		}
	}
	if !added {
		args = append(args, "--console")
	}
	term := "alacritty"
	log.Debug("Relaunching current command in terminal with:", term, args)
	cmd := exec.Command(term, args...)
	err := cmd.Start()
	return err
}
