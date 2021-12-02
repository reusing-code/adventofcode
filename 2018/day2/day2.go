package main

import "sort"

func puzzle1(input []string) int {
	twos := 0
	threes := 0
	for _, val := range input {
		seen := make(map[rune]int, 0)
		for _, ch := range val {
			seen[ch]++
		}
		two := false
		three := false
		for _, count := range seen {
			if count == 2 {
				two = true
			}
			if count == 3 {
				three = true
			}
		}
		if two {
			twos++
		}
		if three {
			threes++
		}
	}
	return twos * threes
}

func puzzle2(input []string) string {
	sort.Strings(input)
	prev := input[0]
	for i := 1; i < len(input); i++ {
		diff := -1
		diffCount := 0
		for j, _ := range input[i] {
			if prev[j] != input[i][j] {
				diffCount++
				diff = j
			}
		}
		if diffCount == 1 {
			return prev[0:diff] + prev[diff+1:]
		}
		prev = input[i]
	}
	return ""
}
