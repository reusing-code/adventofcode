package main

func puzzle1(input []string) int {
	inputInt := parseIntVec(input)
	result := 0
	first := true
	previous := 0
	for _, v := range inputInt {
		if !first {
			{
				if v > previous {
					result++
				}
			}
		}
		first = false
		previous = v
	}

	return result
}

func puzzle2(input []string) int {
	inputInt := parseIntVec(input)
	result := 0
	for i := 0; i < len(inputInt)-3; i++ {
		if inputInt[i] < inputInt[i+3] {
			result++
		}
	}
	return result
}
