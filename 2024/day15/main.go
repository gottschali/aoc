package main

import (
	"aoc2024"
	"bytes"
	"fmt"
	"log"
	"os"
	_ "os"
	"slices"
	"time"
)

type Grid struct {
	width, height int
	data          [][]byte
}

func (g Grid) String() (res string) {
	return string(bytes.Join(g.data, []byte{'\n'}))
}

func (g *Grid) FindFirst(target byte) (ok bool, x int, y int) {
	for y, line := range g.data {
		x := slices.Index(line, target)
		if x != -1 {
			return true, x, y
		}
	}
	return false, -1, -1
}

func (g *Grid) Upscale() (res *Grid) {
	res = &Grid{
		width:  2 * g.width,
		height: g.height,
		data:   make([][]byte, g.height),
	}
	for y, line := range g.data {
		res.data[y] = make([]byte, 2*len(line))
		for x, box := range line {
			res.data[y][2*x] = g.data[y][x]
			res.data[y][2*x+1] = g.data[y][x]
			if box == ROBOT {
				res.data[y][2*x+1] = EMPTY
			} else if box == BOX {
				res.data[y][2*x] = BOX_LEFT
				res.data[y][2*x+1] = BOX_RIGHT
			}
		}
	}
	return res
}

func (g *Grid) Update(x, y int, move Direction) (xOut, yOut int) {
	dx, dy := DirToDelta(move)
	xNext, yNext := x+dx, y+dy

	blockNext, _ := g.At(xNext, yNext)

	switch blockNext {
	case EMPTY:
		logger.Println("walking")
		xOut, yOut = xNext, yNext
	case BOX:
		logger.Print("found a box")
		// Push boxes
		//   find the first non-Box square in the movements direction
		//   @OOOOO. |-> @.OOOOO
		//   @OO# |-> @OO#
		xBox, yBox := xNext, yNext
		for g.data[yBox][xBox] == BOX {
			xBox += dx
			yBox += dy
		}
		if g.data[yBox][xBox] == EMPTY {
			g.data[yNext][xNext], g.data[yBox][xBox] = g.data[yBox][xBox], g.data[yNext][xNext]
			xOut, yOut = xNext, yNext
			logger.Print("   push it")
		} else {
			xOut, yOut = x, y
			logger.Print("   cannot push")
		}
	case OBSTACLE:
		logger.Println("run into obstacle")
		xOut, yOut = x, y
	default:
		logger.Panic("unexpected block")
	}

	// swap the robots position
	g.data[yOut][xOut], g.data[y][x] = g.data[y][x], g.data[yOut][xOut]
	return
}

func (g *Grid) Push(x, y, dx, dy int) (int, int) {
	xn, yn := x+dx, y+dy
	// we know that the next is either [ or ]

	switch {
	case dy == 0 && dx == 1 && g.data[yn][xn] == BOX_LEFT:
		// ?@[
		g.pushHorizontal(xn, yn, dx)
	case dy == 0 && dx == -1 && g.data[yn][xn] == BOX_RIGHT:
		// ]@
		g.pushHorizontal(xn-1, yn, dx)
	case dx == 0 && g.data[yn][xn] == BOX_LEFT:
		g.pushVertical(xn, yn, dy)
	case dx == 0 && g.data[yn][xn] == BOX_RIGHT:
		g.pushVertical(xn-1, yn, dy)
	default:
		log.Panic("unhandled push case")
	}
	xNext, yNext := x+dx, y+dy
	log.Printf("finalize push (%d, %d) %d, %d", dx, dy, x, y)
	if g.data[yNext][xNext] == EMPTY {
		// g.data[yNext][xNext] = ROBOT
		// g.data[y][x] = EMPTY
		return xNext, yNext
	} else {
		return x, y
	}
}

// (x, y) is the position of the left half of the box [
// the box is pushed in direction (dx, 0)
func (g *Grid) pushHorizontal(x, y, dx int) {
	log.Printf("horizontal (%d) %d, %d", dx, x, y)
	if g.data[y][x] != BOX_LEFT {
		panic("box invariant violated")
	}
	if dx == 1 {
		// >[]
		block := g.data[y][x+2]
		switch block {
		case EMPTY, OBSTACLE:
			break
		case BOX_LEFT:
			g.pushHorizontal(x+2, y, dx)
		default:
			log.Panic("unhandled push horizonal case")
		}
		// >[]
		block = g.data[y][x+2]
		if block == EMPTY {
			// [].  --> .[]
			g.data[y][x+0] = EMPTY
			g.data[y][x+1] = BOX_LEFT
			g.data[y][x+2] = BOX_RIGHT
		}
	} else if dx == -1 {
		// []<
		block := g.data[y][x-1]
		switch block {
		case EMPTY, OBSTACLE:
			break
		case BOX_RIGHT:
			// [][]
			g.pushHorizontal(x-2, y, dx)
		default:
			log.Panic("unhandled push horizonal case")
		}
		block = g.data[y][x-1]
		if block == EMPTY {
			// .[]  --> [].
			g.data[y][x+1] = EMPTY
			g.data[y][x-0] = BOX_RIGHT
			g.data[y][x-1] = BOX_LEFT
		}
	} else {
		panic("invalid delta x for horizontal pushing")
	}
}

// Remaining Problem:
// We push some boxes when we shouldn't:
// 2024/12/22 17:27:32 move ^:
// ##############
// ##......##..##
// ##...[][]...##
// ##....[]....##
// ##.....@....##
// ##..........##
// ##############
// 2024/12/22 17:27:32 pushing
// 2024/12/22 17:27:32 finalize push (0, -1) 7, 4
// 2024/12/22 17:27:32 move ^:
// ##############
// ##...[].##..##
// ##.....[]...##
// ##....[]....##
// ##.....@....##
// ##..........##
// ##############
//
// Fix idea:
//  similar helper function to pushVertical
//  determines a depth how far we can push
//  (takes the min of the recursive calls)
//  Then the second update step only
//  pushing until this depth is reached

// (x, y) is the position of the left half of the box [
// the box is pushed in direction (0, dy)
//
// instead, let's make it iterative
func (g *Grid) pushVertical(x, y, dy int) {
	worklist := []int{x}
	todos := make([]struct{ x, y int }, 0)
	ok := true
	for yNext := y + dy; ok && len(worklist) > 0; yNext += dy {
		logger.Printf("pushVertial y=%d, xs=%v", yNext, worklist)
		next := []int{}
		for _, xNext := range worklist {
			if g.data[yNext-dy][xNext] != BOX_LEFT {
				panic("box invariant violated")
			}
			bL := g.data[yNext][xNext]
			bR := g.data[yNext][xNext+1]
			switch {
			case bL == OBSTACLE || bR == OBSTACLE:
				// cannot push further
				// #?, ?#
				// [], []
				ok = false
				break
			case bL == EMPTY && bR == EMPTY:
				// ..   []
				// []   ..
			case bL == BOX_LEFT && bR == BOX_RIGHT:
				// Continue pushing in this line
				// []  []
				// []
				next = append(next, xNext)
			case bL == BOX_RIGHT && bR == EMPTY:
				// Continue pushing the left box
				// [].
				//  []
				next = append(next, xNext-1)
			case bL == EMPTY && bR == BOX_LEFT: // .[
				// Continue pushing the right box
				// .[]
				// []
				next = append(next, xNext+1)
			case bL == BOX_RIGHT && bR == BOX_LEFT: // ][
				// Continue pushing both boxes
				// [][]
				//  []
				next = append(next, xNext-1)
				next = append(next, xNext+1)
			default:
				panic("unhandled pushVertial case")
			}
		}
		if ok {
			for _, xNext := range worklist {
				todos = append(todos, struct{ x, y int }{xNext, yNext})
			}
		}
		worklist = next
	}
	for i := range todos {
		t := todos[len(todos)-1-i]
		// try move the box forward
		if g.data[t.y][t.x] == EMPTY && g.data[t.y][t.x+1] == EMPTY {
			// ..  []
			// []  ..
			g.data[t.y][t.x] = BOX_LEFT
			g.data[t.y][t.x+1] = BOX_RIGHT
			g.data[t.y-dy][t.x] = EMPTY
			g.data[t.y-dy][t.x+1] = EMPTY
		}
	}
}

func (g *Grid) Update2(x, y int, move Direction) (xOut, yOut int) {
	dx, dy := DirToDelta(move)
	xNext, yNext := x+dx, y+dy

	blockNext, _ := g.At(xNext, yNext)

	switch blockNext {
	case EMPTY:
		logger.Println("walking")
		xOut, yOut = xNext, yNext
	case BOX_LEFT, BOX_RIGHT:
		logger.Println("pushing")
		xOut, yOut = g.Push(x, y, dx, dy)
	case OBSTACLE:
		logger.Println("run into obstacle")
		xOut, yOut = x, y
	default:
		logger.Panic("unexpected block")
	}

	// swap the robots position
	g.data[yOut][xOut], g.data[y][x] = g.data[y][x], g.data[yOut][xOut]
	return
}

func (g *Grid) Equals(o Grid) bool {
	if g.width != o.width || g.height != o.height {
		return false
	}
	for y := range g.height {
		if !slices.Equal(g.data[y], o.data[y]) {
			return false
		}
	}
	return true
}

func (g *Grid) At(x, y int) (byte, error) {
	if x < 0 || x >= g.width {
		return 0, fmt.Errorf("x coordinate %d out of bounds (%d, %d)", x, g.width, g.height)
	}
	if y < 0 || y >= g.height {
		return 0, fmt.Errorf("y coordinate %d out of bounds (%d, %d)", y, g.width, g.height)
	}
	return g.data[y][x], nil
}

func DirToDelta(d Direction) (x, y int) {
	switch d {
	case UP:
		return 0, -1
	case RIGHT:
		return 1, 0
	case DOWN:
		return 0, 1
	case LEFT:
		return -1, 0
	default:
		panic("unexpected direction")
	}
}

type Direction byte

const (
	UP    = '^'
	RIGHT = '>'
	DOWN  = 'v'
	LEFT  = '<'
)

type Cell byte

const (
	BOX       = 'O'
	BOX_LEFT  = '['
	BOX_RIGHT = ']'
	OBSTACLE  = '#'
	EMPTY     = '.'
	ROBOT     = '@'
)

type Input struct {
	grid      Grid
	movements []byte
}

func Parse(path string) Input {
	bs := bytes.Split(helper.ReadFile(path), []byte{'\n', '\n'})
	data := bytes.Split(bs[0], []byte{'\n'})

	movements := []byte{}
	if len(bs) >= 2 {
		movements = bytes.ReplaceAll(bs[1], []byte{'\n'}, []byte{})
	}

	return Input{
		grid: Grid{
			width:  len(data[0]),
			height: len(data),
			data:   data,
		},
		movements: movements,
	}

}

var logger = log.Default()

func Simulate(inp Input) {
	state := inp.grid
	_, x, y := state.FindFirst(ROBOT)
	logger.Printf("Initial state:\n%s\n", state)
	for _, move := range inp.movements {
		x, y = state.Update(x, y, Direction(move))
		logger.Printf("move %c:\n%s\n", move, state)
	}
	logger.Printf("final:\nposition: (%d, %d)\n%s\n\n", x, y, state)
}

func Part1(inp Input) (res int) {
	Simulate(inp)

	for y, line := range inp.grid.data {
		for x, box := range line {
			if box == BOX {
				res += y*100 + x
			}
		}
	}
	return res
}

// Observations:
//  invariant: the boxes [] stick together
//    we don't need to store their association
//
//  our swapping trick wont work anymore
//
// Can effect a large area
//    [][]# []
//       [][]
//      [][]
//       []
//       @
//
//        #            []#[]
//    [][] []        [] [][]
//       [][]          [][]
//      [][]            []
//       []	           ^@^
//      ^@^
//
// Idea: propagate per level
//
// In horizontal direction, its still only lines

func Simulate2(inp Input) {
	state := inp.grid
	_, x, y := state.FindFirst(ROBOT)
	logger.Printf("Initial state:\n%s\n", state)
	for _, move := range inp.movements {
		x, y = state.Update2(x, y, Direction(move))
		logger.Printf("move %c:\n%s\n", move, state)
	}
	time.Sleep(time.Millisecond * 100)
	logger.Printf("final:\nposition: (%d, %d)\n%s\n\n", x, y, state)
}

func Part2(inp Input) (res int) {
	Simulate2(inp)
	for y, line := range inp.grid.data {
		for x, box := range line {
			if box == BOX_LEFT {
				res += y*100 + x
			}
		}
	}
	return res
}

func driver(path string) {
	fmt.Println(path)
	input := Parse(path)
	// res1 := Part1(input)
	// fmt.Println("result part1:\n", res1)

	input.grid = *input.grid.Upscale()
	res2 := Part2(input)
	fmt.Println("result part2:\n", res2)
}
func main() {
	logger.SetOutput(os.Stdout)
	// driver("test_small")
	driver("test")
	// driver("input")
}
