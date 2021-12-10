package main

import (
	"sort"
)

var opening = []rune{'(', '[', '{', '<'}
var closing = []rune{')', ']', '}', '>'}
var valueMap = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var valueMap2 = map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}


func puzzle1(input []string) int {
	result := 0
	for _, str := range input {
		stack := make([]rune, 0)
		err := false
		for _, r := range str {
			if err {
				break
			}
			open := false
			for _, o := range opening {
				if r == o {
					open = true
					break
				}
			}
			if open {
				stack = append(stack, r)
			} else {
				for i, c := range closing {
					if r == c {
						// matching closing bracket
						if opening[i] == stack[len(stack)-1] {
							stack = stack[:len(stack)-1]
							break
						} else {
							// un matched closing bracket
							result += valueMap[r]
							err = true
						}
					}
				}
			}
		}
	}
	return result
}

func puzzle2(input []string) int {
	scores := make([]int, 0)
	for _, str := range input {
		stack := make([]rune, 0)
		err := false
		for _, r := range str {
			if err {
				break
			}
			open := false
			for _, o := range opening {
				if r == o {
					open = true
					break
				}
			}
			if open {
				stack = append(stack, r)
			} else {
				for i, c := range closing {
					if r == c {
						// matching closing bracket
						if opening[i] == stack[len(stack)-1] {
							stack = stack[:len(stack)-1]
							break
						} else {
							// un matched closing bracket
							err = true
						}
					}
				}
			}
		}
		if !err && len(stack) > 0 {
			score := 0
			for len(stack) > 0 {
				score *= 5
				score += valueMap2[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[int(len(scores) / 2)]
}
