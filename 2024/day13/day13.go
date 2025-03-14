package day13

import (
	"bytes"
	_ "fmt"
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// var buttonRex = regexp.MustCompile(`Button .: X+([0-9]+), Y+([0-9]+)`)
// var prizeRex = regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)
var numRex = regexp.MustCompile(`[0-9]+`)

func ParseLine(line []byte) Point {
	matches := numRex.FindAll(line, -1)
	x, err := strconv.Atoi(string(matches[0]))
	check(err)
	y, err := strconv.Atoi(string(matches[1]))
	check(err)
	return Point{x, y}
}

type Riddle struct {
	ButtonA, ButtonB, Prize Point
}

func Parse(dat []byte) (rs []Riddle) {
	for _, chunk := range bytes.Split(dat, []byte{'\n', '\n'}) {
		lines := bytes.Split(chunk, []byte{'\n'})
		rs = append(rs, Riddle{
			ParseLine(lines[0]),
			ParseLine(lines[1]),
			ParseLine(lines[2]),
		})
	}
	return
}

const (
	costA = 3
	costB = 1
)

type mat22 struct {
	a, b Point
}

func Solve(r Riddle, limit int) int {
	//
	a := float64(r.ButtonA.x)
	c := float64(r.ButtonA.y)
	b := float64(r.ButtonB.x)
	d := float64(r.ButtonB.y)

	px := float64(r.Prize.x)
	py := float64(r.Prize.y)

	// easier: solve by substitution
	// A bit over-engineered:
	//
	// In standard euclidean basis
	//         P
	//    -------
	//   /  /  /
	//  /  /  /
	// /  /  /
	// ------
	// In basis {ButtonA, ButtonB}
	//       P'
	// -------
	// |     |
	// -------
	//
	// Transform by inverse(ButtonA|ButtonB)
	//
	// if the coordinates of P' are whole numbers
	// they give the number of button presses
	// The costs only scales the solution and have no direct influence

	// computes the inverse of the 2x2 matrix
	// X := ( p1 | p2 ) = ( a b ; c d)
	// det(X) = a * d - b c
	// inv(X) = 1/det(x) * ( d  -b ; -c a )
	s := 1.0 / (a*d - b*c)
	Na := s * (px*d + py*(-b))
	Nb := s * (px*(-c) + py*a)

	if AlmostIntegral(Na) && AlmostIntegral(Nb) {
		iNa := int(math.Round(Na))
		iNb := int(math.Round(Nb))
		if limit < 0 || (iNa <= limit && iNb <= limit) {
			return costA*iNa + costB*iNb
		}
	}
	return 0
}

func AlmostIntegral(f float64) bool {
	return math.Abs(f-float64(math.Round(f))) < 1e-3
}

func Part1(riddles []Riddle) int {
	res := 0
	for _, r := range riddles {
		res += Solve(r, 100)
	}
	return res
}

const shift = 10000000000000

func Part2(riddles []Riddle) int {
	res := 0
	for _, r := range riddles {
		r.Prize.x += shift
		r.Prize.y += shift
		res += Solve(r, -1)
	}
	return res
}
