package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/reusing-code/adventofcode/2019/intcode"
)

func nextPerm(p []int64) {
	for i := len(p) - 1; i >= 0; i-- {
		for i == 0 || p[i] < int64(len(p)-i-1) {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int64) []int64 {
	result := append([]int64{}, orig...)
	for i, v := range p {
		result[i], result[int64(i)+v] = result[int64(i)+v], result[i]
	}
	return result
}

func runProgram(data []int64, phase, input int64) int64 {
	inchan := make(chan int64, 2)
	output := make(chan int64, 1)
	prog := intcode.NewProgram("THRUSTER", data, inchan, output)
	go prog.Execute()
	inchan <- phase
	inchan <- input
	return <-output
}

func runThrusterSequence(data, p []int64, input int64) int64 {
	thrust := input
	for _, v := range p {
		thrust = runProgram(data, v, thrust)
	}
	return thrust
}

func thrusterSequence(data []int64) int64 {
	maxThrust := int64(0)
	orig := []int64{0, 1, 2, 3, 4}
	for p := make([]int64, 5); p[0] < int64(len(p)); nextPerm(p) {
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
		data := make([]int64, len(split))
		for i, v := range split {
			in, _ := strconv.ParseInt(v, 10, 64)
			data[i] = in
		}

		result := thrusterSequence(data)
		fmt.Printf("Result %v\n", result)
	}
}
