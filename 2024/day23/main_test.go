package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	edges, _ := Parse("test")

	first := Edge{"kh", "tc"}
	last := Edge{"td", "yn"}

	if first != edges[0] {
		t.Errorf("first is %v, want %v", edges[0], first)
	}

	if last != edges[len(edges)-1] {
		t.Errorf("last is %v, want %v", edges[len(edges)-1], last)
	}

}

func Test1(t *testing.T) {
	E, _ := Parse("test")
	edg := edgeMap(E)
	G := adjacencyList(E)
	res1 := Part1(edg, G)
	exp := 7
	if res1 != exp {
		t.Errorf("for test part1 got %d but want %d", res1, exp)
	}

	cs := growClique([]Node{"co", "de", "ka"}, edg, G)
	fmt.Println(cs)
}

func Test2(t *testing.T) {
	E, _ := Parse("test")
	edg := edgeMap(E)
	G := adjacencyList(E)
	res2 := Part2(edg, G)
	exp := "co,de,ka,ta"
	if res2 != exp {
		t.Errorf("for test part2 got %s but want %s", res2, exp)
	}
}
