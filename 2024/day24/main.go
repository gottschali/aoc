package main

import (
	helper "aoc2024"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func parseInputs(inp []byte) map[string]int {
	m := make(map[string]int)
	for _, line := range bytes.Split(inp, []byte{'\n'}) {
		parts := bytes.Split(line, []byte(": "))
		name := string(parts[0])
		if parts[1][0] == '1' {
			m[name] = 1
		} else {
			m[name] = 0
		}
	}
	return m
}

func parseGates(inp []byte) (gates []*Gate) {
	for _, line := range bytes.Split(inp, []byte{'\n'}) {
		parts := bytes.Split(line, []byte(" "))
		i1 := string(parts[0])
		i2 := string(parts[2])
		o := string(parts[4])
		var gt Operation
		switch parts[1][0] {
		default:
			panic("unkown operation")
		case 'X':
			gt = XOR
		case 'O':
			gt = OR
		case 'A':
			gt = AND
		}

		gates = append(gates, &Gate{i1, i2, o, false, gt})

	}
	return
}

type Operation int

const (
	XOR Operation = iota
	OR
	AND
)

type Gate struct {
	in1, in2, out string
	activated     bool
	op            Operation
}

type Solver struct {
	inputs  map[string]int
	gates   []*Gate
	results []int
}

func (s *Solver) part1() (res int) {
	// this is way too long...
	s.results = make([]int, len(s.inputs))

	for ok := s.iterate(); ok; ok = s.iterate() {
		fmt.Println(len(s.inputs))
	}

	// get the binary number together
	for n, v := range s.inputs {
		if v == 0 {
			continue
		}
		if !strings.HasPrefix(n, "z") {
			continue
		}
		position, _ := strconv.Atoi(strings.TrimPrefix(n, "z"))
		fmt.Println(n, v, position)
		res += 1 << position
	}
	return
}

func (s *Solver) iterate() (updated bool) {
	for _, g := range s.gates {
		// optimization: remove them between iterations?
		if g.activated {
			continue
		}
		v1, ready1 := s.inputs[g.in1]
		v2, ready2 := s.inputs[g.in2]
		if !ready1 || !ready2 {
			continue
		}
		var v3 int
		switch g.op {
		default:
			panic("unkonwn op")
		case XOR:
			v3 = v1 ^ v2
		case AND:
			v3 = v1 * v2
		case OR:
			v3 = v1 + v2
			if v3 == 2 {
				v3 = 1
			}
		}
		g.activated = true
		s.inputs[g.out] = v3
		updated = true
	}
	return

}

func (s *Solver) Parse(path string) {
	sections := helper.ReadFileSections(path)
	s.inputs = parseInputs(sections[0])
	s.gates = parseGates(sections[1])
}

func main() {
	s := new(Solver)
	s.Parse("input")
	fmt.Println(s.inputs)
	fmt.Println(s.gates)

	res1 := s.part1()
	fmt.Println("res1: ", res1)

}
