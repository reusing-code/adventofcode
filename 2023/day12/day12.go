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

func hashKey(springs []springState, groups []int) string {
	result := make([]byte, 0, len(springs)+1+len(groups))
	for _, v := range springs {
		result = append(result, byte(v))
	}
	result = append(result, 255)
	for _, v := range groups {
		if v >= 255 {
			panic("unexpected value")
		}
		result = append(result, byte(v))

	}
	return string(result)
}

var cache = map[string]int{}

func possibleArrangements(springs []springState, groups []int) int {
	if len(springs) == 0 {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}
	if len(groups) == 0 {
		for _, v := range springs {
			if v == broken {
				return 0
			}
		}
		return 1
	}
	minLength := 0
	for _, v := range groups {
		minLength += v
	}
	minLength += len(groups) - 1
	if len(springs) < minLength {
		return 0
	}

	k := hashKey(springs, groups)
	if v, ok := cache[k]; ok {
		return v
	}

	result := possibleArrangementsStep2(springs, groups)
	cache[k] = result
	return result
}

func possibleArrangementsStep2(springs []springState, groups []int) int {
	if springs[0] == operational {
		i := 0
		for ; i < len(springs); i++ {
			if springs[i] != operational {
				break
			}
		}
		return possibleArrangements(springs[i:], groups)
	}
	// -----------------

	arrangements := 0
	if springs[0] == unknown {
		arrangements += possibleArrangements(springs[1:], groups)
	}
	for i := 0; i < groups[0]; i++ {
		if springs[i] == operational {
			return arrangements
		}
	}
	if len(springs) == groups[0] {
		return arrangements + 1
	}
	if springs[groups[0]] == broken {
		return arrangements
	}
	return arrangements + possibleArrangements(springs[groups[0]+1:], groups[1:])
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
