package day6

import "strconv"

func countSteps(input []int) int {
	usedKeys := make(map[string]bool)
	for steps := 0; ; steps++ {
		key := getKey(input)
		_, exists := usedKeys[key]
		if exists {
			return steps
		}
		usedKeys[key] = true
		input = reallocate(input)
	}

}

func countStepsLoop(input []int) int {
	usedKeys := make(map[string]int)
	for steps := 0; ; steps++ {
		key := getKey(input)
		_, exists := usedKeys[key]
		if exists {
			return steps - usedKeys[key]
		}
		usedKeys[key] = steps
		input = reallocate(input)
	}
}

func getKey(input []int) string {
	res := ""
	for _, v := range input {
		res += strconv.Itoa(v) + ","
	}
	return res
}

func reallocate(in []int) []int {
	curIdx := getMaxIndex(in)
	remainder := in[curIdx]
	in[curIdx] = 0
	for ; remainder > 0; remainder-- {
		curIdx++
		if curIdx == len(in) {
			curIdx = 0
		}
		in[curIdx] += 1

	}
	return in
}

func getMaxIndex(in []int) int {
	max := in[0]
	maxIdx := 0
	for i, v := range in {
		if v > max {
			max = v
			maxIdx = i
		}
	}
	return maxIdx
}
