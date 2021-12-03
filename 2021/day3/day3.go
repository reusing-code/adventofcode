package main

import "strconv"

func puzzle1(input []string) int {
	gamma := 0
	epsilon := 0
	for i := 0; i < len(input[0]); i++ {
		gamma *= 2
		epsilon *= 2
		ones := 0
		zeros := 0
		for _, v := range input {
			if v[i] == '0' {
				ones++
			}
			if v[i] == '1' {
				zeros++
			}
		}
		if ones > zeros {
			gamma++
		} else {
			epsilon++
		}
	}
	return gamma * epsilon
}

func calcRating(input []string, most bool) int {
	for i := 0; i < len(input[0]); i++ {
		onesSlice := make([]string, 0)
		zeroesSlice := make([]string, 0)
		for _, v := range input {
			if v[i] == '1' {
				onesSlice = append(onesSlice, v)
			} else {
				zeroesSlice = append(zeroesSlice, v)
			}
		}

		if (len(onesSlice) >= len(zeroesSlice)) == most {
			input = onesSlice
		} else {
			input = zeroesSlice
		}
		if len(input) <= 1 {
			break
		}
	}

	result, err := strconv.ParseInt(input[0], 2, 64)
	if err != nil {
		panic(err)
	}
	return int(result)
}

func puzzle2(input []string) int {
	return calcRating(input, true) * calcRating(input, false)
}
