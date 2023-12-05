package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	filename := ""
	dir := os.Getenv("AOC_INPUT_DIR")
	if len(dir) > 0 {
		filename = path.Join(dir, "day05.txt")
	}
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if len(filename) == 0 {
		fmt.Printf("No input file specified. Either specify path to input file as the first\n" +
			"(and only one) command line parameter, or set the env variable AOC_INPUT_DIR to a\n" +
			"directory that contains a file named 'day05.txt'")
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

	start1 := time.Now()
	res1 := puzzle1(input)
	elapsed1 := time.Since(start1)
	start2 := time.Now()
	res2 := puzzle2(input)
	elapsed2 := time.Since(start2)

	fmt.Printf("Puzzle1: %v\n", res1)
	fmt.Printf("Puzzle2: %v\n", res2)
	fmt.Printf("Runtime Puzzle1: %v\n", elapsed1)
	fmt.Printf("Runtime Puzzle2: %v\n", elapsed2)

}
