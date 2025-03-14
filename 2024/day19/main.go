package main

import (
	"aoc2024"
	"bytes"
	"fmt"
	"slices"
)

func parsePatterns(bs []byte) [][]byte {
	return bytes.Split(bs, []byte(", "))
}

func parseDesigns(bs []byte) [][]byte {
	return bytes.Split(bs, []byte("\n"))
}

type Problem struct {
	patterns, designs [][]byte
}

func Parse(path string) *Problem {
	sections := helper.ReadFileSections(path)
	patterns := parsePatterns(sections[0])
	designs := parseDesigns(sections[1])
	return &Problem{patterns, designs}
}

func Part(p *Problem, part1 bool) (res int) {
	for _, d := range p.designs {
		res += solve(d, p.patterns, part1)
	}
	return res
}

func solve(design []byte, patterns [][]byte, part1 bool) int {
	dp := make([]int, len(design)+1)
	dp[0] = 1
	for i := range design {
		i += 1
		for _, pat := range patterns {
			start := i - len(pat)
			if start < 0 || dp[start] == 0 {
				continue
			}
			// fmt.Println(string(design[start:i]), string(pat))
			if !slices.Equal(design[start:i], pat) {
				continue
			}
			dp[i] += dp[start]
			// break
		}
	}
	// fmt.Println(string(design), dp)
	if part1 {
		if dp[len(design)] != 0 {
			return 1
		} else {
			return 0
		}
	} else {
		return dp[len(design)]
	}
}

func main() {

	p := Parse("input")
	res1 := Part(p, true)
	fmt.Println("Part1: ", res1)
	res2 := Part(p, false)
	fmt.Println("Part2: ", res2)
}
