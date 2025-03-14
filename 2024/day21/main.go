package main

import "slices"

type numpad int

const (
	NUM0 numpad = iota
	NUM1
	NUM2
	NUM3
	NUM4
	NUM5
	NUM6
	NUM7
	NUM8
	NUM9
	NUMA
)

type keypad int

const (
	A keypad = iota
	UP
	RIGHT
	DOWN
	LEFT
)

func (a keypad) String() string {
	switch a {
	default:
		panic("unexpected key")
	case A:
		return "A"
	case UP:
		return "^"
	case RIGHT:
		return ">"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	}
}
func (n numpad) String() string {
	if n == NUMA {
		return "A"
	} else if 0 <= n && n < 10 {
		return string(n)
	} else {
		panic("unepxected numpad")
	}
}

type state struct {
	k1, k2 keypad
	n      numpad
}
type pair struct {
	x, y state
}

var keyNeighbors = map[keypad][]keypad{
	A:     []keypad{UP, RIGHT},
	UP:    []keypad{A, DOWN},
	RIGHT: []keypad{A, DOWN},
	DOWN:  []keypad{UP, RIGHT, LEFT},
	LEFT:  []keypad{DOWN},
}
var numNeighbors = map[numpad][]numpad{
	NUM0: []numpad{NUMA, NUM2},
	NUM1: []numpad{NUM2, NUM4},
	NUM2: []numpad{NUM0, NUM1, NUM3, NUM5},
	NUM3: []numpad{NUM2, NUM6, NUMA},
	NUM4: []numpad{NUM1, NUM5, NUM7},
	NUM5: []numpad{NUM2, NUM4, NUM6, NUM8},
	NUM6: []numpad{NUM3, NUM5, NUM9},
	NUM7: []numpad{NUM4, NUM8},
	NUM8: []numpad{NUM5, NUM7, NUM9},
	NUM9: []numpad{NUM6, NUM8},
	NUMA: []numpad{NUM0, NUM3},
}

// func floyd() {
// 	dists := make(map[pair]int)
// 	// initialize distances
// 	for k1 := A; k1 <= LEFT; k1 += 1 {
// 		for k2 := A; k2 <= LEFT; k2 += 1 {
// 			for n := NUM0; n <= NUMA; n += 1 {
// 				s := state{k1, k2, n}
// 				dists[pair{s, s}] = 0
// 			}
// 		}
// 	}

// 	for k1 := A; k1 <= LEFT; k1 += 1 {
// 		for k2 := A; k2 <= LEFT; k2 += 1 {

// 		}
// 	}
// }
