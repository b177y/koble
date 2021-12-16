package topsort_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/b177y/netkit/util/topsort"
)

func TestTopographicalSort(t *testing.T) {
	digraph := topsort.NewGraph()
	digraph.AddNode("dns")
	digraph.AddNode("router")
	digraph.AddNode("client")
	err := digraph.AddEdge("dns", "router")
	if err != nil {
		t.Fatal(err)
	}
	err = digraph.AddEdge("client", "dns")
	if err != nil {
		t.Fatal(err)
	}
	order, err := digraph.Sort()
	if err != nil {
		t.Fatal(err)
	}
	correctOrder := []string{"router", "dns", "client"}
	if !reflect.DeepEqual(order, correctOrder) {
		t.Fatal(fmt.Errorf("Order should be %v not %v",
			correctOrder, order))
	}
}

func TestTopographicalCyclic(t *testing.T) {
	digraph := topsort.NewGraph()
	digraph.AddNode("dns")
	digraph.AddNode("router")
	digraph.AddNode("client")
	err := digraph.AddEdge("dns", "router")
	if err != nil {
		t.Fatal(err)
	}
	err = digraph.AddEdge("client", "dns")
	if err != nil {
		t.Fatal(err)
	}
	err = digraph.AddEdge("router", "client")
	if err != nil {
		t.Fatal(err)
	}
	_, err = digraph.Sort()
	if err == nil {
		t.Fatal(fmt.Errorf("Graph has cycle so should error."))
	}
}

func TestTopographicalNonExistentMachine(t *testing.T) {
	digraph := topsort.NewGraph()
	digraph.AddNode("dns")
	digraph.AddNode("router")
	err := digraph.AddEdge("client", "dns")
	if err == nil {
		t.Fatal(fmt.Errorf("Node client doesn't exist, addEdge should fail"))
	}
}

func TestTopographicalNonExistentDependency(t *testing.T) {
	digraph := topsort.NewGraph()
	digraph.AddNode("dns")
	digraph.AddNode("router")
	err := digraph.AddEdge("dns", "client")
	if err == nil {
		t.Fatal(fmt.Errorf("Dependency client doesn't exist, addEdge should fail"))
	}
}
