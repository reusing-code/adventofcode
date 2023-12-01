package main

import "strings"

func puzzle1(input []string) int {
	val := 0
	for _, l := range input {
		first := -1
		last := -1
		for _, c := range l {
			if c >= '0' && c <= '9' {
				n := int(c - '0')
				last = n
				if first < 0 {
					first = n
				}
			}
		}
		val += first*10 + last
	}
	return val
}

var numbers = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func puzzle2(input []string) int {
	val := 0
	for _, l := range input {
		first := -1
		last := -1
		for i := 0; i < len(l); i++ {
			c := l[i]
			if c >= '0' && c <= '9' {
				n := int(c - '0')
				last = n
				if first < 0 {
					first = n
				}
			} else {
				for j, numStr := range numbers {
					if strings.HasPrefix(l[i:], numStr) {
						last = j + 1
						if first < 0 {
							first = j + 1
						}
						break
					}
				}
			}

		}
		val += first*10 + last
	}
	return val
}
