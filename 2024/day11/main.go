package main

import (
	"fmt"
	"strconv"
)

func blink(stones []int) (res []int) {
	for _, s := range stones {
		st := strconv.Itoa(s)
		if s == 0 {
			res = append(res, 1)
		} else if len(st)%2 == 0 {
			s1 := st[:len(st)/2]
			s2 := st[len(st)/2:]
			n1, _ := strconv.Atoi(s1)
			n2, _ := strconv.Atoi(s2)
			res = append(res, n1)
			res = append(res, n2)
		} else {
			res = append(res, s*2024)
		}
	}
	return
}

func Part1(stones []int) int {
	for range 25 {
		stones = blink(stones)
	}
	return len(stones)
}

func blinkMap(stones map[int]int) (res map[int]int) {
	res = make(map[int]int)
	for s, c := range stones {
		st := strconv.Itoa(s)
		if s == 0 {
			res[1] += c
		} else if len(st)%2 == 0 {
			s1 := st[:len(st)/2]
			s2 := st[len(st)/2:]
			n1, _ := strconv.Atoi(s1)
			n2, _ := strconv.Atoi(s2)
			res[n1] += c
			res[n2] += c
		} else {
			res[s*2024] += c
		}
	}
	return
}

func sumMap(stones map[int]int) (res int) {
	for _, v := range stones {
		res += v
	}
	return
}

func Part2(stones []int) int {
	m := make(map[int]int)
	for _, s := range stones {
		m[s] += 1
	}
	for i := range 75 {
		m = blinkMap(m)
		// fmt.Println(i, len(stones))
		fmt.Println(i, sumMap(m))
	}
	return sumMap(m)
}

func main() {
	stones := []int{125, 17}
	res1 := Part1(stones)
	fmt.Println(res1)
	input := []int{27, 10647, 103, 9, 0, 5524, 4594227, 902936}
	res1 = Part1(input)
	fmt.Println(res1)
	res2 := Part2(input)
	fmt.Println(res2)
}
