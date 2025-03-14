package main

import (
	"fmt"
	"testing"
)

func Test123(t *testing.T) {
	s1 := 123
	next := secret(s1)
	expected := []int{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}
	for i, exp := range expected {
		if exp != next {
			t.Errorf("secret number %d, got %d, want %d", i, next, exp)
		}
		next = secret(next)
	}
}

func TestSearch(t *testing.T) {
	secret := 123
	ch := changes(secret, 10)
	fmt.Println(ch)
	price := searchSequence(ch, []int{-1, -1, 0, 2})
	sell := 6
	if price != sell {
		t.Errorf("prize is %d, want %d", price, sell)
	}
}

func Test2000(t *testing.T) {
	ts := []struct{ start, steps, want int }{
		{
			start: 1,
			steps: 2000,
			want:  8685429,
		},
		{
			start: 10,
			steps: 2000,
			want:  4700978,
		},
		{
			start: 100,
			steps: 2000,
			want:  15273692,
		},
		{
			start: 2024,
			steps: 2000,
			want:  8667524,
		},
	}
	for _, tt := range ts {
		t.Run(fmt.Sprintf("test%d", tt.start), func(t *testing.T) {
			got := simulate(tt.start, tt.steps)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func Test2(t *testing.T) {
	inp := []int{1, 2, 3, 2024}
	got := part2(inp)
	exp := 23

	if got != exp {
		t.Errorf("got %d want %d", got, exp)
	}
}
