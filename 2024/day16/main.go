package main

import (
	"aoc2024"
	"container/heap"
	"fmt"
	"slices"
)

type Grid = [][]byte

func FindFirst(g *Grid, target byte) (x int, y int) {
	for y, line := range *g {
		x := slices.Index(line, target)
		if x != -1 {
			return x, y
		}
	}
	panic("not found")
}

type Block byte

const (
	WALL  Block = '#'
	FREE        = '.'
	START       = 'S'
	END         = 'E'
)

type Direction int

const (
	EAST Direction = iota
	SOUTH
	WEST
	NORTH
)

func Parse(path string) Grid {
	return helper.ReadFileLines(path)
}

type State struct {
	x, y      int
	direction Direction
}

type Item struct {
	value    State // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
	// parent *Item
	path map[Coord]bool
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value State, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func nextStates(s State) (states []State) {
	x, y := s.x, s.y
	return []State{
		{x + 1, y, s.direction},
		{x - 1, y, s.direction},
		{x, y + 1, s.direction},
		{x, y - 1, s.direction},
		{x, y, (s.direction + 1) % 4},
		{x, y, (s.direction + 3) % 4},
	}
}

type Coord struct {
	x, y int
}

//	func Backtrack(i Item) map[Coord]bool {
//		m := make(map[Coord]bool)
//		m[Coord{i.value.x, i.value.y}] = true
//		for p := i.parent; p != nil; p = p.parent {
//			m[Coord{p.value.x, p.value.y}] = true
//		}
//		return m
//	}
func MapCopy(m map[Coord]bool) map[Coord]bool {
	res := make(map[Coord]bool)
	for k, v := range m {
		res[k] = v
	}
	return res
}

func Solve(grid *Grid, part2 bool) (res int) {

	// inspired by
	// https://github.com/teivah/advent-of-code/blob/main/2024/day16-go/main.go
	costs := make(map[State]int)
	var paths map[Coord]bool
	best := -1 << 31

	startX, startY := FindFirst(grid, START)
	endX, endY := FindFirst(grid, END)
	// reindeer starts facing east
	startState := State{startX, startY, EAST}
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		value:    startState,
		priority: 0,
		index:    0,
		// parent:   nil,
		path: map[Coord]bool{{startX, startY}: true},
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		prevCost, ok := costs[item.value]
		if ok && prevCost > item.priority {
			continue
		}
		costs[item.value] = item.priority

		x, y := item.value.x, item.value.y
		if x == endX && y == endY {
			// path := Backtrack(*item)
			if item.priority > best {
				// best shortest path
				best = item.priority
				paths = item.path
			} else if item.priority == best {
				// tied cost path
				for c := range item.path {
					paths[c] = true
				}
			}
			continue
		}
		// - Continue moving forward
		c := Coord{x, y}.Forward(item.value.direction)
		if (*grid)[c.y][c.x] != byte(WALL) {
			path := MapCopy(item.path)
			path[c] = true
			moveForward := Item{
				value: State{
					x:         c.x,
					y:         c.y,
					direction: item.value.direction,
				},
				priority: item.priority - 1,
				index:    -1,
				// parent:   item,
				path: path,
			}
			heap.Push(&pq, &moveForward)
		}
		// - Turn 90°
		heap.Push(&pq, &Item{
			value: State{
				x:         item.value.x,
				y:         item.value.y,
				direction: (item.value.direction + 1) % 4,
			},
			priority: item.priority - 1000,
			index:    -1,
			// parent:   item,
			path: item.path,
		})
		// - Turn -90°
		heap.Push(&pq, &Item{
			value: State{
				x:         item.value.x,
				y:         item.value.y,
				direction: (item.value.direction + 3) % 4,
			},
			priority: item.priority - 1000,
			index:    -1,
			// parent:   item,
			path: item.path,
		})

	}
	// TODO for part2 we might need to continue
	// searching for other best paths
	// (drain the pq)
	// state := startState
	// cost := 0

	// for state.x != endX && state.y != endY {
	// 	x, y := state.x, state.y

	// 	for _, sn := range nextStates(state) {
	// 		cn := costs[sn]
	// 	}
	// }
	if part2 {
		// fmt.Println(paths)
		return len(paths)
	}
	return best
}

func (c Coord) Forward(d Direction) Coord {
	x, y := c.x, c.y
	switch d {
	default:
		panic("unkown direction")
	case EAST:
		x += 1
	case SOUTH:
		y += 1
	case WEST:
		x -= 1
	case NORTH:
		y -= 1
	}
	return Coord{x, y}
}

func driver(path string) {
	fmt.Println(path)
	input := Parse(path)
	// res1 := Solve(&input, false)
	// fmt.Println("result part1:\n", res1)
	res2 := Solve(&input, true)
	fmt.Println("result part2:\n", res2)
}
func main() {
	driver("test1")
	driver("test2")
	driver("input")
}
