package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func parse(path string) (xs, ys []int) {
	dat, err := os.ReadFile(path)
	check(err)
	data := string(dat)
	for _, line := range strings.Split(data, "\n") {
		nums := strings.Split(line, "   ")
		if len(nums) < 2 {
			continue
		}
		x, err := strconv.Atoi(nums[0])
		check(err)
		y, err := strconv.Atoi(nums[1])
		check(err)
		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func part1(xs, ys []int) int {
	slices.Sort(xs)
	slices.Sort(ys)
	sum := 0
	for i := range xs {
		sum += Abs(xs[i] - ys[i])
	}
	return sum
}

func count(xs []int) (counter map[int]int) {
	counter = make(map[int]int)
	for _, x := range xs {
		counter[x]++
	}
	return
}

func part2(xs, ys []int) int {
	c := count(ys)
	res := 0
	for _, x := range xs {
		res += x * c[x]
	}
	return res
}

func main() {
	path := os.Args[1]
	xs, ys := parse(path)
	res1 := part1(xs, ys)
	fmt.Println(res1)
	res2 := part2(xs, ys)
	fmt.Println(res2)
}
