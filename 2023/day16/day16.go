package main

import (
	"fmt"

	h "github.com/reusing-code/adventofcode/gohelpers"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

func nextDirs(dir direction, c byte) []direction {
	switch c {
	case '.':
		return []direction{dir}
	case '/':
		switch dir {
		case south:
			return []direction{west}
		case west:
			return []direction{south}
		case north:
			return []direction{east}
		case east:
			return []direction{north}
		}
	case '\\':
		switch dir {
		case south:
			return []direction{east}
		case west:
			return []direction{north}
		case north:
			return []direction{west}
		case east:
			return []direction{south}
		}
	case '-':
		switch dir {
		case south:
			return []direction{east, west}
		case west:
			return []direction{west}
		case north:
			return []direction{east, west}
		case east:
			return []direction{east}
		}
	case '|':
		switch dir {
		case south:
			return []direction{south}
		case west:
			return []direction{north, south}
		case north:
			return []direction{north}
		case east:
			return []direction{north, south}
		}
	}
	return []direction{}
}

type step struct {
	c   h.Coord
	dir direction
}

func moveDir(c h.Coord, dir direction) h.Coord {
	switch dir {
	case north:
		return h.Coord{c.X - 1, c.Y}
	case east:
		return h.Coord{c.X, c.Y + 1}
	case south:
		return h.Coord{c.X + 1, c.Y}
	case west:
		return h.Coord{c.X, c.Y - 1}
	}
	return h.Coord{0, 0}
}

func nextSteps(field [][]byte, s step) []step {
	height := len(field)
	width := len(field[0])
	result := []step{}

	dirs := nextDirs(s.dir, field[s.c.X][s.c.Y])
	for _, d := range dirs {
		c := moveDir(s.c, d)
		if c.X < 0 || c.X >= height || c.Y < 0 || c.Y >= width {
			continue
		}
		result = append(result, step{c, d})
	}
	return result
}

func hashKey(s step) string {
	return fmt.Sprintf("%d,%d,%d", s.c.X, s.c.Y, s.dir)
}

func countEnergized(field [][]byte, start step) int {
	energized := make(map[h.Coord]bool)
	energized[start.c] = true
	visited := make(map[string]bool)
	stack := []step{start}
	visited[hashKey(stack[0])] = true
	for len(stack) > 0 {
		next := nextSteps(field, stack[0])
		for _, n := range next {
			energized[n.c] = true
			hk := hashKey(n)
			if _, ok := visited[hk]; !ok {
				visited[hk] = true
				stack = append(stack, n)
			}
		}
		stack = stack[1:]
	}
	return len(energized)
}

func puzzle1(input []string) int {
	field := h.ParseCharField(input)
	return countEnergized(field, step{h.Coord{0, 0}, east})
}

func puzzle2(input []string) int {
	field := h.ParseCharField(input)
	height := len(field)
	width := len(field[0])
	max := 0
	// left
	for i := 0; i < height; i++ {
		e := countEnergized(field, step{h.Coord{i, 0}, east})
		if e > max {
			max = e
		}
	}
	// top
	for i := 0; i < width; i++ {
		e := countEnergized(field, step{h.Coord{0, i}, south})
		if e > max {
			max = e
		}
	}
	// right
	for i := 0; i < height; i++ {
		e := countEnergized(field, step{h.Coord{i, width - 1}, west})
		if e > max {
			max = e
		}
	}
	// bottom
	for i := 0; i < width; i++ {
		e := countEnergized(field, step{h.Coord{height - 1, i}, north})
		if e > max {
			max = e
		}
	}
	return max
}
