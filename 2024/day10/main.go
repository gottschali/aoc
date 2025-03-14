package main

import (
	"aoc2024"
	"fmt"
	"strconv"
)

func Parse(path string) (grid [][]int) {
	grid = make([][]int, 0)
	for i, line := range helper.ReadFileLines(path) {
		grid = append(grid, make([]int, len(line)))
		for j, b := range line {
			n, _ := strconv.Atoi(string(b))
			grid[i][j] = n
		}
	}
	return grid
}

type Point struct {
	x, y int
}

var delta = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func neighbors(p Point, g [][]int) (ps []Point) {
	for _, d := range delta {
		x, y := p.x+d.x, p.y+d.y
		if !(x >= 0 && x < len(g[0])) {
			continue
		}
		if !(y >= 0 && y < len(g)) {
			continue
		}
		ps = append(ps, Point{x, y})
	}
	return
}

func explore(p Point, g [][]int, part1 bool) (res int) {
	q := []Point{p}
	visited := make(map[Point]bool)
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		height := g[n.y][n.x]
		if height == 9 {
			res += 1
			continue
		}
		for _, neigh := range neighbors(n, g) {
			if _, ok := visited[neigh]; part1 && ok {
				continue
			}
			nextHeight := g[neigh.y][neigh.x]
			if nextHeight == height+1 {
				visited[neigh] = true
				q = append(q, neigh)
			}
		}

	}
	return
}

func Solve(g [][]int, part1 bool) int {
	res := 0
	for y, line := range g {
		for x, b := range line {
			if b == 0 {
				res += explore(Point{x, y}, g, part1)
			}
		}
	}
	return res
}

func main() {
	s := Parse("test")
	res1 := Solve(s, true)
	fmt.Println("part1: ", res1)
	res2 := Solve(s, false)
	fmt.Println("part2: ", res2)
	s = Parse("input")
	res1 = Solve(s, true)
	fmt.Println("part1: ", res1)
	res2 = Solve(s, false)
	fmt.Println("part2: ", res2)
}
