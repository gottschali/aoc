package main

import (
	"aoc2024"
	"fmt"
	"slices"
)

const (
	WALL  = '#'
	START = 'S'
	END   = 'E'
	EMPTY = '.'
)

func FindFirst(g [][]byte, target byte) (x int, y int) {
	for y, line := range g {
		x := slices.Index(line, target)
		if x != -1 {
			return x, y
		}
	}
	panic("not found")
}

type Point struct {
	x, y int
}

type Solver struct {
	part1           bool
	data            [][]byte
	cheatingAllowed bool
}

type State struct {
	x, y                 int
	cheated              bool
	cheatStart, cheatEnd Point
}

func (solver *Solver) neighborStates(s State) (res []State) {
	for _, p := range solver.neighborPoints(Point{s.x, s.y}) {
		x, y := p.x, p.y
		if solver.data[y][x] == WALL {
			if solver.cheatingAllowed && !s.cheated {
				res = append(res, State{
					x:          x,
					y:          y,
					cheated:    true,
					cheatStart: Point{s.x, s.y},
					cheatEnd:   Point{x, y},
				})
			}
			// if s.cheat < solver.maxCheat {
			// 	res = append(res, State{x, y, s.cheat + 1})
			// }
		}
		if solver.data[y][x] != WALL {
			res = append(res, State{x, y, s.cheated, s.cheatStart, s.cheatEnd})
		}
	}
	return
}

func (solver *Solver) neighborPoints(p Point) (res []Point) {
	deltas := []struct{ x, y int }{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	width := len(solver.data[0])
	height := len(solver.data)
	for _, delta := range deltas {
		x, y := p.x+delta.x, p.y+delta.y
		if x < 0 || x >= width {
			continue
		}
		if y < 0 || y >= height {
			continue
		}
		if solver.data[y][x] != WALL {
			res = append(res, Point{x, y})
		}
	}
	return
}

func (solver *Solver) solve() int {
	solver.cheatingAllowed = false
	best := solver.part1Inefficient(-1)
	fmt.Println(best)
	solver.cheatingAllowed = true
	return solver.part1Inefficient(best)
}

// if best is positive, this is the shortest path
// else we return the shortest path without cheating
func (solver *Solver) part1Inefficient(best int) int {
	sx, sy := FindFirst(solver.data, START)
	ex, ey := FindFirst(solver.data, END)
	start := State{sx, sy, false, Point{-1, -1}, Point{-1, -1}}
	dist := map[State]int{start: 0}
	q := []State{start}

	cheatCount := make(map[int]int)

	for len(q) > 0 {
		// pop an element
		node := q[0]
		q = q[1:]
		// fmt.Println(node, dist[node])
		if node.x == ex && node.y == ey {
			fmt.Println(node, dist[node])
			if !node.cheated {
				if best <= 0 {
					return dist[node]
				}
				break // no better cheats can be found
			} else {
				cheatCount[dist[node]] += 1
			}
			continue
			// return dist[node]
		}
		// add neighbors
		for _, n := range solver.neighborStates(node) {
			if _, ok := dist[n]; !ok {
				dist[n] = dist[node] + 1
				q = append(q, n)
			}
		}
	}
	// fmt.Println(cheatCount)
	res := 0
	for d, c := range cheatCount {
		// fmt.Println(best-d, c)
		if best-d >= 100 {
			res += c
		}
	}
	return res
}

func (solver *Solver) part1Efficient() int {
	ex, ey := FindFirst(solver.data, END)
	sx, sy := FindFirst(solver.data, START)
	fromStart := solver.pathsFrom(sx, sy)
	toEnd := solver.pathsFrom(ex, ey)
	bestWithout := toEnd[Point{sx, sy}] // == fromStart[Point{ex, ey}]
	fmt.Println("without cheating: ", bestWithout)
	cheatCount := make(map[int]int)
	for y, line := range solver.data {
		for x, b := range line {
			if b != WALL {
				continue
			}
			// We cheat here     b
			// 				    a#.
			for _, cheatStart := range solver.neighborPoints(Point{x, y}) {
				for _, cheatEnd := range solver.neighborPoints(Point{x, y}) {
					if cheatStart == cheatEnd {
						continue
					}
					cost := fromStart[cheatStart] + 2 + toEnd[cheatEnd]
					cheatCount[cost] += 1
				}
			}
		}
	}

	res := 0
	for d, c := range cheatCount {
		// if bestWithout-d > 0 {
		// 	fmt.Println(bestWithout-d, c)
		// }
		if bestWithout-d >= 100 {
			res += c
		}
	}
	return res
}

func Abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func FindAll(grid [][]byte, target byte) (res []Point) {
	for y, line := range grid {
		for x, b := range line {
			if b == target {
				res = append(res, Point{x, y})
			}
		}
	}
	return
}

func (solver *Solver) part2() int {
	ex, ey := FindFirst(solver.data, END)
	sx, sy := FindFirst(solver.data, START)
	solver.data[sy][sx] = EMPTY
	solver.data[ey][ex] = EMPTY
	width, height := len(solver.data[0]), len(solver.data)
	fromStart := solver.pathsFrom(sx, sy)
	toEnd := solver.pathsFrom(ex, ey)
	bestWithout := toEnd[Point{sx, sy}]
	if bestWithout != fromStart[Point{ex, ey}] {
		panic("shortest path invariant violated")
	}
	fmt.Println("without cheating: ", bestWithout)
	cheatCount := make(map[int]int)
	const CHEAT_LEN = 20
	for _, cheatStart := range FindAll(solver.data, EMPTY) {
		// start cheating    x
		// at position S    xxx
		// 				   xxSxx
		// exit cheatmode   xxx
		// at any x		     x
		x, y := cheatStart.x, cheatStart.y
		// consider all exit locations with manhattan distance <= CHEAT_LEN
		for dx := -CHEAT_LEN; dx <= CHEAT_LEN; dx += 1 {
			for dy := -CHEAT_LEN; dy <= CHEAT_LEN; dy += 1 {
				// how many steps taken
				L := Abs(dx) + Abs(dy)
				if L > CHEAT_LEN {
					continue
				}
				// bounds check
				if x+dx < 0 || x+dx >= width || y+dy < 0 || y+dy >= height {
					continue
				}
				cheatEnd := Point{x + dx, y + dy}
				if solver.data[cheatEnd.y][cheatEnd.x] == WALL {
					continue
				}
				if cheatStart == cheatEnd {
					continue
				}
				cost := fromStart[cheatStart] + L + toEnd[cheatEnd]
				cheatCount[cost] += 1
			}
		}
	}

	res := 0
	for d, c := range cheatCount {
		if bestWithout-d >= 50 {
			fmt.Println(bestWithout-d, c)
		}
		if bestWithout-d >= 100 {
			res += c
		}
	}
	return res
}
func (solver *Solver) pathsFrom(x, y int) map[Point]int {
	dist := map[Point]int{{x, y}: 0}
	q := []Point{{x, y}}
	// BFS
	for len(q) > 0 {
		// pop an element
		node := q[0]
		q = q[1:]
		// add neighbors
		for _, n := range solver.neighborPoints(node) {
			if _, ok := dist[n]; !ok {
				dist[n] = dist[node] + 1
				q = append(q, n)
			}
		}
	}
	return dist
}

func (solver *Solver) parse(path string) {
	solver.data = helper.ReadFileLines(path)
}

func main() {
	s := Solver{part1: true}
	s.parse("test")
	res1 := s.part1Efficient()
	fmt.Println("Part1: ", res1)
	res2 := s.part2()
	fmt.Println("Part2: ", res2)
	s.parse("input")
	// res1 = s.solve()
	res1 = s.part1Efficient()
	fmt.Println("(input) Part1: ", res1)
	// 1360
	res2 = s.part2()
	fmt.Println("Part2: ", res2)
}
