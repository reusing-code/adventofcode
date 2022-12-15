package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func main() {
	filename := ""
	dir := os.Getenv("AOC_INPUT_DIR")
	if len(dir) > 0 {
		filename = path.Join(dir, "day15.txt")
	}
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if len(filename) == 0 {
		fmt.Printf("No input file specified. Either specify path to input file as the first\n" +
			"(and only one) command line parameter, or set the env variable AOC_INPUT_DIR to a\n" +
			"directory that contains a file named 'day15.txt'")
		os.Exit(1)
	}

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

	res1 := puzzle1(input, 2000000)
	res2 := puzzle2(input, 4000000)

	fmt.Printf("Puzzle1: %v\n", res1)
	fmt.Printf("Puzzle2: %v\n", res2)
}
