package helper

import (
	"bytes"
	"log"
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFileLines(path string) [][]byte {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read input file: ", err)
	}
	dat = bytes.TrimSpace(dat)
	return bytes.Split(dat, []byte{'\n'})
}

func ReadFileSections(path string) [][]byte {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read input file: ", err)
	}
	dat = bytes.TrimSpace(dat)
	return bytes.Split(dat, []byte("\n\n"))
}

func ReadFileLinesString(path string) []string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read input file: ", err)
	}
	dat = bytes.TrimSpace(dat)
	return strings.Split(string(dat), "\n")
}

func ReadFile(path string) []byte {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read input file: ", err)
	}
	dat = bytes.TrimSpace(dat)
	return dat
}

type Point struct {
	X, Y int
}

func (p *Point) Add(o *Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

func (p *Point) Scale(s int) Point {
	return Point{p.X * s, p.Y * s}
}
