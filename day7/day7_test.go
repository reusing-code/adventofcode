package day7

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	result := findRoot("exampleInput.txt")
	if result.name != "tknk" {
		t.Errorf("Input created wrong result %s", result.name)
	}
}

func TestPuzzle(t *testing.T) {
	fmt.Println(findRoot("input.txt").name)
}
