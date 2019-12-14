package day5

import (
	"fmt"
	"testing"
)

var example = []int{
	0, 3, 0, 1, -3,
}

func TestExample(t *testing.T) {
	result := countSteps(example)
	if result != 5 {
		t.Errorf("Input created wrong result")
	}
}

func TestExample2(t *testing.T) {
	result := countStepsStrange(example)
	if result != 10 {
		t.Errorf("Input created wrong result")
	}
}

func TestPuzzle(t *testing.T) {
	fmt.Println(countStepsFromFile("input.txt", countSteps))
	fmt.Println(countStepsFromFile("input.txt", countStepsStrange))
}
