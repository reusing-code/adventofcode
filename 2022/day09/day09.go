package main

import (
	"fmt"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type puzzle int

const (
	p1 puzzle = iota
	p2
)

type coord struct {
	x int
	y int
}

type rope struct {
	head        coord
	tails       []coord
	tailVisited map[coord]bool
}

func createRope(numberOfTails int) *rope {
	r := &rope{coord{0, 0}, make([]coord, 0), make(map[coord]bool)}
	r.tailVisited[coord{0, 0}] = true
	for i := 0; i < numberOfTails; i++ {
		r.tails = append(r.tails, coord{0, 0})
	}
	return r
}

func (r *rope) moveTail(noOfTails int) {
	for i := 0; i < noOfTails; i++ {
		var h *coord
		if i == 0 {
			h = &r.head
		} else {
			h = &r.tails[i-1]
		}
		t := &r.tails[i]
		if gohelpers.Abs(h.x-t.x) <= 1 && gohelpers.Abs(h.y-t.y) <= 1 {
			return
		}
		t.x += gohelpers.Sign(h.x - t.x)
		t.y += gohelpers.Sign(h.y - t.y)
	}
	r.tailVisited[r.tails[noOfTails-1]] = true
}

func (r *rope) up() {
	r.head.x += 1
}

func (r *rope) down() {
	r.head.x -= 1
}

func (r *rope) left() {
	r.head.y -= 1
}

func (r *rope) right() {
	r.head.y += 1
}

func runpuzzle(input []string, p puzzle) int {
	numberOfTails := 1
	if p == p2 {
		numberOfTails = 9
	}
	r := createRope(numberOfTails)
	for _, inLine := range input {
		direction := ""
		moves := 0
		fmt.Sscanf(inLine, "%s %d", &direction, &moves)
		for i := 0; i < moves; i++ {
			switch direction {
			case "U":
				r.up()
			case "D":
				r.down()
			case "L":
				r.left()
			case "R":
				r.right()
			}
			r.moveTail(numberOfTails)
		}
	}
	return len(r.tailVisited)
}

func puzzle1(input []string) int {
	return runpuzzle(input, p1)
}

func puzzle2(input []string) int {
	return runpuzzle(input, p2)
}
