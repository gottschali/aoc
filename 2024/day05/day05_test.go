package main

import (
	"aoc2024"
	"bytes"
	"slices"
	"testing"
)

func TestPart1(t *testing.T) {
	dat := bytes.Split(helper.ReadFile("test"), []byte("\n\n"))
	rules := ParseRules(dat[0])
	manuals := ParseManuals(dat[1])
	expected := 143
	got := Part1(rules, manuals)
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestPart2(t *testing.T) {
	dat := bytes.Split(helper.ReadFile("test"), []byte("\n\n"))
	rules := ParseRules(dat[0])
	manuals := ParseManuals(dat[1])
	expected := 123
	got := Part2(rules, manuals)
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}

}

func TestPart2Reorder(t *testing.T) {
	dat := bytes.Split(helper.ReadFile("test"), []byte("\n\n"))
	rules := ParseRules(dat[0])
	before := makeBefore(rules)
	wrongManuals := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	correctManuals := [][]int{
		{97, 75, 47, 61, 53},
		{61, 29, 13},
		{97, 75, 47, 29, 13},
	}
	for i, wrong := range wrongManuals {
		got := Reorder(wrong, before)
		want := correctManuals[i]
		if slices.Compare(got, want) != 0 {
			t.Errorf("want %d, got %d", want, got)
		}
	}
}

func Part1Input(t *testing.T) {
	dat := bytes.Split(helper.ReadFile("input"), []byte("\n\n"))
	rules := ParseRules(dat[0])
	manuals := ParseManuals(dat[1])
	expected := 4135
	got := Part1(rules, manuals)
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
