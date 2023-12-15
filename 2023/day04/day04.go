package main

import (
	"math"
	"slices"
	"strings"

	"github.com/reusing-code/adventofcode/gohelpers"
)

func intersect(in1, in2 []int) []int {
	result := make([]int, 0)
	slices.Sort(in1)
	slices.Sort(in2)
	i1, i2 := 0, 0
	for i1 < len(in1) && i2 < len(in2) {
		if in1[i1] == in2[i2] {
			result = append(result, in1[i1])
		}
		if in1[i1] < in2[i2] {
			i1++
		} else {
			i2++
		}
	}
	return result
}

func puzzle1(input []string) int {
	result := 0
	for _, line := range input {
		splitColon := strings.Split(line, ":")
		splitPipe := strings.Split(splitColon[1], "|")
		winningNumbers := gohelpers.ParseIntVecSingleLineSpace(splitPipe[0])
		playedNumbers := gohelpers.ParseIntVecSingleLineSpace(splitPipe[1])
		winners := intersect(winningNumbers, playedNumbers)
		if len(winners) > 0 {
			result += int(math.Pow(2, float64(len(winners)-1)))
		}
	}
	return result
}

func puzzle2(input []string) int {
	scratchCards := make([]int, len(input))
	for i := range scratchCards {
		scratchCards[i] = 1
	}

	for card, line := range input {
		splitColon := strings.Split(line, ":")
		splitPipe := strings.Split(splitColon[1], "|")
		winningNumbers := gohelpers.ParseIntVecSingleLineSpace(splitPipe[0])
		playedNumbers := gohelpers.ParseIntVecSingleLineSpace(splitPipe[1])
		winners := intersect(winningNumbers, playedNumbers)
		for i := 1; i <= len(winners); i++ {
			scratchCards[card+i] += scratchCards[card]
		}
	}

	result := 0
	for _, v := range scratchCards {
		result += v
	}
	return result
}
