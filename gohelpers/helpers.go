package gohelpers

import (
	"strconv"
	"strings"
)

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

func ParseIntVecSingleLine(in string) []int {
	result := make([]int, 0, (len(in)/2)+1)
	stringSlice := strings.Split(in, ",")
	for _, str := range stringSlice {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		result = append(result, v)
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

func ParseIntField(in []string) [][]int {
	field := make([][]int, len(in))
	for i, str := range in {
		field[i] = make([]int, len(str))
		for j, s := range str {
			n, err := strconv.Atoi(string(s))
			if err != nil {
				panic(err)
			}
			field[i][j] = n
		}
	}
	return field
}

func ParseCharField(in []string) [][]byte {
	field := make([][]byte, len(in))
	for i, str := range in {
		field[i] = make([]byte, len(str))
		for j, s := range str {
			field[i][j] = byte(s)
		}
	}
	return field
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	if x > 0 {
		return 1
	}
	panic("")
}
