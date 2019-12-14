package day6

import (
	"fmt"
	"testing"
)

var example = []int{0, 2, 7, 0}

func TestExample(t *testing.T) {
	result := countSteps(example)
	if result != 5 {
		t.Errorf("Input created wrong result %d", result)
	}
}

func TestExample2(t *testing.T) {
	result := countStepsLoop(example)
	if result != 4 {
		t.Errorf("Input created wrong result %d", result)
	}
}

var puzzleInput = []int{4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3}

func TestPuzzle(t *testing.T) {
	fmt.Println(countSteps(puzzleInput))
	fmt.Println(countStepsLoop(puzzleInput))
}
