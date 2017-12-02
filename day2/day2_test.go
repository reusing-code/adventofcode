package day2

import (
	"testing"
	"fmt"
)

func TestExample(t *testing.T) {
	res := checksum("test_input.txt")
	if (res != 18) {
		t.Errorf("Wrong result, was %d expected %d", res, 18)
	}

	res = checksumEvenDivisable("test_input_part2.txt")
	if (res != 9) {
		t.Errorf("Wrong result, was %d expected %d", res, 9)
	}
}

func TestRealInput(t *testing.T) {
	fmt.Println(checksum("input.txt"))
	fmt.Println(checksumEvenDivisable("input.txt"))
}


