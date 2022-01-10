package koble

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/b177y/koble/driver"
	"github.com/b177y/koble/util/topsort"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func fileExists(name string) (exists bool) {
	_, err := os.Lstat(name)
	return err == nil
}

func SaveLab(lab *Lab) error {
	lab.Name = ""
	lab.Directory = ""
	labYaml, err := yaml.Marshal(lab)
	if err != nil {
		return err
	}
	err = os.WriteFile("lab.yml", labYaml, 0644)
	return err
}

func RenderTable(headers []string, list [][]string) {
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
	table.AppendBulk(list)
	table.Render()
}

func MachineInfoToStringArr(machines []driver.MachineInfo, showNS bool) (mlist [][]string, headers []string) {
	headers = append(headers, "name")
	if showNS {
		headers = append(headers, "namespace")
	}
	// headers = append(headers, "lab")
	headers = append(headers, "image")
	// headers = append(headers, "networks")
	headers = append(headers, "state")

	for _, m := range machines {
		var minfo []string
		minfo = append(minfo, m.Name)
		if showNS {
			if len(m.Namespace) >= 8 {
				minfo = append(minfo, m.Namespace[:8])
			} else {
				minfo = append(minfo, m.Namespace)
			}
		}
		// minfo = append(minfo, m.Lab)
		minfo = append(minfo, filepath.Base(m.Image))
		// minfo = append(minfo, strings.Join(m.Networks, ","))
		minfo = append(minfo, m.State)
		// Add machine info to list of machines
		mlist = append(mlist, minfo)
	}
	return mlist, headers
}

func NetInfoToStringArr(networks []driver.NetInfo, showLab bool) (nlist [][]string, headers []string) {
	headers = append(headers, "name")
	// if showLab {
	// 	headers = append(headers, "lab")
	// }
	// headers = append(headers, "interface")
	headers = append(headers, "external")
	headers = append(headers, "gateway")
	headers = append(headers, "subnet")

	for _, n := range networks {
		var ninfo []string
		ninfo = append(ninfo, n.Name)
		// if showLab {
		// 	ninfo = append(ninfo, n.Lab)
		// }
		// ninfo = append(ninfo, n.Interface)
		ninfo = append(ninfo, strconv.FormatBool(n.External))
		ninfo = append(ninfo, n.Gateway)
		ninfo = append(ninfo, n.Subnet)
		// Add net info to list of networks
		nlist = append(nlist, ninfo)
	}
	return nlist, headers
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

func (nk *Koble) LaunchInTerm() error {
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
	cmd.Env = os.Environ()
	err = cmd.Start()
	return err
}

func orderMachines(machines map[string]driver.MachineConfig) (ordered map[string]driver.MachineConfig,
	err error) {
	dg := topsort.NewGraph()
	for name, m := range machines {
		dg.AddNode(name)
		for _, d := range m.Dependencies {
			if d == name {
				return ordered, fmt.Errorf("Machine %s cannot depend on itself!", name)
			} else if _, ok := machines[d]; !ok {
				return ordered, fmt.Errorf("Machine %s does not exist!", d)
			}
			dg.AddEdge(name, d)
		}
	}
	sorted, err := dg.Sort()
	if err != nil {
		return ordered, err
	}
	ordered = make(map[string]driver.MachineConfig, 0)
	for _, name := range sorted {
		ordered[name] = machines[name]
	}
	return ordered, nil
}

func multiHeading(heading string, list []string) (header string, value string) {
	if len(list) > 1 {
		header = heading + "(s)"
	} else {
		header = heading
	}
	value = strings.Join(list, ",")
	return header, value
}

// Walk up directories looking for a "lab.yml" file
func getLabRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for dir != "" {
		if fileExists(path.Join(dir, "lab.yml")) {
			return dir, nil
		}
		dir, _ = path.Split(dir)
		dir = strings.TrimRight(dir, "/")
	}
	return "", nil
}
