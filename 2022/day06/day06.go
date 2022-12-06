package main

func puzzle(input []string, length int) int {
	buffer := make([]byte, length)
	idx := 0
	for i, c := range input[0] {
		buffer[idx] = byte(c)
		idx = (idx + 1) % length
		if i >= length-1 {
			different := true
			seen := make(map[byte]bool)
			for _, d := range buffer {
				if _, prs := seen[d]; prs {
					different = false
				}
				seen[d] = true
			}
			if different {
				return i + 1
			}
		}
	}
	return 0
}

func puzzle1(input []string) int {
	return puzzle(input, 4)
}

func puzzle2(input []string) int {
	return puzzle(input, 14)
}
