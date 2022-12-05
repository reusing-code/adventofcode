package main

import (
	"fmt"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type stack = []string

func parseStacks(in []string) []stack {
	size := (len(in[0]) + 1) / 4
	stacks := make([]stack, size)
	for i := len(in) - 2; i >= 0; i-- {
		for j := 0; j < size; j++ {
			crate := in[i][j*4 : j*4+3]
			if crate != "   " {
				stacks[j] = append(stacks[j], string(crate[1]))
			}
		}
	}
	return stacks
}

func puzzle1(input []string) string {
	splitInput := gohelpers.SplitByEmptyLine(input)
	stacks := parseStacks(splitInput[0])
	moves := splitInput[1]

	for _, move := range moves {
		count := 0
		from := 0
		to := 0
		fmt.Sscanf(move, "move %d from %d to %d", &count, &from, &to)
		from--
		to--
		lf := len(stacks[from])
		cratesToMove := stacks[from][lf-count:]
		for i := len(cratesToMove) - 1; i >= 0; i-- {
			stacks[to] = append(stacks[to], cratesToMove[i])
		}
		stacks[from] = stacks[from][:lf-count]
	}

	result := ""
	for _, s := range stacks {
		result += string(s[len(s)-1])
	}
	return result
}

func puzzle2(input []string) string {
	splitInput := gohelpers.SplitByEmptyLine(input)
	stacks := parseStacks(splitInput[0])
	moves := splitInput[1]

	for _, move := range moves {
		count := 0
		from := 0
		to := 0
		fmt.Sscanf(move, "move %d from %d to %d", &count, &from, &to)
		from--
		to--
		lf := len(stacks[from])
		cratesToMove := stacks[from][lf-count:]
		stacks[to] = append(stacks[to], cratesToMove...)
		stacks[from] = stacks[from][:lf-count]
	}

	result := ""
	for _, s := range stacks {
		result += string(s[len(s)-1])
	}
	return result
}
