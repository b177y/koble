package netkit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTopographicalSort(t *testing.T) {
	digraph := newGraph()
	digraph.addNode("dns")
	digraph.addNode("router")
	digraph.addNode("client")
	err := digraph.addEdge("dns", "router")
	if err != nil {
		t.Fatal(err)
	}
	err = digraph.addEdge("client", "dns")
	if err != nil {
		t.Fatal(err)
	}
	order, err := digraph.sort()
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
	digraph := newGraph()
	digraph.addNode("dns")
	digraph.addNode("router")
	digraph.addNode("client")
	err := digraph.addEdge("dns", "router")
	if err != nil {
		t.Fatal(err)
	}
	err = digraph.addEdge("client", "dns")
	if err != nil {
		t.Fatal(err)
	}
	err = digraph.addEdge("router", "client")
	if err != nil {
		t.Fatal(err)
	}
	_, err = digraph.sort()
	if err == nil {
		t.Fatal(fmt.Errorf("Graph has cycle so should error."))
	}
}

func TestTopographicalNonExistentMachine(t *testing.T) {
	digraph := newGraph()
	digraph.addNode("dns")
	digraph.addNode("router")
	err := digraph.addEdge("client", "dns")
	if err == nil {
		t.Fatal(fmt.Errorf("Node client doesn't exist, addEdge should fail"))
	}
}

func TestTopographicalNonExistentDependency(t *testing.T) {
	digraph := newGraph()
	digraph.addNode("dns")
	digraph.addNode("router")
	err := digraph.addEdge("dns", "client")
	if err == nil {
		t.Fatal(fmt.Errorf("Dependency client doesn't exist, addEdge should fail"))
	}
}
