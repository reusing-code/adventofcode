package main

import (
	"strings"

	h "github.com/reusing-code/adventofcode/gohelpers"
)

type springState int

const (
	operational springState = iota
	broken
	unknown
)

func possibleArrangements(springs []springState, groups []int) int {
	checkGroups := make([]int, 0)
	currentGroupSize := 0
	for i, v := range springs {
		if v == unknown {
			copy := append([]springState(nil), springs...)
			copy[i] = broken
			springs[i] = operational
			return possibleArrangements(springs, groups) + possibleArrangements(copy, groups)
		}
		if v == broken {
			currentGroupSize++
		} else {
			if currentGroupSize != 0 {
				checkGroups = append(checkGroups, currentGroupSize)
			}
			currentGroupSize = 0
		}
	}
	if currentGroupSize != 0 {
		checkGroups = append(checkGroups, currentGroupSize)
	}
	groupsSame := len(groups) == len(checkGroups)
	if groupsSame {
		for i, v := range groups {
			if checkGroups[i] != v {
				groupsSame = false
			}
		}
	}
	if groupsSame {
		return 1
	}
	return 0
}

func puzzle1(input []string) int {
	result := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		groups := h.ParseIntVecSingleLine(split[1])
		springs := make([]springState, 0)
		for _, c := range split[0] {
			switch c {
			case '#':
				springs = append(springs, broken)
			case '.':
				springs = append(springs, operational)
			case '?':
				springs = append(springs, unknown)
			}
		}
		arrangements := possibleArrangements(springs, groups)
		result += arrangements
	}

	return result
}

func puzzle2(input []string) int {
	result := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		groups := h.ParseIntVecSingleLine(split[1])
		springs := make([]springState, 0)
		for _, c := range split[0] {
			switch c {
			case '#':
				springs = append(springs, broken)
			case '.':
				springs = append(springs, operational)
			case '?':
				springs = append(springs, unknown)
			}
		}
		springsUnfolded := make([]springState, 0, 5*len(springs)+4)
		springsUnfolded = append(springsUnfolded, springs...)
		for i := 0; i < 4; i++ {
			springsUnfolded = append(springsUnfolded, unknown)
			springsUnfolded = append(springsUnfolded, springs...)
		}
		groupsUnfolded := make([]int, 0, 5*len(groups))
		for i := 0; i < 5; i++ {
			groupsUnfolded = append(groupsUnfolded, groups...)
		}
		arrangements := possibleArrangements(springsUnfolded, groupsUnfolded)
		result += arrangements
	}

	return result
}
