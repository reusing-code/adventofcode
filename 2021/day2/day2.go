package main

import (
	"strconv"
	"strings"
)

func puzzle1(input []string) int {
	depth := 0
	hor := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			continue
		}
		val, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		if split[0] == "forward" {
			hor += val
		} else if split[0] == "up" {
			depth -= val
		} else if split[0] == "down" {
			depth += val
		} else {
			panic(split[0])
		}
	}
	return depth * hor
}

func puzzle2(input []string) int {
	depth := 0
	hor := 0
	aim := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			continue
		}
		val, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		if split[0] == "forward" {
			hor += val
			depth += aim * val
		} else if split[0] == "up" {
			aim -= val
		} else if split[0] == "down" {
			aim += val
		} else {
			panic(split[0])
		}
	}
	return depth * hor
}
