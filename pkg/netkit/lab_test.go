package netkit_test

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/b177y/netkit/driver/mock"
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
		t.Errorf("Could not load newly initialised lab: %w", err)
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
	err := netkit.InitLab("", "test description", []string{}, []string{}, []string{})
	if err != nil {
		t.Errorf("Could not initialise lab in current directory: %w", err)
	}
	var lab netkit.Lab
	_, err = netkit.GetLab(&lab)
	if err != nil {
		t.Errorf("Could not load newly initialised lab: %w", err)
	}
	os.Chdir(TMPDIR)
	os.RemoveAll("newlab")
	// TODO test lab info is saved
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
	err := netkit.InitLab("newlab", "", []string{}, []string{}, []string{})
	if err != nil {
		t.Errorf("Could not initialise lab for test: %w", err)
	}
	os.Chdir("newlab")
	err = netkit.AddMachineToLab("testmachine", []string{"n1", "n2"}, "test-image")
	if err != nil {
		t.Errorf("Could not add machine to lab: %w", err)
	}
	var lab netkit.Lab
	_, err = netkit.GetLab(&lab)
	if err != nil {
		t.Errorf("Could not load lab for test: %w", err)
	}
	if len(lab.Machines) != 1 {
		t.Errorf("Lab should have exactly one machine")
	}
	m := lab.Machines[0]
	if m.Name != "testmachine" {
		t.Errorf("Machine name has not saved correctly")
	}
	if len(m.Networks) != 2 {
		t.Errorf("Machine should have two networks")
	}
	if m.Networks[0] != "n1" || m.Networks[1] != "n2" {
		t.Errorf("Machine networks have not saved correctly")
	}
	if m.Image != "test-image" {
		t.Errorf("Machine image has not saved correctly")
	}
	os.Chdir(TMPDIR)
	os.RemoveAll("newlab")
}

func TestAddNetworkToLab(t *testing.T) {
	err := netkit.InitLab("newlab", "", []string{}, []string{}, []string{})
	if err != nil {
		t.Errorf("Could not initialise lab for test: %w", err)
	}
	os.Chdir("newlab")
	ip, cidr, _ := net.ParseCIDR("172.16.10.1/24")
	err = netkit.AddNetworkToLab("testnet", true, ip, *cidr, false)
	if err != nil {
		t.Errorf("Could not add network to lab: %w", err)
	}
	var lab netkit.Lab
	_, err = netkit.GetLab(&lab)
	if err != nil {
		t.Errorf("Could not load lab for test: %w", err)
	}
	if len(lab.Networks) != 1 {
		t.Error("Lab should have exactly one network")
	}
	n := lab.Networks[0]
	if n.Name != "testnet" || n.External != true || n.IPv6 == true {
		t.Error("Network details has not saved correctly")
	}
	if n.Subnet != cidr.String() || n.Gateway.String() != ip.String() {
		t.Error("Network addresses not saved correctly")
	}
}

func TestValidate(t *testing.T) {
	t.Errorf("Test not made (feature not yet added)")
}

func TestLabStart(t *testing.T) {
	t.Errorf("Test not made")
	// test fails when no lab.yml
	// test works with simple abr
}

func TestLabDestroy(t *testing.T) {
	t.Errorf("Test not made")
	// test lab destroy succeeds and all machines are gone
}

func TestLabHalt(t *testing.T) {
	t.Errorf("Test not made")
	// test lab stops but machines are still present
}

func TestLabInfo(t *testing.T) {
	t.Errorf("Test not made")
	// load simple lab, check no errors
	// while running
	// while not running
	// after halted
	// after destroyed
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
	nk.Driver = new(mock.MockDriver)
	err = nk.Driver.SetupDriver(new(netkit.Config).Driver.ExtraConf)
	c := tm.Run()
	err = os.RemoveAll(TMPDIR)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	os.Exit(c)
}
