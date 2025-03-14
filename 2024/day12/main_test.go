package main

import (
	"testing"
)

func Test3(t *testing.T) {
	s := Parse("test3")
	res2 := s.Part2()
	exp := 368

	if res2 != exp {
		t.Errorf("expectec %d, got %d", exp, res2)
	}
}

func TestE(t *testing.T) {
	s := Parse("testE")
	res2 := s.Part2()
	exp := 236
	if res2 != exp {
		t.Errorf("expected %d, got %d", exp, res2)
	}
}

func Test2_2(t *testing.T) {
	s := Parse("test2")
	res2 := s.Part2()
	exp := 1206
	if res2 != exp {
		t.Errorf("expected %d, got %d", exp, res2)
	}
}
