package day3

import (
	"fmt"
	"testing"
)

var examples = []struct {
	input  int
	output int
}{
	{1, 0},
	{2, 1},
	{4, 1},
	{9, 2},
	{12, 3},
	{23, 2},
	{1024, 31},
}

func TestExamples(t *testing.T) {
	for _, test := range examples {
		result := getSpiralMemorySteps(test.input)
		if result != test.output {
			t.Errorf("Input '%d' created wrong result %d, expected %d", test.input, result, test.output)
		}
	}
}

const puzzleInput int = 368078

func TestPuzzle(t *testing.T) {
	fmt.Println(getSpiralMemorySteps(puzzleInput))
	fmt.Println(getStressTestValue(puzzleInput))
}

var examples2 = []struct {
	input  int
	output int
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 4},
	{5, 5},
	{6, 10},
	{7, 11},
	{8, 23},
	{9, 25},
	{10, 26},
	{15, 133},
	{17, 147},
	{21, 362},
	{23, 806},
}

func TestExamples2(t *testing.T) {
	for _, test := range examples2 {
		result := getStressTestValueOfCell(test.input)
		if result != test.output {
			t.Errorf("Input '%d' created wrong result %d, expected %d", test.input, result, test.output)
		}
	}
}
