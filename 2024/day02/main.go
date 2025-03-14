package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse(path string) [][]int {
	dat, err := os.ReadFile(path)
	check(err)
	mat := make([][]int, 0)
	for _, line := range strings.Split(string(dat), "\n") {
		if len(line) == 0 {
			continue
		}
		l := make([]int, 0)
		for _, num := range strings.Split(line, " ") {
			n, err := strconv.Atoi(strings.TrimSpace(num))
			check(err)
			l = append(l, n)
		}
		mat = append(mat, l)
	}
	return mat
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func safe(xs []int) bool {
	increasing := xs[0] < xs[1]
	for i := range xs {
		if i == len(xs)-1 {
			break
		}
		if xs[i] == xs[i+1] {
			return false
		}
		if xs[i] < xs[i+1] != increasing {
			return false
		}
		if Abs(xs[i]-xs[i+1]) > 3 {
			return false
		}
	}
	return true
}

func part1(mat [][]int) int {
	safeReports := 0
	for _, report := range mat {
		if safe(report) {
			safeReports += 1
		}
	}
	return safeReports
}

// omg...
// report:  [8 6 4 4 1]
// 0 [6 4 4 1]
// 1 [6 4 1 1]
// 2 [6 4 1 1]
// 3 [6 4 1 1]
// 4 [6 4 1 1]
func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func safe2(report []int) bool {
	for i := 0; i < len(report); i += 1 {
		// delete element i
		rep := make([]int, len(report)-1)
		copy(rep[:i], report[:i])
		copy(rep[i:], report[i+1:])
		if safe(rep) {
			return true
		}
	}
	// without removing any level
	return safe(report)
}

func part2(mat [][]int) int {
	safeReports := 0
	for _, report := range mat {
		if safe2(report) {
			safeReports += 1
		}
	}
	return safeReports
}

func main() {
	path := os.Args[1]
	mat := parse(path)
	res1 := part1(mat)
	fmt.Println(res1)
	res2 := part2(mat)
	fmt.Println(res2)
}
