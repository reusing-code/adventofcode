package adventofcode

import "strconv"

func inverseCaptcha(input string) int {
	var res int = 0

	intArr := toIntArray(input)

	for i, val := range intArr {
		nextIdx := i + 1
		if nextIdx >= len(intArr) {
			nextIdx = 0
		}

		if val == intArr[nextIdx] {
			res += val
		}
	}

	return res
}

func toIntArray(input string) []int {
	intArr := make([]int, 0, len(input))
	// byte length not rune length, but who cares...
	for _, r := range input {
		val, _ := strconv.ParseInt(string(r), 10, 0)
		intArr = append(intArr, int(val))
	}
	return intArr
}

func inverseCaptchaCircular(input string) int {
	var res int = 0

	intArr := toIntArray(input)

	for i, val := range intArr {
		cmpIdx := (i + len(intArr) / 2) % len(intArr)

		if val == intArr[cmpIdx] {
			res += val
		}
	}

	return res
}