package main

import (
	"aoc2024"
	"fmt"
)

const (
	EMPTY byte = '.'
)

func Parse(path string) solver {
	return solver{data: helper.ReadFileLines(path)}
}

type solver struct {
	data [][]byte
}

func (s *solver) Antennas() (ps map[byte][]helper.Point) {
	ps = make(map[byte][]helper.Point)
	for y, line := range s.data {
		for x, b := range line {
			if b != EMPTY {
				ps[b] = append(ps[b], helper.Point{X: x, Y: y})
			}
		}
	}
	return
}

func (s *solver) bounds(p helper.Point) bool {
	if p.X < 0 || p.X >= len(s.data[0]) {
		return false
	}
	if p.Y < 0 || p.Y >= len(s.data) {
		return false
	}
	return true
}

func Abs(n int) int {
	if n <= 0 {
		return -n
	} else {
		return n
	}
}

func (s *solver) antinode(a1, a2 helper.Point) (ps []helper.Point) {
	dx := Abs(a2.X - a1.X)
	dy := Abs(a2.Y - a1.Y)

	p1, p2 := a1, a2

	if p1.X < p2.X {
		p1.X -= dx
		p2.X += dx
	} else {
		p1.X += dx
		p2.X -= dx
	}
	if p1.Y < p2.Y {
		p1.Y -= dy
		p2.Y += dy
	} else {
		p1.Y += dy
		p2.Y -= dy
	}
	if s.bounds(p1) {
		ps = append(ps, p1)
	}

	if s.bounds(p2) {
		ps = append(ps, p2)
	}
	return
}

func (s *solver) Antinodes() map[helper.Point]int {
	ants := s.Antennas()
	antinodes := make(map[helper.Point]int)
	// different frequencies
	for _, as := range ants {
		for _, a1 := range as {
			for _, a2 := range as {
				if a1 == a2 {
					continue
				}
				for _, anti := range s.antinode(a1, a2) {
					// fmt.Println(string(freq), a1, a2, anti)
					antinodes[anti] += 1
				}
			}
		}
	}
	// fmt.Println(antinodes)
	return antinodes
}

func (s *solver) Part1() (res int) {
	antinodes := s.Antinodes()
	return len(antinodes)
}

func (s *solver) Part2() (res int) {
	ants := s.Antennas()
	antinodes := make(map[helper.Point]int)
	for _, as := range ants {
		for _, a1 := range as {
			for _, a2 := range as {
				if a1 == a2 {
					continue
				}
				dx := a2.X - a1.X
				dy := a2.Y - a1.Y
				delta := helper.Point{dx, dy}
				// iterate in one direction
				t := a1
				for s.bounds(t) {
					antinodes[t] += 1
					t = t.Add(&delta)
				}
				// iterate the other direction
				t = a1
				delta = delta.Scale(-1)
				for s.bounds(t) {
					antinodes[t] += 1
					t = t.Add(&delta)
				}
			}
		}
	}
	return len(antinodes)
}

func main() {
	s := Parse("test")
	res1 := s.Part1()
	fmt.Println("part1: ", res1)
	res2 := s.Part2()
	fmt.Println("part2: ", res2)
	s = Parse("input")
	res1 = s.Part1()
	fmt.Println("part1: ", res1)
	res2 = s.Part2()
	fmt.Println("part2: ", res2)
}
