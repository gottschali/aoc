package main

import (
	"aoc2024"
	"testing"
)

func TestPart1(t *testing.T) {
	grid := helper.ReadFileLines("test")
	got := Part1(grid)
	expected := 41
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestInput1(t *testing.T) {
	grid := helper.ReadFileLines("input")
	got := Part1(grid)
	expected := 4663
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
