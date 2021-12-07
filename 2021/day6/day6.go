package main

import (
	"strings"
	"strconv"
)

func puzzle1(input []string) int {
	stringSlice := strings.Split(input[0], ",")

	var ages [9]int
	for _, str := range stringSlice {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ages[n]++
	}

	for i := 0; i < 80; i++ {
		var newAges [9]int
		for age, count := range ages {
			if age != 0 {
				newAges[age - 1] += count
			} else {
				newAges[6] += count
				newAges[8] += count
			}
		}
		ages = newAges
	}

	result := 0
	for _, count := range ages {
		result += count
	}

	return result
}

func puzzle2(input []string) int {
	stringSlice := strings.Split(input[0], ",")

	var ages [9]int
	for _, str := range stringSlice {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ages[n]++
	}

	for i := 0; i < 256; i++ {
		var newAges [9]int
		for age, count := range ages {
			if age != 0 {
				newAges[age - 1] += count
			} else {
				newAges[6] += count
				newAges[8] += count
			}
		}
		ages = newAges
	}

	result := 0
	for _, count := range ages {
		result += count
	}

	return result
}
