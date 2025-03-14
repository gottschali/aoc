package main

import (
	helper "aoc2024"
	"fmt"
	"slices"
	"strconv"
)

const modulus = 16777216

func secret(s int) (res int) {
	s = (s ^ (s << 6)) % modulus
	s = (s ^ (s >> 5)) % modulus
	s = (s ^ (s << 11)) % modulus
	return s
}

func simulate(initialSecret int, steps int) (finalSecret int) {
	s := initialSecret
	for range steps {
		s = secret(s)
	}
	return s
}

func changes(initialSecret int, steps int) (change []int) {
	s := initialSecret
	var prev int
	change = append(change, s%10)
	for range steps {
		prev = s
		s = secret(s)
		change = append(change, (s%10)-(prev%10))
	}
	return
}

func searchSequence(changes []int, seq []int) (prize int) {
	for i, c := range changes {
		if i >= len(seq) && slices.Equal(changes[i-len(seq):i], seq) {
			return prize
		}
		prize += c
	}
	return 0
}

func parse(path string) (nums []int) {
	for _, s := range helper.ReadFileLinesString(path) {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	return
}

func Part1(nums []int) (res int) {
	for _, n := range nums {
		res += simulate(n, 2000)
	}
	return
}

type sequence [4]int

func shift(s sequence, x int) sequence {
	return [4]int{s[1], s[2], s[3], x}
}

func solve2(initialSecret int, steps int, prizes map[sequence]int) {
	seen := make(map[sequence]bool)
	s := initialSecret
	var prev int
	seq := [4]int{-10, -10, -10, s % 10}
	prize := s % 10
	for range steps {
		prev = s
		s = secret(s)
		change := (s % 10) - (prev % 10)
		seq = shift(seq, change)
		prize += change
		if !seen[seq] {
			prizes[seq] += prize
		}
		seen[seq] = true
		// fmt.Println(seq, prize)
	}
	return
}

func part2(nums []int) int {
	prizes := make(map[sequence]int)
	for _, n := range nums {
		solve2(n, 2_000, prizes)
	}
	fmt.Println("len(prizes)= ", len(prizes))

	maxPrize := -1
	for seq, prize := range prizes {
		if prize > maxPrize {
			maxPrize = prize
			fmt.Println("prize", maxPrize, " for seq, ", seq)
		}
	}
	return maxPrize
}

func main() {
	nums := parse("input")
	res1 := Part1(nums)
	fmt.Println("res1: ", res1)
	res2 := part2(nums)
	fmt.Println("res2: ", res2)
}
