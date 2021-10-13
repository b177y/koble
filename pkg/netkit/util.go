package netkit

import (
	"io/ioutil"
	"os"

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
