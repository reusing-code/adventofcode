package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type jumpFunc func([]int) int

func countStepsFromFile(fileName string, f jumpFunc) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	offsets := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			n, _ := strconv.Atoi(line)
			offsets = append(offsets, n)
		}
	}

	return f(offsets)
}

func countSteps(offsets []int) int {
	idx := 0
	for steps := 0; ; steps++ {
		if idx < 0 || idx >= len(offsets) {
			return steps
		}
		offset := offsets[idx]
		offsets[idx] = offset + 1
		idx = idx + offset
	}
}

func countStepsStrange(offsets []int) int {
	idx := 0
	for steps := 0; ; steps++ {
		if idx < 0 || idx >= len(offsets) {
			return steps
		}
		offset := offsets[idx]
		offsets[idx] = offset + getIncrement(offset)
		idx = idx + offset
	}
}

func getIncrement(a int) int {
	if a >= 3 {
		return -1
	}
	return 1
}
