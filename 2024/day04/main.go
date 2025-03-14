package main

import (
	"bytes"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Grid [][]byte

func parse(path string) Grid {
	dat, err := os.ReadFile(path)
	check(err)
	// 10 -> \n
	dat, _ = bytes.CutSuffix(dat, []byte{'\n'})
	return bytes.Split(dat, []byte{'\n'})
}

func checkXmas(grid Grid, x, y int) (res int) {
	bs := []byte{'M', 'A', 'S'}
	deltas := [...]int{-1, 0, 1}
	for _, dy := range deltas {
		for _, dx := range deltas {
			if dy == 0 && dx == 0 {
				continue
			}
			for i, c := range bs {
				x_ := x + (i+1)*dx
				y_ := y + (i+1)*dy
				if y_ < 0 || y_ >= len(grid) {
					break
				}
				if x_ < 0 || x_ >= len(grid[y_]) {
					break
				}
				if grid[y_][x_] != c {
					break
				}
				if i == len(bs)-1 {
					res += 1
				}
			}
		}
	}
	return res
}

func part1(grid Grid) int {
	words := 0
	for y, line := range grid {
		for x, char := range line {
			if char != 'X' {
				continue
			}
			words += checkXmas(grid, x, y)
		}
	}
	return words
}

// M.S
// .A.
// M.S

// Not allowd
// M.S
// .A.
// S.M
func part2(grid Grid) int {
	words := 0
	for y, line := range grid {
		for x, char := range line {
			if char != 'A' {
				continue
			}
			if y == 0 || y >= len(grid)-1 {
				continue
			}
			if x == 0 || x >= len(grid[y])-1 {
				continue
			}
			a := grid[y-1][x-1]
			b := grid[y-1][x+1]
			c := grid[y+1][x+1]
			d := grid[y+1][x-1]
			if a == 'M' && c == 'S' || a == 'S' && c == 'M' {
				if b == 'M' && d == 'S' || b == 'S' && d == 'M' {
					words += 1
				}
			}
		}
	}
	return words
}

func main() {
	path := os.Args[1]
	grid := parse(path)
	// fmt.Println(grid)
	res := part1(grid)
	fmt.Println(res)

	res2 := part2(grid)
	fmt.Println(res2)
}
