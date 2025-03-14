package main

import (
	helper "aoc2024"
	"bytes"
	"fmt"
	"strconv"
)

func parseRegister(line []byte) int {
	n, _ := strconv.Atoi(string(bytes.Split(line, []byte(": "))[1]))
	return n
}

func parseProgram(line []byte) (res []int) {
	memory := bytes.Split(line, []byte(": "))[1]
	for _, b := range bytes.Split(memory, []byte(",")) {
		n, _ := strconv.Atoi(string(b))
		res = append(res, n)
	}
	return
}

func Parse(path string) *Emulator {
	bs := helper.ReadFileLines(path)

	return NewEmulator(
		parseRegister(bs[0]),
		parseRegister(bs[1]),
		parseRegister(bs[2]),
		parseProgram(bs[4]),
	)
}

func main() {
	e := Parse("input")
	fmt.Println(e)

	// e.Run(-1)
	A := e.SearchQuine()
	fmt.Println("A: ", A)
}
