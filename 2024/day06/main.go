package main

import (
	"aoc2024"
	"fmt"
)

type Grid [][]byte

const (
	Obstacle = '#'
	Free     = '.'
	Guard    = '^'
)

var directions = []helper.Point{
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
	{-1, 0}, // LEFT
}

func Part1(grid Grid) (res int) {
	return len(visited(grid))
}

func visited(grid Grid) map[helper.Point]bool {
	height := len(grid)
	width := len(grid[0])
	inBounds := func(p helper.Point) bool {
		return 0 <= p.X && p.X < width && 0 <= p.Y && p.Y < height
	}

	visited := make(map[helper.Point]bool)

	// find the guard
	var guard helper.Point
	for y := range height {
		for x := range width {
			if grid[y][x] == Guard {
				guard = helper.Point{X: x, Y: y}
			}
		}
	}

	turn := func() func() helper.Point {
		dir := len(directions) - 1
		return func() helper.Point {
			dir = (dir + 1) % len(directions)
			return directions[dir]
		}
	}()
	direction := turn()

	visited[guard] = true
	nextPos := guard.Add(&direction)
	for ; inBounds(nextPos); nextPos = guard.Add(&direction) {
		if grid[nextPos.Y][nextPos.X] == Obstacle {
			direction = turn()
		} else {
			guard = nextPos
		}
		visited[guard] = true
	}
	return visited
}

type PointDir struct {
	p, v helper.Point
}

func loops(grid Grid) (loops bool) {
	height := len(grid)
	width := len(grid[0])
	inBounds := func(p helper.Point) bool {
		return 0 <= p.X && p.X < width && 0 <= p.Y && p.Y < height
	}
	visited := make(map[PointDir]bool)

	// find the guard
	var guard helper.Point
	for y := range height {
		for x := range width {
			if grid[y][x] == Guard {
				guard = helper.Point{X: x, Y: y}
			}
		}
	}

	turn := func() func() helper.Point {
		dir := len(directions) - 1
		return func() helper.Point {
			dir = (dir + 1) % len(directions)
			return directions[dir]
		}
	}()
	direction := turn()

	visited[PointDir{guard, direction}] = true
	nextPos := guard.Add(&direction)
	for ; inBounds(nextPos); nextPos = guard.Add(&direction) {
		if grid[nextPos.Y][nextPos.X] == Obstacle {
			direction = turn()
		} else {
			guard = nextPos
		}
		if visited[PointDir{guard, direction}] {
			return true
		}
		visited[PointDir{guard, direction}] = true
	}
	return false
}

func Part2(grid Grid) (res int) {
	vs := visited(grid)
	// Optimization
	//  place obstacles only on visited area
	//  BROKEN does not produce the right result
	//  (as i got wiht brute force)
	//  maybe the loops check is broken?
	for p := range vs {
		grid[p.Y][p.X] = Obstacle
		if loops(grid) {
			res += 1
		}
		grid[p.Y][p.X] = Free
	}
	return res
}

func driver(path string) {
	fmt.Println(path)
	grid := helper.ReadFileLines(path)
	res1 := Part1(grid)
	fmt.Println("result part1:\n", res1)

	res2 := Part2(grid)
	fmt.Println("result part2:\n", res2)
}
func main() {
	driver("test")
	driver("input")
}
