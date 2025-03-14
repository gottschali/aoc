package main

import (
	"aoc2024"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Calibration struct {
	value   int
	numbers []int
}
type Input []Calibration

func Parse(path string) (cs []Calibration) {
	for _, line := range helper.ReadFileLinesString(path) {
		cs = append(cs, ParseLine(line))
	}
	return
}
func ParseLine(bs string) Calibration {
	parts := strings.Split(bs, ": ")
	p1 := strings.Split(parts[1], " ")
	nums := make([]int, len(p1))
	for i := range nums {
		n, _ := strconv.Atoi(p1[i])
		nums[i] = n
	}
	v, _ := strconv.Atoi(parts[0])
	return Calibration{v, nums}
}

func Solve(c Calibration) bool {
	poss := []int{c.numbers[0]}
	// Brute Force baseline: 5.886s
	//
	// optimization 1:
	//   stop if we overshoot c.target
	//   note we can no longer preallocate fresh
	//
	//   1.373s
	for i, n := range c.numbers {
		if i == 0 {
			continue
		}
		fresh := make([]int, 0)
		for _, p := range poss {
			// (+)
			addition := p + n
			if addition <= c.value {
				fresh = append(fresh, addition)
			}
			// (*)
			multiplication := p * n
			if multiplication <= c.value {
				fresh = append(fresh, multiplication)
			}
		}
		poss = fresh
	}
	// logger.Printf("%d possibilities, input len %d", len(poss), len(c.numbers))
	for _, p := range poss {
		if p == c.value {
			return true
		}
	}
	return false
}
func Solve2(c Calibration) bool {
	poss := []int{c.numbers[0]}
	for i, n := range c.numbers {
		if i == 0 {
			continue
		}
		fresh := make([]int, 0)
		for _, p := range poss {
			// (+)
			addition := p + n
			if addition <= c.value {
				fresh = append(fresh, addition)
			}
			// (*)
			multiplication := p * n
			if multiplication <= c.value {
				fresh = append(fresh, multiplication)
			}

			// Concatenation
			concatenation := Concat(p, n)
			if concatenation <= c.value {
				fresh = append(fresh, concatenation)
			}
		}
		poss = fresh
	}
	// logger.Printf("%d possibilities, input len %d", len(poss), len(c.numbers))
	for _, p := range poss {
		if p == c.value {
			return true
		}
	}
	return false
}

func Concat(p, n int) int {
	var t int
	if n == 0 {
		t = 0
	} else {
		t = int(math.Floor(math.Log10(float64(n))))
	}
	e := int(math.Pow10(t + 1))
	return p*e + n
}

var logger = log.Default()

func Part1(cs Input) (res int) {
	for _, c := range cs {
		if Solve(c) {
			res += c.value
		}
	}
	return
}
func Part2(cs Input) (res int) {
	for _, c := range cs {
		if Solve2(c) {
			res += c.value
		}
	}
	return
}

func driver(path string) {
	fmt.Println(path)
	input := Parse(path)
	res1 := Part1(input)
	fmt.Println("result part1:\n", res1)
	res2 := Part2(input)
	fmt.Println("result part2:\n", res2)
}
func main() {
	driver("test")
	driver("input")
}
