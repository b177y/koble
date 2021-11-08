package netkit

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
	lab.Machines, err = orderMachines(lab.Machines)
	if err != nil {
		return true, err
	}
	lab.Name = filepath.Base(dir)
	lab.Directory = dir
	return true, nil
}

func saveLab(lab *Lab) error {
	lab.Name = ""
	lab.Directory = ""
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

func NetInfoToStringArr(networks []driver.NetInfo, showLab bool) (nlist [][]string, headers []string) {
	headers = append(headers, "name")
	if showLab {
		headers = append(headers, "lab")
	}
	headers = append(headers, "interface")
	headers = append(headers, "external")
	headers = append(headers, "gateway")
	headers = append(headers, "subnet")

	for _, n := range networks {
		var ninfo []string
		ninfo = append(ninfo, n.Name)
		if showLab {
			ninfo = append(ninfo, n.Lab)
		}
		ninfo = append(ninfo, n.Interface)
		ninfo = append(ninfo, strconv.FormatBool(n.External))
		ninfo = append(ninfo, n.Gateway)
		ninfo = append(ninfo, n.Subnet)
		// Add net info to list of networks
		nlist = append(nlist, ninfo)
	}
	return nlist, headers
}

func (nk *Netkit) getTerm() (term Terminal, err error) {
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

func (nk *Netkit) LaunchInTerm() error {
	term, err := nk.getTerm()
	if err != nil {
		return err
	}
	var args []string
	if tl := len(term.Command); tl == 0 {
		return errors.New("Terminal command must not be empty")
	} else if tl != 1 {
		args = append(args, term.Command[1:]...)
	}
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
	log.Info("Relaunching current command in terminal with:", term.Name, args)
	cmd := exec.Command(term.Command[0], args...)
	err = cmd.Start()
	return err
}

func orderMachines(machines []driver.Machine) (ordered []driver.Machine,
	err error) {
	dg := newGraph()
	mappedMachines := map[string]driver.Machine{}
	for _, m := range machines {
		mappedMachines[m.Name] = m
	}
	for _, m := range machines {
		dg.addNode(m.Name)
		for _, d := range m.Dependencies {
			if d == m.Name {
				return ordered, fmt.Errorf("Machine %s cannot depend on itself!", m.Name)
			} else if _, ok := mappedMachines[d]; !ok {
				return ordered, fmt.Errorf("Machine %s does not exist!", d)
			}
			dg.addEdge(m.Name, d)
		}
	}
	sorted, err := dg.sort()
	if err != nil {
		return ordered, err
	}
	for _, m := range sorted {
		ordered = append(ordered, mappedMachines[m])
	}
	return ordered, nil
}
