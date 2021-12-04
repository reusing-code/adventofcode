package gohelpers

import "strconv"

func ParseIntVec(in []string) []int {
	result := make([]int, len(in))
	for i, str := range in {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		result[i] = v
	}
	return result
}

func SplitByEmptyLine(in []string) [][]string {
	result := make([][]string, 0)
	current := make([]string, 0)
	for _, str := range in {
		if len(str) != 0 {
			current = append(current, str)
		} else {
			if len(current) > 0 {
				result = append(result, current)
				current = make([]string, 0)
			}
		}
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}
