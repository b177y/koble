// https://www.youtube.com/watch?v=cIBFEhD77b4
package netkit

import (
	"errors"
)

type graph struct {
	vertices []string
	edges    map[string][]string
	sorted   []string
}

func newGraph() graph {
	return graph{
		vertices: []string{},
		edges:    make(map[string][]string),
		sorted:   []string{},
	}
}

func (g *graph) addEdge(machine, dependency string) {
	_, ok := g.edges[machine]
	if !ok {
		g.edges[machine] = []string{dependency}
	} else {
		g.edges[machine] = append(g.edges[machine], dependency)
	}
}

func (g *graph) addNode(machine string) {
	g.vertices = append(g.vertices, machine)
}

func (g *graph) sort() ([]string, error) {
	n := len(g.vertices)
	// Make indegree as map of ints, counting connected vertices
	in_degree := make(map[string]int)
	for _, n := range g.vertices {
		in_degree[n] = 0
	}
	for k := range g.edges {
		for _, to := range g.edges[k] {
			in_degree[to]++
		}
	}
	queue := []string{}
	for _, k := range g.vertices {
		if in_degree[k] == 0 {
			queue = append(queue, k)
		}
	}

	index := 0
	order := make([]string, n)
	for len(queue) != 0 {
		at := queue[0]
		queue = queue[1:]
		order[index] = at
		index++
		for _, to := range g.edges[at] {
			in_degree[to]--
			if in_degree[to] == 0 {
				queue = append(queue, to)
			}
		}
	}
	if index != n {
		return []string{}, errors.New("Graph is not Directed Acyclic Graph :(")
	}
	return order, nil
}
