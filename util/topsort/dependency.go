// https://www.youtube.com/watch?v=cIBFEhD77b4
package topsort

import (
	"errors"
	"fmt"
)

type graph struct {
	vertices []string
	edges    map[string][]string
	sorted   []string
}

func NewGraph() graph {
	return graph{
		vertices: []string{},
		edges:    make(map[string][]string),
		sorted:   []string{},
	}
}

func (g *graph) HasNode(node string) bool {
	for _, n := range g.vertices {
		if n == node {
			return true
		}
	}
	return false
}

func (g *graph) AddEdge(machine, dependency string) error {
	if !g.HasNode(machine) {
		return fmt.Errorf("Machine %s has not been added as a node.", machine)
	} else if !g.HasNode(dependency) {
		return fmt.Errorf("Dependency %s has not been added as a node.", dependency)
	}
	if _, ok := g.edges[machine]; !ok {
		g.edges[machine] = []string{dependency}
	} else {
		g.edges[machine] = append(g.edges[machine], dependency)
	}
	return nil
}

func (g *graph) AddNode(machine string) {
	g.vertices = append(g.vertices, machine)
}

func (g *graph) Sort() ([]string, error) {
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
	reversed := []string{}
	for i := range order {
		index = n - 1 - i
		reversed = append(reversed, order[index])
	}
	return reversed, nil
}
