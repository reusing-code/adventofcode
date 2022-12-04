package main

import (
	"fmt"
)

type sectionRange struct {
	from, to int
}

func parseRanges(in string) [2]sectionRange {
	result := [2]sectionRange{{1, 2}, {2, 3}}
	fmt.Sscanf(in, "%d-%d,%d-%d", &result[0].from, &result[0].to, &result[1].from, &result[1].to)
	return result
}

func fullyContained(in string) bool {
	r := parseRanges(in)
	r1, r2 := r[0], r[1]
	if r1.to-r1.from < r2.to-r2.from {
		r1, r2 = r2, r1
	}
	res := r2.from >= r1.from && r2.to <= r1.to
	return res
}

func anyOverlap(in string) bool {
	r := parseRanges(in)
	r1, r2 := r[0], r[1]
	if r1.to > r2.to {
		r1, r2 = r2, r1
	}
	return r2.from <= r1.to
}

func puzzle1(input []string) int {
	result := 0
	for _, v := range input {
		if fullyContained(v) {
			result += 1
		}
	}
	return result
}

func puzzle2(input []string) int {
	result := 0
	for _, v := range input {
		if anyOverlap(v) {
			result += 1
		}
	}
	return result
}
