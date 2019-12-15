package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcModuleFuel(mass int) int {
	result := mass / 3
	result -= 2
	return result
}

func calcModuleFuelAdvanced(mass int) int {
	result := mass / 3
	result -= 2
	if result <= 0 {
		return 0
	}
	return result + calcModuleFuelAdvanced(result)
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		sum1 += calcModuleFuel(input)
		sum2 += calcModuleFuelAdvanced(input)

	}
	fmt.Printf("Result1: %v\n", sum1)
	fmt.Printf("Result2: %v\n", sum2)
}
