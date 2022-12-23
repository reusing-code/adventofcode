package main

import (
	"fmt"

	"github.com/reusing-code/adventofcode/gohelpers"
)

func parseField(input []string) map[gohelpers.Coord]gohelpers.Coord {
	result := make(map[gohelpers.Coord]gohelpers.Coord)
	charField := gohelpers.ParseCharField(input)
	for i, line := range charField {
		for j, v := range line {
			if v == '#' {
				result[gohelpers.Coord{i, j}] = gohelpers.Coord{i, j}
			}
		}
	}
	return result
}

type proposition struct {
	name  string
	dir   gohelpers.Coord
	check []gohelpers.Coord
}

var propositions []proposition = []proposition{
	{name: "N", dir: gohelpers.Coord{-1, 0}, check: []gohelpers.Coord{{-1, -1}, {-1, 0}, {-1, 1}}},
	{name: "S", dir: gohelpers.Coord{1, 0}, check: []gohelpers.Coord{{1, -1}, {1, 0}, {1, 1}}},
	{name: "W", dir: gohelpers.Coord{0, -1}, check: []gohelpers.Coord{{-1, -1}, {0, -1}, {1, -1}}},
	{name: "E", dir: gohelpers.Coord{0, 1}, check: []gohelpers.Coord{{-1, 1}, {0, 1}, {1, 1}}},
}

func minMax(field map[gohelpers.Coord]gohelpers.Coord) (gohelpers.Coord, gohelpers.Coord) {
	max := gohelpers.Coord{-(1 << 32), -(1 << 32)}
	min := gohelpers.Coord{1 << 32, 1 << 32}
	for c := range field {
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
	}
	return min, max
}

func draw(field map[gohelpers.Coord]gohelpers.Coord) {
	min, max := minMax(field)
	draw := make([][]byte, 0, max.X-min.X+1)
	for x := min.X; x <= max.X; x++ {
		draw = append(draw, make([]byte, 0, max.Y-min.Y+1))
		for y := min.Y; y <= max.Y; y++ {
			draw[x-min.X] = append(draw[x-min.X], '.')
		}
	}
	for elve := range field {
		draw[elve.X-min.X][elve.Y-min.Y] = '#'
	}

	for _, line := range draw {
		for _, v := range line {
			fmt.Print(string(v))
		}
		fmt.Println("")
	}
	fmt.Println("")
	return
}

func simulateRound(elveField map[gohelpers.Coord]gohelpers.Coord, round int) (map[gohelpers.Coord]gohelpers.Coord, bool) {
	anyElveMoved := false
	targets := make(map[gohelpers.Coord]gohelpers.Coord)
	for elve := range elveField {
		moved := false
		needsToMove := false
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				if _, prs := elveField[gohelpers.Coord{elve.X + dx, elve.Y + dy}]; prs {
					needsToMove = true
				}
			}
		}
		if needsToMove {
			for pi := range propositions {
				prop := propositions[(round+pi)%len(propositions)]
				canMove := true
				for _, c := range prop.check {
					if _, prs := elveField[c.Add(elve)]; prs {
						canMove = false
					}
				}
				if canMove {
					target := prop.dir.Add(elve)
					if _, prs := targets[target]; prs {
						targets[elve] = elve
						targets[targets[target]] = targets[target]
						delete(targets, target)
					} else {
						targets[target] = elve
					}
					moved = true
					break
				}
			}
		}
		if !moved {
			targets[elve] = elve
		} else {
			anyElveMoved = true
		}
	}
	return targets, anyElveMoved
}

func puzzle1(input []string) int {
	elveField := parseField(input)

	for i := 0; i < 10; i++ {
		elveField, _ = simulateRound(elveField, i)
	}

	min, max := minMax(elveField)
	return (max.X-min.X+1)*(max.Y-min.Y+1) - len(elveField)
}

func puzzle2(input []string) int {
	elveField := parseField(input)

	for i := 0; true; i++ {
		var moved bool
		elveField, moved = simulateRound(elveField, i)
		if !moved {
			return i + 1
		}
	}
	return 0
}
