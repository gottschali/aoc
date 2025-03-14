package main

import (
	"fmt"
	"math"
	"slices"
)

type Operation int

const (
	Adv Operation = iota
	Bxl
	Bst
	Jnz
	Bxc
	Out
	Bdv
	Cdv
)

type State struct {
	RegA, RegB, RegC int
	IP               int
}

type Emulator struct {
	s      State
	memory []int
	output []int
}

func NewEmulator(a, b, c int, memory []int) *Emulator {
	return &Emulator{
		s:      State{a, b, c, 0},
		memory: memory,
		output: make([]int, 0),
	}
}

func (e *Emulator) combo(operand int) int {
	switch operand {
	default:
		panic("unknown combo operand")
	case 0, 1, 2, 3:
		return operand
	case 4:
		return e.s.RegA
	case 5:
		return e.s.RegB
	case 6:
		return e.s.RegC
	}

}

func dv(x, y int) int {
	return int(math.Trunc(float64(x) / float64(int(1)<<y)))
}

func (e *Emulator) calculate(operation Operation, operand int) {
	co := e.combo(operand)
	switch operation {
	default:
		panic("unknown operation")
	case Adv:
		e.s.RegA = dv(e.s.RegA, co)
	case Bxl:
		e.s.RegB = e.s.RegB ^ operand
	case Bst:
		e.s.RegB = co % 8
	case Jnz:
		if e.s.RegA != 0 {
			e.s.IP = operand
		} else {
			e.s.IP += 2
		}
	case Bxc:
		e.s.RegB = e.s.RegB ^ e.s.RegC
	case Out:
		// if len(e.output) > 0 {
		// 	fmt.Printf(",")
		// }
		// fmt.Printf("%d", co%8)
		e.output = append(e.output, co%8)
	case Bdv:
		e.s.RegB = dv(e.s.RegA, co)
	case Cdv:
		e.s.RegC = dv(e.s.RegA, co)
	}
}

func (e *Emulator) Step() (halt bool) {
	if e.s.IP >= len(e.memory) {
		// fmt.Print("\n")
		return true
	}
	operation := Operation(e.memory[e.s.IP])
	operand := e.memory[e.s.IP+1]
	// fmt.Println(operation, operand)
	e.calculate(operation, operand)

	if operation != Jnz {
		e.s.IP += 2
	}
	return false
}

func (s State) String() string {
	return fmt.Sprintf("A=%d, B=%d, C=%d, IP=%d",
		s.RegA, s.RegB, s.RegC, s.IP)
}

// Runs the program until it halts
// or if n is postitive stops after n steps
func (e *Emulator) Run(n int) {
	// fmt.Printf("%s\n", e.s)
	for i := 0; n < 0 || i < n; i += 1 {
		if halt := e.Step(); halt {
			return
		}
		// fmt.Printf("%s\n", e.s)
	}
}

// returns the smallest non-negative value
// for the initial value of register A,
// such that the program outputs a copy of itself
func (e *Emulator) SearchQuine() int {
	initialState := e.s
	for A := 0; ; A += 1 {
		if A%100_000 == 0 {
			fmt.Println("A: ", A, " out: ", e.output)
		}
		// reset the state
		e.s = initialState
		e.output = []int{}
		e.s.RegA = A
		// unoptimized
		for i := 0; ; i += 1 {
			if halt := e.Step(); halt {
				break
			}
			// OPTIMIZATION 1: quit if output too long
			if len(e.output) >= len(e.memory) {
				break
			}

			// OPTIMIZATION 2: quit after first difference is found
			if len(e.output) > 0 && e.output[len(e.output)-1] != e.memory[len(e.output)-1] {
				break
			}
		}

		if slices.Equal(e.memory, e.output) {
			return A
		}
	}

}
