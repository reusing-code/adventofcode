package main

import (
	"github.com/reusing-code/adventofcode/gohelpers"
)

type directionType int

const (
	R directionType = iota
	D
	L
	U
)

type hurricaneType map[gohelpers.Coord][]directionType

func (h *hurricaneType) add(c *gohelpers.Coord, dir directionType) {
	if _, prs := (*h)[*c]; !prs {
		(*h)[*c] = make([]directionType, 0)
	}
	(*h)[*c] = append((*h)[*c], dir)
}

var (
	width  int
	height int
)

func doHurricaneMin(old *hurricaneType) *hurricaneType {
	newHurricanes := make(hurricaneType)
	for k, v := range *old {
		for _, d := range v {
			coord := k
			switch d {
			case R:
				coord.Y++
				if coord.Y == width-1 {
					coord.Y = 1
				}
			case D:
				coord.X++
				if coord.X == height-1 {
					coord.X = 1
				}
			case L:
				coord.Y--
				if coord.Y == 0 {
					coord.Y = width - 2
				}
			case U:
				coord.X--
				if coord.X == 0 {
					coord.X = height - 2
				}
			}
			newHurricanes.add(&coord, d)
		}
	}
	return &newHurricanes
}

func travelTime(field *[][]byte, hurricanes *hurricaneType, start, end gohelpers.Coord) (int, *hurricaneType) {
	reachable := make(map[gohelpers.Coord]bool)
	reachable[start] = true
	currentHurricanes := hurricanes

	for min := 1; true; min++ {
		currentHurricanes = doHurricaneMin(currentHurricanes)
		newReachable := make(map[gohelpers.Coord]bool)
		check := func(x, y int) {
			if x < 0 || x == len(*field) {
				return
			}
			if (*field)[x][y] == '#' {
				return
			}
			c := gohelpers.Coord{x, y}
			if _, prs := (*currentHurricanes)[c]; prs {
				return
			}
			newReachable[c] = true
		}
		for k := range reachable {
			check(k.X, k.Y)
			check(k.X, k.Y+1)
			check(k.X, k.Y-1)
			check(k.X+1, k.Y)
			check(k.X-1, k.Y)
		}
		if _, prs := newReachable[end]; prs {
			return min, currentHurricanes
		}
		reachable = newReachable
	}
	return 0, nil
}

func parseInput(input []string) ([][]byte, hurricaneType, gohelpers.Coord, gohelpers.Coord) {
	field := gohelpers.ParseCharField(input)
	start := gohelpers.Coord{}
	end := gohelpers.Coord{}
	width = len(field[0])
	height = len(field)
	for i, v := range field[0] {
		if v == '.' {
			start.Y = i
			break
		}
	}
	for i, v := range field[len(field)-1] {
		if v == '.' {
			end.X = len(field) - 1
			end.Y = i
		}
	}

	hurricanes := make(hurricaneType)

	for x, line := range field {
		for y, v := range line {
			switch v {
			case '>':
				hurricanes.add(&gohelpers.Coord{x, y}, R)
			case 'v':
				hurricanes.add(&gohelpers.Coord{x, y}, D)
			case '<':
				hurricanes.add(&gohelpers.Coord{x, y}, L)
			case '^':
				hurricanes.add(&gohelpers.Coord{x, y}, U)
			}
			if v != '#' {
				field[x][y] = '.'
			}
		}
	}
	return field, hurricanes, start, end
}

func puzzle1(input []string) int {
	field, hurricanes, start, end := parseInput(input)

	result, _ := travelTime(&field, &hurricanes, start, end)

	return result
}

func puzzle2(input []string) int {
	field, hurricanes, start, end := parseInput(input)

	r1, currentHurricanes := travelTime(&field, &hurricanes, start, end)
	r2, currentHurricanes := travelTime(&field, currentHurricanes, end, start)
	r3, _ := travelTime(&field, currentHurricanes, start, end)

	return r1 + r2 + r3
}
