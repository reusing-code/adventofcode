package day4

import (
	"fmt"
	"testing"
)

var examples = []struct {
	input  string
	output bool
}{
	{"aa bb cc dd ee", true},
	{"aa bb cc dd aa", false},
	{"aa bb cc dd aaa", true},
}

func TestExamples(t *testing.T) {
	for _, test := range examples {
		result := checkPhrase(test.input)
		if result != test.output {
			t.Errorf("Input '%s' created wrong result %d, expected %d", test.input, result, test.output)
		}
	}
}

func TestPuzzle(t *testing.T) {
	fmt.Println(checkPhrasesFromFile("input.txt"))
	fmt.Println(checkPhrasesAnagramFromFile("input.txt"))
}

var examples2 = []struct {
	input  string
	output bool
}{
	{"abcde fghij", true},
	{"abcde xyz ecdab", false},
	{"a ab abc abd abf abj", true},
	{"iiii oiii ooii oooi oooo", true},
	{"oiii ioii iioi iiio", false},
}

func TestExamples2(t *testing.T) {
	for _, test := range examples2 {
		result := checkPhraseAnagram(test.input)
		if result != test.output {
			t.Errorf("Input '%s' created wrong result %d, expected %d", test.input, result, test.output)
		}
	}
}
