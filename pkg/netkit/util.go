package netkit

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
