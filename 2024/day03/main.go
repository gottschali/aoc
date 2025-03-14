package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Pair struct {
	x, y int
}

func parse(path string) []Pair {
	dat, err := os.ReadFile(path)
	check(err)
	sdat := string(dat)
	mul := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	res := []Pair{}
	for _, match := range mul.FindAllStringSubmatch(sdat, -1) {
		if len(match) < 2 {
			continue
		}
		d1, err := strconv.Atoi(match[1])
		check(err)
		d2, err := strconv.Atoi(match[2])
		check(err)
		res = append(res, Pair{d1, d2})
	}
	return res
}

func part1(data []Pair) (r int) {
	for _, p := range data {
		r += p.x * p.y
	}
	return
}

func project(xs [][]int, k int) []int {
	res := make([]int, len(xs))
	for i := 0; i < len(xs); i += 1 {
		res[i] = xs[i][k]
	}
	return res
}

func bisect(xs []int, target int) int {
	l, r, mid := 0, len(xs), len(xs)
	for l < r {
		mid = (r-l)/2 + l
		if xs[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}

func parsePart2(path string) {
	dat, err := os.ReadFile(path)
	check(err)
	sdat := string(dat)
	mul := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	do := regexp.MustCompile(`do\(\)`)
	dont := regexp.MustCompile(`don't\(\)`)

	res := 0
	matches := mul.FindAllStringSubmatch(sdat, -1)
	idxs := mul.FindAllStringIndex(sdat, -1)

	dos := project(do.FindAllStringIndex(sdat, -1), 0)
	donts := project(dont.FindAllStringIndex(sdat, -1), 0)

	fmt.Println(dos)
	fmt.Println(donts)

	for i := range len(matches) {
		idx := idxs[i]
		start := idx[0]

		// rather we could have two sweeping indices
		d0 := bisect(dos, start)
		d1 := bisect(donts, start)

		// ugly logic
		if d0 == 0 && d1 != 0 && donts[d1-1] < start {
			continue
		} else if d1 != 0 && donts[d1-1] < start && !(donts[d1-1] < dos[d0-1] && dos[d0-1] < start) {
			continue
		}

		match := matches[i]
		if len(match) < 2 {
			continue
		}
		d1, err := strconv.Atoi(match[1])
		check(err)
		d2, err := strconv.Atoi(match[2])
		check(err)
		res += d1 * d2
	}

	fmt.Println(res)
}

func main() {
	// fmt.Println(bisect([]int{0, 2, 5, 7, 9}, 3))
	// fmt.Println(bisect([]int{0, 2, 5, 7, 9}, 0))
	// fmt.Println(bisect([]int{0, 2, 5, 7, 9}, -1))
	// fmt.Println(bisect([]int{0, 2, 5, 7, 9}, 10))
	path := os.Args[1]
	data := parse(path)
	res1 := part1(data)
	fmt.Println(res1)

	parsePart2(path)
}
