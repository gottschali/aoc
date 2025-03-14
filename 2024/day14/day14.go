package main

import (
	"aoc2024"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func walk(x, v, s, t int) (r int) {
	r = (x + ((t%s)*(v%s))%s) % s
	if r < 0 {
		r += s
	}
	return
}

type Robot struct {
	px, py, vx, vy int
}

func Parse(dat []byte) (rs []Robot) {
	num := regexp.MustCompile(`-?[0-9]+`)

	for _, line := range bytes.Split(dat, []byte{'\n'}) {
		matches := num.FindAll(line, -1)
		if len(matches) < 4 {
			continue
		}
		px, err := strconv.Atoi(string(matches[0]))
		check(err)
		py, err := strconv.Atoi(string(matches[1]))
		check(err)
		vx, err := strconv.Atoi(string(matches[2]))
		check(err)
		vy, err := strconv.Atoi(string(matches[3]))
		check(err)

		rs = append(rs, Robot{px, py, vx, vy})
	}
	return rs
}

type Point struct {
	x, y int
}

func Solve(r Robot, w, h, t int) Point {
	return Point{
		x: walk(r.px, r.vx, w, t),
		y: walk(r.py, r.vy, h, t),
	}
}

func constructGrid(t int, data []Robot, grid [][]bool) [][]bool {
	height := len(grid)
	width := len(grid[0])
	for _, r := range data {
		x := walk(r.px, r.vx, width, t)
		y := walk(r.py, r.vy, height, t)
		grid[y][x] = true
	}
	return grid
}

func Part2(data []Robot) int {
	width := 101
	height := 103
	interesting := make(chan int)
	// thought I had to parallelize for a better search,
	// the problem was wrong width and height values
	// It looks like it is periodic after some time
	// lcm of all vx, vy maybe?
	NUM_WORKER := 8
	for worker := 1; worker <= NUM_WORKER; worker += 1 {
		go func(worker int, c chan int) {
			grid := make([][]bool, height)
			for i := range grid {
				grid[i] = make([]bool, width)
			}
			for t := worker; ; t += NUM_WORKER {
				// clear
				for y := range height {
					for x := range width {
						grid[y][x] = false
					}
				}
				grid = constructGrid(t, data, grid)
				if Heuristic(grid) {
					interesting <- t
				} else if (t-worker)%1_000_000 == 0 {
					fmt.Printf("[worker %d] time: %d\n", worker, t)
				}
			}
		}(worker, interesting)
	}

	for t := range interesting {
		// clear the terminal
		fmt.Println("time: ", t)
		fmt.Print("\033[H\033[2J")
		grid := make([][]bool, height)
		for i := range grid {
			grid[i] = make([]bool, width)
		}
		grid = constructGrid(t, data, grid)
		printGrid(grid)
		time.Sleep(time.Millisecond * 200)
	}

	return -1
}

func Heuristic(grid [][]bool) bool {
	H := len(grid)
	for y := 1; y < H-1; y += 1 {
		W := len(grid[y])
		for x := 1; x < W-1; x += 1 {
			// .x.
			// xxx
			// xxx
			if !grid[y][x] {
				continue
			}
			if grid[y-1][x] && grid[y+1][x] && grid[y][x-1] && grid[y][x+1] {
				if !grid[y-1][x-1] && !grid[y-1][x+1] {
					if grid[y+1][x-1] && grid[y+1][x+1] {
						return true
					}
				}
			}
		}
	}
	return false
}

func Part1(data []Robot) int {
	var q1, q2, q3, q4 int
	width := 101
	height := 103
	for _, r := range data {
		end := Solve(r, width, height, 101)

		// Quadrants Bookkeeping
		if end.x < width/2 {
			if end.y < height/2 {
				q1 += 1
			} else if end.y > height/2 {
				q2 += 1
			}
		} else if end.x > width/2 {
			if end.y < height/2 {
				q3 += 1
			} else if end.y > height/2 {
				q4 += 1
			}
		}
	}
	fmt.Println("quadrant counts: ", q1, q2, q3, q4)
	return q1 * q2 * q3 * q4
}

func printGrid(grid [][]bool) {
	for _, line := range grid {
		for _, c := range line {
			if c {
				fmt.Print(" ")
			} else {
				fmt.Print("█")
			}
		}
		fmt.Print("\n")
	}

}

func main() {
	var path string
	if len(os.Args) < 2 {
		path = "test.txt"
	} else {
		path = os.Args[1]
	}
	fmt.Println(path)
	dat := helper.ReadFile(path)
	data := Parse(dat)
	res1 := Part1(data)
	fmt.Println("Result part 1: \n", res1)
	res2 := Part2(data)
	fmt.Println("Result part 2: \n", res2)
}
