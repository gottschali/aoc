package main

import (
	helper "aoc2024"
	"fmt"
	"slices"
	"strings"
)

type Node string

type Edge struct {
	u, v Node
}

func Parse(path string) (edges []Edge, nodes map[Node]bool) {
	nodes = make(map[Node]bool)
	for _, bs := range helper.ReadFileLines(path) {
		u := Node(string(bs[:2]))
		v := Node(string(bs[3:]))
		nodes[u] = true
		nodes[v] = true
		edges = append(edges, Edge{u, v})
	}
	return
}

type Adj map[Node][]Node

func adjacencyList(edges []Edge) (a Adj) {
	a = make(Adj)
	for _, edge := range edges {
		a[edge.u] = append(a[edge.u], edge.v)
		a[edge.v] = append(a[edge.v], edge.u)
	}
	return
}

type Edgy map[Edge]bool

func edgeMap(edges []Edge) (e Edgy) {
	e = make(Edgy)
	for _, edge := range edges {
		e[edge] = true
		e[Edge{edge.v, edge.u}] = true
	}
	return
}

// find all triangles that contain  at least one
// node starting with t
func triangles(edg Edgy, adj Adj, prefix string) map[[3]Node]bool {
	seen := make(map[[3]Node]bool)
	for v, neighbors := range adj {
		if !strings.HasPrefix(string(v), prefix) {
			continue
		}
		for i, n1 := range neighbors {
			for j, n2 := range neighbors {
				if i == j {
					continue
				}
				if !edg[Edge{n1, n2}] {
					continue
				}
				s := []Node{v, n1, n2}
				slices.Sort(s)
				if seen[[3]Node(s)] {
					continue
				}
				seen[[3]Node(s)] = true
				// fmt.Println(v, n1, n2)
			}
		}
	}
	return seen
}

func Part1(edg Edgy, adj Adj) (res int) {
	return len(triangles(edg, adj, "t"))
}

func Join(ns []Node, sep string) string {
	res := ""
	for i, n := range ns {
		res += string(n)
		if i != len(ns)-1 {
			res += sep
		}
	}
	return res
}

func Part2(edg Edgy, adj Adj) (res string) {
	triangs := triangles(edg, adj, "")

	var worklist = [][]Node{}
	for t := range triangs {
		worklist = append(worklist, t[:])
	}
	for len(worklist) > 0 {
		fmt.Println(len(worklist[0]), len(worklist))
		if len(worklist) == 1 {
			lanParty := worklist[0]
			slices.Sort(lanParty)
			return Join(lanParty, ",")
		}
		next := [][]Node{}
		seen := make(map[string]bool)
		for _, clique := range worklist {
			for _, cand := range growClique(clique, edg, adj) {
				larger := append([]Node(nil), clique...)
				larger = append(larger, cand)
				slices.Sort(larger)
				stringRep := Join(larger, "")
				if !seen[stringRep] {
					next = append(next, larger)
					seen[stringRep] = true
				}
			}
		}
		worklist = next
	}
	return ""
}

func growClique(clique []Node, edg Edgy, adj Adj) (candidates []Node) {
	for v := range adj {
		for _, c := range clique {
			if v == c {
				goto done
			}
			if !edg[Edge{v, c}] {
				goto done
			}
		}
		candidates = append(candidates, v)
	done:
	}
	return
}

func main() {
	E, _ := Parse("input")
	edg := edgeMap(E)
	G := adjacencyList(E)
	res1 := Part1(edg, G)
	fmt.Println("res1: ", res1)
	res2 := Part2(edg, G)
	fmt.Println("res2: ", res2)
}
