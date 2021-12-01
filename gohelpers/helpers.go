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
