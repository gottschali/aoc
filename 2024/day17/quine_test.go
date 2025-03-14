package main

import (
	"testing"
)

func BechmarkQuine(b *testing.B) {
	e := Parse("quine")
	for i := 0; i < b.N; i += 1 {
		e.SearchQuine()
	}
}
