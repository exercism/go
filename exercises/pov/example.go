package pov

import (
	"fmt"
	"log"
)

const testVersion = 2

type Graph struct {
	arcs   [][]int
	labels []string
	nodes  map[string]int
}

func New() *Graph { return &Graph{nodes: map[string]int{}} }

func (g *Graph) node(l string) int {
	if n, ok := g.nodes[l]; ok {
		return n
	}
	n := len(g.labels)
	g.labels = append(g.labels, l)
	g.arcs = append(g.arcs, nil)
	g.nodes[l] = n
	return n
}

func (g *Graph) AddNode(label string) {
	g.node(label)
}

func (g *Graph) AddArc(from, to string) {
	// use CI run to validate test cases
	if _, ok := g.nodes[to]; !ok {
		log.Println("AddArc from, to:", from, to)
		log.Fatal("test program promises bottom-up construction")
	}
	fr := g.node(from)
	g.arcs[fr] = append(g.arcs[fr], g.node(to))
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	// use CI run to validate test cases
	if !g.isTree(oldRoot) {
		log.Println("not a tree, or", oldRoot, "not root")
		log.Fatal("test program promises to pass root of a tree")
	}
	nr := g.nodes[newRoot]
	var f func(int) bool
	f = func(n int) (found bool) {
		if n == nr {
			return true
		}
		a := g.arcs[n]
		for i, to := range a {
			if f(to) {
				last := len(a) - 1
				a[i] = a[last]
				g.arcs[n] = a[:last]
				g.arcs[to] = append(g.arcs[to], n)
				return true
			}
		}
		return false
	}
	f(g.nodes[oldRoot])
	return g
}

func (g *Graph) ArcList() (s []string) {
	for fr, to := range g.arcs {
		for _, to := range to {
			s = append(s, fmt.Sprintf("%s -> %s", g.labels[fr], g.labels[to]))
		}
	}
	return
}

func (g *Graph) isTree(root string) bool {
	a := g.arcs
	v := make([]bool, len(a))
	nv := 0
	var df func(int) bool
	df = func(n int) bool {
		if v[n] {
			return false
		}
		v[n] = true
		nv++
		for _, to := range a[n] {
			if !df(to) {
				return false
			}
		}
		return true
	}
	return df(g.nodes[root]) && nv == len(a)
}
