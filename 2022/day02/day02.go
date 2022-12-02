package main

import "github.com/reusing-code/adventofcode/gohelpers"

var (
	lookupP1 = [3][3]int{{1 + 3, 2 + 6, 3 + 0}, {1 + 0, 2 + 3, 3 + 6}, {1 + 6, 2 + 0, 3 + 3}}
	lookupP2 = [3][3]int{{0 + 3, 3 + 1, 6 + 2}, {0 + 1, 3 + 2, 6 + 3}, {0 + 2, 3 + 3, 6 + 1}}
)

func calcScore(in []byte, lookup [3][3]int) int {
	return lookup[in[0]-'A'][in[2]-'X']
}

func puzzle1(input []string) int {
	field := gohelpers.ParseCharField(input)
	score := 0
	for _, round := range field {
		score += calcScore(round, lookupP1)
	}
	return score
}

func puzzle2(input []string) int {
	field := gohelpers.ParseCharField(input)
	score := 0
	for _, round := range field {
		score += calcScore(round, lookupP2)
	}
	return score
}
