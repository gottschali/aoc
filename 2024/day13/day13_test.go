package day13

import (
	"testing"
)

func TestButtonParse(t *testing.T) {
	s := []byte("Button A: X+94, Y+34")
	want := Point{94, 34}
	got := ParseLine(s)
	if want != got {
		t.Error("Failed to parse button line")
	}
}

func TestPrizeParse(t *testing.T) {
	s := []byte("Prize: X=8400, Y=5400")
	want := Point{8400, 5400}
	got := ParseLine(s)
	if want != got {
		t.Error("Failed to parse button line")
	}
}

type TestRiddle struct {
	name     string
	riddle   Riddle
	expected int
}

func TestPart1Examples(t *testing.T) {
	tests := []TestRiddle{
		{"A", Riddle{Point{94, 34}, Point{22, 67}, Point{8400, 5400}}, 280},
		{"B", Riddle{Point{26, 66}, Point{67, 21}, Point{12748, 12176}}, 0},
		{"C", Riddle{Point{17, 86}, Point{84, 37}, Point{7870, 6450}}, 200},
		{"D", Riddle{Point{69, 23}, Point{27, 71}, Point{18641, 10279}}, 0},
		{"rounding", Riddle{Point{16, 58}, Point{83, 72}, Point{2007, 4300}}, 58*3 + 13},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Solve(tc.riddle, 100)
			if got != tc.expected {
				t.Errorf("Solve(%d) = %d; want %d", tc.riddle, got, tc.expected)
			}
		})
	}
}
