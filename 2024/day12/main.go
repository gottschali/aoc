package main

import (
	"aoc2024"
	"fmt"
	"log"
)

type solver struct {
	data [][]byte
}

func Parse(path string) solver {
	return solver{data: helper.ReadFileLines(path)}
}

type Point struct {
	x, y int
}

var deltas = []Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func (s *solver) neighbors(p Point) (ps []Point) {
	for _, d := range deltas {
		x, y := p.x+d.x, p.y+d.y
		if !(x >= 0 && x < len(s.data[0])) {
			continue
		}
		if !(y >= 0 && y < len(s.data)) {
			continue
		}
		ps = append(ps, Point{x, y})
	}
	return
}

func (s *solver) allNeighbors(p Point) (ps []Point) {
	for _, d := range deltas {
		x, y := p.x+d.x, p.y+d.y
		ps = append(ps, Point{x, y})
	}
	return
}

// m2 is merged into the map m1
func Combine(m1, m2 map[Point]bool) {
	for k, v := range m2 {
		m1[k] = v
	}
}

func (s *solver) dfs(p Point, symbol byte) (perimeter int, visited map[Point]bool) {
	visited = make(map[Point]bool)
	stack := []Point{p}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = true
		var neighbors []Point
		for _, n := range s.neighbors(node) {
			if s.data[n.y][n.x] == symbol {
				neighbors = append(neighbors, n)
			}
		}
		perimeter += 4 - len(neighbors)
		for _, neighbor := range neighbors {
			stack = append(stack, neighbor)
		}
	}
	return perimeter, visited
}

func (s *solver) Part1() (res int) {
	visited := make(map[Point]bool)
	for y, line := range s.data {
		for x, cell := range line {
			p := Point{x, y}
			if _, ok := visited[p]; ok {
				continue
			}
			perimeter, vis := s.dfs(p, cell)
			area := len(vis)
			Combine(visited, vis)
			res += area * perimeter
		}
	}
	return
}

func (s *solver) bounds(p Point) bool {
	if p.x < 0 || p.x >= len(s.data[0]) {
		return false
	}
	if p.y < 0 || p.y >= len(s.data) {
		return false
	}
	return true
}

type asf struct {
	p Point
	d int
}

func (s *solver) sides(region int, position Point, identifiers map[Point]int) (res int) {
	visited := make(map[asf]bool)

	candidates := make(map[asf]bool)
	for p, r := range identifiers {
		if r != region {
			continue
		}
		for _, n := range s.allNeighbors(p) {
			if r2 := identifiers[n]; r2 != region {
				candidates[asf{p, r2}] = true
			}
		}
	}
	logger := log.Default()
	for a := range candidates {
		if visited[a] {
			continue
		}
		position = a.p
		target := a.d
		// direction := a.d
		direction := 0
		for d, n := range s.allNeighbors(position) {
			if r2 := identifiers[n]; r2 == target {
				direction = d
				break
			}
		}
		// heading := deltas[direction]
		// target := identifiers[Point{position.x + heading.x, position.y + heading.y}]
		direction = (direction + 1) % 4
		endPosition := position
		endDirection := direction
		for {

			log.Println(position, direction, endPosition, endDirection, target)
			visited[asf{position, target}] = true
			heading := deltas[direction]
			next := Point{position.x + heading.x, position.y + heading.y}
			lookLeft := deltas[(direction+3)%4]
			onLeft := Point{position.x + lookLeft.x, position.y + lookLeft.y}
			if identifiers[onLeft] == target && identifiers[next] == region {
				logger.Println("forward", position)
				position = next
				next = Point{position.x + heading.x, position.y + heading.y}
			} else {
				logger.Println("turn left")
				res += 1
				direction = (direction + 3) % 4
			}

			// if identifiers[onLeft] == target {
			// 	logger.Println("turn left and go forward")
			// 	res += 1
			// 	direction = (direction + 3) % 4
			// 	position = onLeft
			// } else if s.bounds(next) && identifiers[next] != target {
			// 	logger.Println("forward", position)
			// 	position = next
			// 	next = Point{position.x + heading.x, position.y + heading.y}
			// } else {
			// 	logger.Println("turn right")
			// 	res += 1
			// 	direction = (direction + 1) % 4
			// }
			// !visited[position]
			if position == endPosition && direction == endDirection {
				break
			}
			// time.Sleep(time.Millisecond * 20)
		}
	}
	// fmt.Println("Region ", string(symbol))
	// fmt.Println("sides: ", res)
	return res
}

func (s *solver) Part2() (res int) {
	identifiers := make(map[Point]int)
	startpoint := make(map[int]Point)
	regionId := 0
	for y, line := range s.data {
		for x, cell := range line {
			p := Point{x, y}
			if _, ok := identifiers[p]; ok {
				continue
			}
			regionId += 1
			startpoint[regionId] = p
			_, vis := s.dfs(p, cell)
			for k := range vis {
				identifiers[k] = regionId
			}
			area := len(vis)
			side := s.sides(regionId, p, identifiers)
			fmt.Println(regionId, side, area)
			res += side * area
		}
	}
	return
}

func main() {
	// s := Parse("test1")
	// res1 := s.Part1()
	// fmt.Println("test1 part1: ", res1)
	// res2 := s.Part2()
	// fmt.Println("test1 part2: ", res2)
	// s := Parse("test2")
	// // res1 = s.Part1()
	// // fmt.Println("test2 part1: ", res1)
	// res2 := s.Part2()
	// fmt.Println("test2 part2: ", res2)
	s := Parse("testE")
	res2 := s.Part2()
	fmt.Println("testE part2: ", res2)
	// s = Parse("input")
	// res1 = s.Part1()
	// fmt.Println("part1: ", res1)
	// res2 = s.Part2()
	// fmt.Println("part2: ", res2)
}
