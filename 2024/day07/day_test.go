package main

import (
	"fmt"
	"slices"
	"testing"
)

func (c *Calibration) Equal(o Calibration) bool {
	if c.value != o.value {
		return false
	}
	return slices.Equal(c.numbers, o.numbers)
}

func TestParsingLines(t *testing.T) {
	var exp, got Calibration

	got = ParseLine("190: 10 19")
	exp = Calibration{190, []int{10, 19}}
	if !exp.Equal(got) {
		t.Errorf("exp %d, got %d", exp, got)
	}

	got = ParseLine("21037: 9 7 18 13")
	exp = Calibration{21037, []int{9, 7, 18, 13}}
	if !exp.Equal(got) {
		t.Errorf("exp %d, got %d", exp, got)
	}

	got = ParseLine("0: 1")
	exp = Calibration{0, []int{1}}
	if !exp.Equal(got) {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func TestPart1(t *testing.T) {
	input := Parse("test")
	got := Part1(input)
	exp := 3749
	if exp != got {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func TestPart2(t *testing.T) {
	input := Parse("test")
	got := Part2(input)
	exp := 11387
	if exp != got {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func TestInput1(t *testing.T) {
	input := Parse("input")
	got := Part1(input)
	exp := 1298300076754
	if exp != got {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := Parse("input")
	for range b.N {
		Part1(input)
	}
}

func TestInput2(t *testing.T) {
	input := Parse("input")
	got := Part2(input)
	exp := 248427118972289
	if exp != got {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func TestConcat(t *testing.T) {
	ts := []struct {
		a, b int
		want int
	}{
		{0, 0, 0},
		{1, 0, 10},
		{0, 1, 1},
		{9, 9, 99},
		{42, 0, 420},
		{69, 69, 6969},
	}
	for _, tt := range ts {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Concat(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
