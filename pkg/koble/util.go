package koble

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/b177y/koble/pkg/driver"
	"github.com/b177y/koble/util/topsort"
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
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
	headers = append(headers, "image")
	headers = append(headers, "state")
	headers = append(headers, "created at")
	headers = append(headers, "started at")
	// headers = append(headers, "volumes")
	// headers = append(headers, "ports")
	headers = append(headers, "networks")

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
		minfo = append(minfo, filepath.Base(m.Image))
		minfo = append(minfo, m.State)
		if !m.CreatedAt.IsZero() {
			minfo = append(minfo, humanize.Time(m.CreatedAt))
		} else {
			minfo = append(minfo, "")
		}
		if !m.StartedAt.IsZero() {
			minfo = append(minfo, humanize.Time(m.StartedAt))
		} else {
			minfo = append(minfo, "")
		}
		// minfo = append(minfo, "") // TODO volumes
		// minfo = append(minfo, "") // TODO ports
		minfo = append(minfo, strings.Join(m.Networks, ", "))
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
