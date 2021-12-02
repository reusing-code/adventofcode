package main

import "github.com/reusing-code/adventofcode/gohelpers"

func puzzle1(input []string) int {
	vals := gohelpers.ParseIntVec(input)
	freq := 0
	for _, v := range vals {
		freq += v
	}
	return freq
}

func puzzle2(input []string) int {
	vals := gohelpers.ParseIntVec(input)
	freq := 0
	seen := make(map[int]bool, 0)
	seen[0] = true
	for i := 0; ; i = (i + 1) % len(vals) {
		freq += vals[i]
		if _, ok := seen[freq]; ok {
			return freq
		}
		seen[freq] = true
	}
	return freq
}
