package main

import (
	"fmt"
	"math"

	"github.com/reusing-code/adventofcode/gohelpers"
)

func poly(input []string, steps int) int {
	split := gohelpers.SplitByEmptyLine(input)
	polymer := split[0][0]

	rules := make(map[string]rune)
	for _, rule := range split[1] {
		var left string
		var right rune
		_, err := fmt.Sscanf(rule, "%v -> %c", &left, &right)
		if err != nil {
			panic(err)
		}
		rules[left] = right
	}

	counts := make(map[rune]int)
	for _, v := range polymer {
		counts[v]++
	}

	for i := 0; i < steps; i++ {
		for j := 0; j < len(polymer)-1; j++ {
			if v, ok := rules[polymer[j:j+2]]; ok {
				polymer = polymer[:j+1] + string(v) + polymer[j+1:]
				counts[v]++
				j++
			}
		}
	}

	max := 0
	min := math.MaxInt
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}

func puzzle1(input []string) int {
	return poly(input, 10)
}

type newPairs struct {
	p1, p2 string
}

func puzzle2(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	polymer := split[0][0]

	pairs := make(map[string]int)
	for i := 0; i < len(polymer)-1; i++ {
		pairs[polymer[i:i+2]]++
	}

	rules := make(map[string]newPairs)
	for _, rule := range split[1] {
		var left string
		var right rune
		_, err := fmt.Sscanf(rule, "%v -> %c", &left, &right)
		if err != nil {
			panic(err)
		}
		rules[left] = newPairs{left[:1] + string(right), string(right) + left[1:]}
	}

	for i := 0; i < 40; i++ {
		origPairs := pairs
		pairs = make(map[string]int)
		for k, v := range origPairs {
			np := rules[k]
			pairs[np.p1] += v
			pairs[np.p2] += v
		}
	}

	counts := make(map[rune]int)
	for k, v := range pairs {
		counts[rune(k[0])] += v
	}
	counts[rune(polymer[len(polymer)-1])]++

	max := 0
	min := math.MaxInt
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}
