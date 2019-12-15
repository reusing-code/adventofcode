package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		for i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func runProgram(data []int, phase, input int) int {
	prog := newProgram(data)
	prog.input = []int{phase, input}
	prog.execute()
	return prog.output[0]
}

func runThrusterSequence(data, p []int, input int) int {
	thrust := input
	for _, v := range p {
		thrust = runProgram(data, v, thrust)
	}
	return thrust
}

func thrusterSequence(data []int) int {
	maxThrust := 0
	orig := []int{0, 1, 2, 3, 4}
	for p := make([]int, 5); p[0] < len(p); nextPerm(p) {
		thrusterConfig := getPerm(orig, p)
		thrust := runThrusterSequence(data, thrusterConfig, 0)

		if thrust > maxThrust {
			maxThrust = thrust
		}
	}

	return maxThrust
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

		result := thrusterSequence(data)
		fmt.Printf("Result %v\n", result)
	}
}
