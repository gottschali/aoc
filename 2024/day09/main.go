package main

import (
	"aoc2024"
	"fmt"
	"strconv"
)

const empty int = -1

func Parse(path string) (ns []int) {
	for _, b := range helper.ReadFile(path) {
		n, _ := strconv.Atoi(string(b))
		ns = append(ns, n)
	}
	return
}

func checksum(ns []int) (c int) {
	for i, n := range ns {
		if n != empty {
			c += i * n
		}
	}
	return
}

func expand(compat []int) (res []int) {
	for i, n := range compat {
		for range n {
			if i%2 == 0 {
				res = append(res, i/2)
			} else {

				res = append(res, empty)
			}
		}
	}
	return
}

func Part1(ns []int) int {
	// fmt.Println(ns)
	ns = expand(ns)
	// fmt.Println(ns)
	l := 0
	h := len(ns) - 1
	for l < h {
		if ns[h] == empty {
			h -= 1
			continue
		}
		if ns[l] == empty {
			ns[l], ns[h] = ns[h], ns[l]
			h -= 1
		}
		l += 1
		// fmt.Println(ns)
	}
	return checksum(ns)
}

type block struct {
	start, length int
}

func blockInfo(compat []int) (files []block, gaps []block) {
	pos := 0
	for i, n := range compat {
		if i%2 == 0 {
			files = append(files, block{start: pos, length: n})
		} else {
			gaps = append(gaps, block{start: pos, length: n})
		}
		pos += n
	}
	return
}

func virtualChecksum(files []block) (c int) {
	for fileID, f := range files {
		for l := 0; l < f.length; l += 1 {
			c += (f.start + l) * fileID
		}
	}
	return
}

func printFiles(files []block, n int) string {
	buffer := make([]byte, n)
	for i := range buffer {
		buffer[i] = '.'
	}
	for i, f := range files {
		for l := range f.length {
			buffer[f.start+l] = byte(i%10) + '0'
		}
	}
	return string(buffer)
}

func Part2(ns []int) int {
	// fmt.Println(ns)
	files, gaps := blockInfo(ns)
	// fmt.Println(files)
	last := files[len(files)-1]
	n := last.start + last.length
	fmt.Println(printFiles(files, n))
	// fmt.Println(gaps)
	for fileID := len(files) - 1; fileID >= 0; fileID -= 1 {
		f := files[fileID]
		for gapID, gap := range gaps {
			// block must be to the left of a file
			if gap.start >= f.start {
				break
			}
			if gap.length >= f.length {
				files[fileID].start = gap.start
				gaps[gapID] = block{start: gap.start + f.length, length: gap.length - f.length}
				break
			}
		}
	}
	fmt.Println(printFiles(files, n))
	return virtualChecksum(files)
}

func main() {
	data := Parse("test")
	res1 := Part1(data)
	fmt.Println("part1: ", res1)
	res2 := Part2(data)
	fmt.Println("part2: ", res2)
	data = Parse("input")
	res1 = Part1(data)
	fmt.Println("part1: ", res1)
	res2 = Part2(data)
	fmt.Println("part2: ", res2)
}
