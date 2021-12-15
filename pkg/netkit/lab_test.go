package netkit_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/b177y/netkit/pkg/netkit"
)

var nk *netkit.Netkit

var TESTPATH string
var TMPDIR string

func TestInitLabNewDir(t *testing.T) {
	err := netkit.InitLab("newlab", "", []string{}, []string{}, []string{})
	if err != nil {
		t.Errorf("Could not initialise lab in new directory: %w", err)
	}
	var lab netkit.Lab
	os.Chdir("newlab")
	_, err = netkit.GetLab(&lab)
	if err != nil {
		t.Errorf("Could not create newly initialised lab: %w", err)
	}
	os.Chdir(TMPDIR)
	err = os.RemoveAll("newlab")
	if err != nil {
		t.Error(err)
	}
}

func TestInitLabCurrDir(t *testing.T) {
	os.Mkdir("newlab", 0700)
	os.Chdir("newlab")
	err := netkit.InitLab("", "", []string{}, []string{}, []string{})
	if err != nil {
		t.Errorf("Could not initialise lab in current directory: %w", err)
	}
	var lab netkit.Lab
	_, err = netkit.GetLab(&lab)
	if err != nil {
		t.Errorf("Could not create newly initialised lab: %w", err)
	}
	os.Chdir(TMPDIR)
	err = os.RemoveAll("newlab")
	if err != nil {
		t.Error(err)
	}
}

func TestInitLabNewDirExists(t *testing.T) {
	os.Mkdir("newlab", 0700)
	err := netkit.InitLab("newlab", "", []string{}, []string{}, []string{})
	if err == nil {
		t.Error("Managed to make new lab where the directory already exists")
	}
	err = os.RemoveAll("newlab")
	if err != nil {
		t.Error(err)
	}
}

func TestInitLabCurrDirExists(t *testing.T) {
	os.Mkdir("newlab", 0700)
	os.Chdir("newlab")
	os.Create("lab.yml")
	err := netkit.InitLab("", "", []string{}, []string{}, []string{})
	if err == nil {
		t.Error("Managed to init new lab which already has lab.yml")
	}
	os.Chdir(TMPDIR)
	err = os.RemoveAll("newlab")
	if err != nil {
		t.Error(err)
	}
}

func TestAddMachineToLab(t *testing.T) {
	t.Errorf("Test not made")
}

func TestAddNetworkToLab(t *testing.T) {
	t.Errorf("Test not made")
}

func TestValidate(t *testing.T) {
	t.Errorf("Test not made")
}

func TestLabStart(t *testing.T) {
	t.Errorf("Test not made")
}

func TestLabDestroy(t *testing.T) {
	t.Errorf("Test not made")
}

func TestLabHalt(t *testing.T) {
	t.Errorf("Test not made")
}

func TestLabInfo(t *testing.T) {
	t.Errorf("Test not made")
}

func TestMain(tm *testing.M) {
	cwd, _ := os.Getwd()
	TESTPATH = filepath.Join(cwd, "..", "..", "tests")
	TMPDIR = filepath.Join(TESTPATH, "temp")
	err := os.MkdirAll(TMPDIR, 0700)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	os.Chdir(TMPDIR)
	nk, err = netkit.NewNetkit("GLOBAL")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	c := tm.Run()
	err = os.RemoveAll(TMPDIR)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	os.Exit(c)
}
