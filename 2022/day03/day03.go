package main

func calcPriority(c rune) int {
	if c >= 97 {
		return int(c) - 96
	}
	return int(c) - 65 + 27
}

func calcRSValue(in string) int {
	comp1 := in[:len(in)/2]
	comp2 := in[len(in)/2:]

	c1map := make(map[byte]bool)
	for _, c := range comp1 {
		c1map[byte(c)] = true
	}

	for _, c := range comp2 {
		if _, contained := c1map[byte(c)]; contained {
			return calcPriority(c)
		}
	}
	return 0
}

func puzzle1(input []string) int {
	sum := 0
	for _, rs := range input {
		sum += calcRSValue(rs)
	}
	return sum
}

func calcGroupValue(input []string) int {
	c1map := make(map[byte]bool)
	for _, c := range input[0] {
		c1map[byte(c)] = true
	}
	c2map := make(map[byte]bool)
	for _, c := range input[1] {
		if _, contained := c1map[byte(c)]; contained {
			c2map[byte(c)] = true
		}
	}
	for _, c := range input[2] {
		if _, contained := c2map[byte(c)]; contained {
			return calcPriority(c)
		}
	}
	return 0
}

func puzzle2(input []string) int {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		sum += calcGroupValue(input[i : i+3])
	}
	return sum
}
