package main

import (
	"fmt"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type node struct {
	name  string
	left  string
	right string
}

func parseMap(input []string) map[string]node {
	result := make(map[string]node)
	for _, line := range input {
		var name, left, right string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &name, &left, &right)
		result[name] = node{name, left, right}
	}
	return result
}

func puzzle1(input []string) int {
	split := gohelpers.SplitByEmptyLine(input)
	instructions := split[0][0]
	m := parseMap(split[1])
	current := "AAA"
	for i := 0; ; i++ {
		if current == "ZZZ" {
			return i
		}
		step := instructions[i%len(instructions)]
		if step == 'R' {
			current = m[current].right
		} else {
			current = m[current].left
		}
	}
}

type pathLoop struct {
	looplength   int64
	possibleEnds []int64
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int64) int64 {
	return (a / gcd(a, b)) * b
}

func lcmSlice(s []int64) int64 {
	current := s[0]
	for i := 1; i < len(s); i++ {
		current = lcm(current, s[i])
	}
	return current
}

func puzzle2(input []string) int64 {
	split := gohelpers.SplitByEmptyLine(input)
	instructions := split[0][0]
	m := parseMap(split[1])
	starts := make([]string, 0)
	for k := range m {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}
	pathLoops := make([]pathLoop, 0)
	for _, c := range starts {
		current := c
		loops := make(map[string]int64)
		zs := make([]int64, 0)
		looplen := int64(-1)
		for i := int64(0); ; i++ {
			if current[2] == 'Z' {
				zs = append(zs, i)
			}
			if i%int64(len(instructions)) == 0 {
				if v, ok := loops[current]; ok {
					looplen = i - v
					break
				}
				loops[current] = i
			}
			step := instructions[i%int64(len(instructions))]
			if step == 'R' {
				current = m[current].right
			} else {
				current = m[current].left
			}

		}
		pathLoops = append(pathLoops, pathLoop{looplen, zs})
	}
	currentMin := make([]int64, len(pathLoops))
	for i := range currentMin {
		currentMin[i] = pathLoops[i].possibleEnds[0]
	}

	return lcmSlice(currentMin)

	// for {
	// 	currentMax := int64(0)
	// 	allEqual := true
	// 	result := currentMin[0]
	// 	for _, v := range currentMin {
	// 		if v > currentMax {
	// 			currentMax = v
	// 		}
	// 		if v != result {
	// 			allEqual = false
	// 		}
	// 	}
	// 	if allEqual {
	// 		return result
	// 	}

	// 	for i, v := range currentMin {
	// 		if v == currentMax {
	// 			continue
	// 		}
	// 		l := int64(currentMax / pathLoops[i].looplength)
	// 		for _, v := range pathLoops[i].possibleEnds {
	// 			if l*pathLoops[i].looplength+v >= currentMax {
	// 				currentMin[i] = l*pathLoops[i].looplength + v
	// 				break
	// 			}
	// 		}
	// 		if currentMin[i] < currentMax {
	// 			currentMin[i] = (l+1)*pathLoops[i].looplength + pathLoops[i].possibleEnds[0]
	// 		}
	// 	}
	// }
}
