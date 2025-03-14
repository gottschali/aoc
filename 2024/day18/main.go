package main

import (
	"aoc2024"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func Parse(path string) (ps []Point) {
	for _, s := range helper.ReadFileLinesString(path) {
		nums := strings.Split(s, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		ps = append(ps, Point{x, y})
	}
	return
}

func adjacent(p Point) []Point {
	return []Point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}
}

func shortestPath(width, height int, walls map[Point]bool) int {
	start := Point{0, 0}
	end := Point{width - 1, height - 1}
	dists := map[Point]int{start: 0}
	worklist := []Point{start}
	for len(worklist) > 0 {
		node := worklist[0]
		if node == end {
			return dists[end]
		}
		worklist = worklist[1:]
		for _, n := range adjacent(node) {
			if n.x >= width || n.x < 0 {
				continue
			}
			if n.y >= height || n.y < 0 {
				continue
			}
			if walls[n] {
				continue
			}
			if _, ok := dists[n]; ok {
				continue
			}
			dists[n] = dists[node] + 1
			worklist = append(worklist, n)
		}
	}
	return -1
}

func SliceToMap(s []Point) map[Point]bool {
	m := make(map[Point]bool)
	for _, p := range s {
		m[p] = true
	}
	return m
}

func Reachable(width, height int, points []Point) bool {
	walls := SliceToMap(points)
	return shortestPath(width, height, walls) != -1
}

func Part2(width int, height int, points []Point) Point {
	low, high := 0, len(points)
	for low < high {
		mid := (high-low)/2 + low
		// fmt.Println(low, mid, high)
		if Reachable(width, height, points[:mid]) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return points[low-1]
}

func main() {
	testPoints := Parse("test")
	walls := SliceToMap(testPoints)
	fmt.Println(shortestPath(7, 7, walls))
	fmt.Println(shortestPath(7, 7, SliceToMap(testPoints[:12])))
	fmt.Println(shortestPath(7, 7, SliceToMap(testPoints[:20])))
	fmt.Println(shortestPath(7, 7, SliceToMap(testPoints[:21])))

	points := Parse("input")
	fmt.Println(shortestPath(71, 71, SliceToMap(points[:1024])))

	fmt.Println(Part2(7, 7, testPoints))
	fmt.Println(Part2(71, 71, points))
}
