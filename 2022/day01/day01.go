package main

import (
	"sort"
	"strconv"
)

func puzzle1(input []string) int {
	maxElf := 0
	currentElf := 0
	for _, val := range input {
		if len(val) == 0 {
			if currentElf > maxElf {
				maxElf = currentElf
			}
			currentElf = 0
			continue
		}
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		currentElf += v
	}
	if currentElf > maxElf {
		maxElf = currentElf
		currentElf = 0
	}
	return maxElf
}

func puzzle2(input []string) int {
	maxElves := make([]int, 3)
	currentElf := 0
	for _, val := range input {
		if len(val) == 0 {
			if currentElf > maxElves[0] {
				maxElves[0] = currentElf
				sort.Ints(maxElves)
			}
			currentElf = 0
			continue
		}
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		currentElf += v
	}
	if currentElf > maxElves[0] {
		maxElves[0] = currentElf
	}
	sum := 0
	for _, v := range maxElves {
		sum += v
	}
	return sum
}
