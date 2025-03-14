package main

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}

func TestSimulations(t *testing.T) {
	tests := []struct{ path, finalPath string }{
		{"test_small", "test_small_final"},
		{"test", "test_final"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			input := Parse(tt.path)
			exp := Parse(tt.finalPath)
			Simulate(input)
			if !exp.grid.Equals(input.grid) {
				t.Errorf("exp %v, got %v", exp.grid, input.grid)
			}
		})
	}
}

func TestUpscale(t *testing.T) {
	tests := []struct{ path, finalPath string }{
		{"upscale", "upscale_final"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			input := Parse(tt.path)
			exp := Parse(tt.finalPath)
			got := input.grid.Upscale()
			if !exp.grid.Equals(*got) {
				t.Errorf("exp \n%v, got \n%v for input\n %v", exp.grid, got, input.grid)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := Parse("test")
	got := Part1(input)
	exp := 10092
	if exp != got {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func TestPart2(t *testing.T) {
	t.Skip()
	input := Parse("test")
	got := Part2(input)
	exp := 11387
	if exp != got {
		t.Errorf("exp %d, got %d", exp, got)
	}
}
