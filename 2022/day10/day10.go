package main

import (
	"fmt"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type instructionType int

const (
	addx instructionType = iota
	noop
)

type instruction struct {
	iType instructionType
	param int
}

type program struct {
	instructions []instruction
	x            int
}

func (p *program) parseInstructions(input []string) {
	p.instructions = make([]instruction, 0, len(input))
	for _, str := range input {
		if str == "noop" {
			p.instructions = append(p.instructions, instruction{noop, 0})
		} else {
			addVal := 0
			fmt.Sscanf(str, "addx %d", &addVal)
			p.instructions = append(p.instructions, instruction{addx, addVal})
		}
	}
}

func (p *program) run(measureCycles []int) int {
	result := 0
	cycle := 0
	nextMeasure := 0
	for _, inst := range p.instructions {
		add := 0
		switch inst.iType {
		case noop:
			cycle += 1
		case addx:
			cycle += 2
			add += inst.param
		}

		if cycle >= measureCycles[nextMeasure] {
			result += p.x * measureCycles[nextMeasure]
			nextMeasure++
			if nextMeasure >= len(measureCycles) {
				return result
			}
		}
		p.x += add
	}
	return result
}

func paint(crt [][]bool, x int, cycle int) {
	row := int(cycle / len(crt[0]))
	column := cycle % len(crt[0])
	val := false
	if gohelpers.Abs(x-column) <= 1 {
		val = true
	}
	crt[row][column] = val
}

func (p *program) run2(crt [][]bool) {
	cycle := 0
	for _, inst := range p.instructions {
		cycle += 1
		if cycle >= len(crt)*len(crt[0]) {
			return
		}
		paint(crt, p.x, cycle-1)
		if inst.iType == addx {
			cycle += 1
			if cycle >= len(crt)*len(crt[0]) {
				return
			}
			paint(crt, p.x, cycle-1)
		}
		p.x += inst.param
	}
	return
}

func puzzle1(input []string) int {
	p := &program{nil, 1}
	p.parseInstructions(input)
	return p.run([]int{20, 60, 100, 140, 180, 220})
}

func puzzle2(input []string) int {
	p := &program{nil, 1}
	crt := make([][]bool, 6)
	for i := range crt {
		crt[i] = make([]bool, 40)
	}
	p.parseInstructions(input)
	p.run2(crt)

	for _, line := range crt {
		for _, val := range line {
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	return 1
}
