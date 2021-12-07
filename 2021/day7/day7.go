package main

import (
	"sort"

	"github.com/reusing-code/adventofcode/gohelpers"
)

func puzzle1(input []string) int {
	vec := gohelpers.ParseIntVecSingleLine(input[0])
	sort.Ints(vec)
	midIdx := len(vec) / 2
	midVal := vec[midIdx]
	fuel := 0
	for _, v := range vec {
		if v < midVal {
			fuel += midVal -v
		} else {
			fuel += v - midVal
		}
	}
	return fuel
}

func puzzle2(input []string) int {
	vec := gohelpers.ParseIntVecSingleLine(input[0])
	sort.Ints(vec)
	sum := 0
	for _, v := range vec {
		sum += v
	}	
	midVal := int(float64(sum) / float64(len(vec)))
	fuel1 := 0
	for _, v := range vec {
		var distance int
		if v < midVal {
			distance = midVal -v
		} else {
			distance = v - midVal
		}
		fuel1 += (distance * (distance + 1)) / 2
	}
	midVal++
	fuel2 := 0
	for _, v := range vec {
		var distance int
		if v < midVal {
			distance = midVal -v
		} else {
			distance = v - midVal
		}
		fuel2 += (distance * (distance + 1)) / 2
	}
	if fuel2 > fuel1 {
		return fuel1
	} else {
		return fuel2
	}
}
