package main

import "github.com/reusing-code/adventofcode/gohelpers"

var (
	snafuDigit2decimal map[byte]int = map[byte]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}
	decimalDigit2snafu map[int]byte = map[int]byte{2: '2', 1: '1', 0: '0', -1: '-', -2: '='}
)

func snafu2decimal(in string) int {
	result := 0
	for _, c := range in {
		result *= 5
		result += snafuDigit2decimal[byte(c)]
	}
	return result
}

func decimal2snafu(in int) string {
	result := ""
	for in > 0 {
		digit := in % 5
		sub := 0
		if digit >= 3 {
			digit = digit - 5
			sub = gohelpers.Abs(digit)
		}
		result += string(decimalDigit2snafu[digit])
		in += sub
		in /= 5
	}
	realResult := ""
	for i := len(result) - 1; i >= 0; i-- {
		realResult += string(result[i])
	}
	return realResult
}

func puzzle1(input []string) string {
	result := 0
	for _, v := range input {
		result += snafu2decimal(v)
	}
	return decimal2snafu(result)
}

func puzzle2(input []string) int {
	return 1
}
