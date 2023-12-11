package main

import (
	h "github.com/reusing-code/adventofcode/gohelpers"
)

func findS(field [][]byte) h.Coord {
	for i, v := range field {
		for j, c := range v {
			if c == byte('S') {
				return h.Coord{X: i, Y: j}
			}
		}
	}
	return invalid
}

var (
	invalid = h.Coord{-1, -1}
	north   = h.Coord{-1, 0}
	south   = h.Coord{1, 0}
	west    = h.Coord{0, -1}
	east    = h.Coord{0, 1}
)

type mapping map[h.Coord]h.Coord

var lookupMapping = map[byte]mapping{
	'|': map[h.Coord]h.Coord{north: south, south: north},
	'-': map[h.Coord]h.Coord{west: east, east: west},
	'L': map[h.Coord]h.Coord{north: east, east: north},
	'J': map[h.Coord]h.Coord{north: west, west: north},
	'7': map[h.Coord]h.Coord{west: south, south: west},
	'F': map[h.Coord]h.Coord{east: south, south: east},
}

func followLoop(field [][]byte, from, current h.Coord) h.Coord {
	diff := h.Coord{from.X - current.X, from.Y - current.Y}
	if v, ok := lookupMapping[field[current.X][current.Y]]; ok {
		if r, ok2 := v[diff]; ok2 {
			next := current.Add(r)
			if next.X < 0 || next.X >= len(field) || next.Y < 0 || next.Y >= len(field[next.X]) {
				return invalid
			}
			return next
		} else {
			return invalid
		}
	} else {
		return invalid
	}
}

func puzzle1(input []string) int {
	field := h.ParseCharField(input)
	start := findS(field)
	checkDirections := []h.Coord{north, east, south}
	for _, dir := range checkDirections {
		last := start
		current := start.Add(dir)
		for i := 1; ; i++ {
			next := followLoop(field, last, current)
			if next == invalid {
				break
			}
			if next == start {
				return (i + 1) / 2
			}
			last = current
			current = next
		}
	}
	return 0
}

func getLoop(field [][]byte, start h.Coord) []h.Coord {
	checkDirections := []h.Coord{north, east, south}
	for _, dir := range checkDirections {
		loop := []h.Coord{start}
		last := start
		current := start.Add(dir)
		for i := 1; ; i++ {
			loop = append(loop, current)
			next := followLoop(field, last, current)
			if next == invalid {
				break
			}
			if next == start {
				return loop
			}
			last = current
			current = next
		}
	}
	return []h.Coord{}
}

func puzzle2(input []string) int {
	field := h.ParseCharField(input)
	start := findS(field)
	loop := getLoop(field, start)
	newField := make([][]rune, len(field))
	for i, v := range field {
		newField[i] = make([]rune, len(v))
		for j := range v {
			newField[i][j] = '.'
		}
	}

	for _, v := range loop {
		var r rune
		switch field[v.X][v.Y] {
		case '|':
			r = '│'
		case '-':
			r = '─'
		case 'L':
			r = '└'
		case 'J':
			r = '┘'
		case '7':
			r = '┐'
		case 'F':
			r = '┌'
		case 'S':
			r = 'S'
		}
		newField[v.X][v.Y] = r
	}

	h.PrintRuneFieldToFile(newField, "out.txt")

	return 0
}
