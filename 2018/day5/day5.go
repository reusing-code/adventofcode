package main

func checkReact(a uint8, b uint8) bool {
	if a == b {
		return false
	}
	if a | 32 == b | 32 {
		return true
	} 
	return false
}

func react(poly []uint8) int {
	for i := 0;; {
		if i + 1 >= len(poly) {
			break
		}
		if checkReact(poly[i], poly[i+1]) {
			poly = append(poly[:i], poly[i+2:]...)
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return len(poly)
}

func removeUnit(in []uint8, unit uint8) []uint8 {
	for i := 0; i < len(in); i++ {
		if in[i] | 32 == unit {
			in = append(in[:i], in[i+1:]...)
			i--
		} 
	}
	return in
}

func puzzle1(input []string) int {
	poly := []uint8(input[0])
	return react(poly)
}

func puzzle2(input []string) int {
	shortest := len(input[0])
	for c := uint8(97); c <= 122; c++ {
		poly := []uint8(input[0])
		poly = removeUnit(poly, c)
		n := react(poly)
		if n < shortest {
			shortest = n
		}
	}
	return shortest
}
