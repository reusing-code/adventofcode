package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseIntVec(in []string) []int {
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

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	res1 := puzzle1(input)
	res2 := puzzle2(input)

	fmt.Printf("Puzzle1: %v\n", res1)
	fmt.Printf("Puzzle2: %v\n", res2)

}
