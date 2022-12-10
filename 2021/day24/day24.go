package main

import (
	"fmt"
	"strconv"
	"strings"
)

func run(program []string, input int, z int) int {
	x := 0
	y := 0
	w := 0
	for _, str := range program {
		if str == "inp w" {
			w = input
			continue
		}
		a := 0
		b := 0
		split := strings.Split(str, " ")
		if split[1] == "x" {
			a = x
		} else if split[1] == "y" {
			a = y
		} else if split[1] == "z" {
			a = z
		} else if split[1] == "w" {
			a = w
		} else {
			panic(split[1])
		}
		if split[2] == "x" {
			b = x
		} else if split[2] == "y" {
			b = y
		} else if split[2] == "z" {
			b = z
		} else if split[2] == "w" {
			b = w
		} else {
			b, _ = strconv.Atoi(split[2])
		}

		if split[0] == "add" {
			a = a + b
		} else if split[0] == "mul" {
			a = a * b
		} else if split[0] == "div" {
			a = int(a / b)
		} else if split[0] == "mod" {
			a = a % b
		} else if split[0] == "eql" {
			if a == b {
				a = 1
			} else {
				a = 0
			}
		} else {
			panic(split[0])
		}

		if split[1] == "x" {
			x = a
		} else if split[1] == "y" {
			y = a
		} else if split[1] == "z" {
			z = a
		} else if split[1] == "w" {
			w = a
		} else {
			panic(split[1])
		}

	}
	return z
}

func SplitByInp(in []string) [][]string {
	result := make([][]string, 0)
	current := make([]string, 0)
	for _, str := range in {
		if str == "inp w" {
			if len(current) > 0 {
				result = append(result, current)
				current = make([]string, 0)
			}
		}

		current = append(current, str)
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}

func puzzle1(input []string) int {
	in := SplitByInp(input)
	possibleZ := make([]map[int]bool, 0)
	possibleZ = append(possibleZ, make(map[int]bool, 0))
	possibleZ[0][0] = true
	for i, prog := range in {
		nextPossibleZ := make(map[int]bool, 0)
		for z, _ := range possibleZ[i] {
			for w := 1; w <= 9; w++ {
				res := run(prog, w, z)
				nextPossibleZ[res] = true

			}
		}
		fmt.Println(i, len(nextPossibleZ))

		possibleZ = append(possibleZ, nextPossibleZ)
	}

	return 1
}

func puzzle2(input []string) int {
	return 1
}
