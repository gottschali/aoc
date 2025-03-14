package main

import (
	"aoc2024"
	"bytes"
	"fmt"
	"strconv"
)

type Rule struct {
	first, second int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseRules(dat []byte) (rs []Rule) {
	dat = bytes.TrimSpace(dat)
	for _, line := range bytes.Split(dat, []byte{'\n'}) {
		nums := bytes.Split(line, []byte("|"))
		f, err := strconv.Atoi(string(nums[0]))
		check(err)
		s, err := strconv.Atoi(string(nums[1]))
		check(err)
		rs = append(rs, Rule{f, s})
	}
	return rs
}

func ParseManuals(dat []byte) (ms [][]int) {
	dat = bytes.TrimSpace(dat)
	for _, line := range bytes.Split(dat, []byte{'\n'}) {
		nums := bytes.Split(line, []byte(","))
		manual := make([]int, 0, len(nums))
		for _, n := range nums {
			page, err := strconv.Atoi(string(n))
			check(err)
			manual = append(manual, page)

		}
		ms = append(ms, manual)
	}
	return ms
}

func checkOrder(manual []int, before map[Rule]bool) bool {
	for i, m := range manual {
		for j := i + 1; j < len(manual); j += 1 {
			if _, ok := before[Rule{m, manual[j]}]; ok {
				return false
			}
		}
	}
	return true
}

func makeBefore(rules []Rule) map[Rule]bool {
	before := make(map[Rule]bool)
	for _, rule := range rules {
		before[Rule{rule.second, rule.first}] = true
	}
	return before
}

func Part1(rules []Rule, manuals [][]int) (res1 int) {
	before := makeBefore(rules)
	for _, manual := range manuals {
		if checkOrder(manual, before) {
			res1 += manual[len(manual)/2]
		}
	}
	return res1
}

func Part2(rules []Rule, manuals [][]int) (res2 int) {
	before := makeBefore(rules)
	for _, manual := range manuals {
		if !checkOrder(manual, before) {
			manual = Reorder(manual, before)
			res2 += manual[len(manual)/2]
		}
	}
	return res2
}

func Reorder(manual []int, before map[Rule]bool) (res []int) {
	res = append(res, manual...)
	for i := range manual {
		for j := i + 1; j < len(manual); j += 1 {
			if _, ok := before[Rule{res[i], res[j]}]; ok {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}

func driver(path string) {
	dat := bytes.Split(helper.ReadFile(path), []byte("\n\n"))
	rules := ParseRules(dat[0])
	manuals := ParseManuals(dat[1])
	// fmt.Println(rules, manuals)
	res1 := Part1(rules, manuals)
	fmt.Println("result part1:\n", res1)

	res2 := Part2(rules, manuals)
	fmt.Println("result part2:\n", res2)
}

func main() {
	driver("test")
	driver("input")
}
