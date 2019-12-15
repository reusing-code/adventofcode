package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func program(input []int, noun int, verb int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = -1
		}
	}()
	data := make([]int, len(input))
	copy(data, input)
	data[1] = noun
	data[2] = verb
	data = intcode(data)
	return data[0]
}

func intcode(input []int) []int {
	for idx := 0; true; idx += 4 {
		switch input[idx] {
		case 99:
			return input
		case 1:
			input[input[idx+3]] = input[input[idx+1]] + input[input[idx+2]]
		case 2:
			input[input[idx+3]] = input[input[idx+1]] * input[input[idx+2]]
		}
	}
	return input
}

func findInputs(data []int, want int) int {

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			res := program(data, noun, verb)
			if res == want {
				return noun*100 + verb
			}
		}
	}

	return -1
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		data := make([]int, len(split))
		for i, v := range split {
			in, _ := strconv.Atoi(v)
			data[i] = in
		}
		want := 19690720
		result := findInputs(data, want)
		fmt.Printf("Result %v\n", result)
	}

}
