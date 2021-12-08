package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func diffOneRune(a string, b string) rune {
	long := []rune(a)
	short := []rune(b)
	if len(long) < len(short) {
		long = []rune(b)
		short = []rune(a)
	}
	for i, r := range short {
		if r != long[i] {
			return long[i]
		}
	}
	return long[len(long) - 1]
}

func contains(long string, short string) bool {
	l := []rune(long)
	s := []rune(short)

	for _, r1 := range s {
		cont := false
		for _, r2 := range l {
			if r1 == r2 {
				cont = true
			}
		}
		if !cont {
			return false
		}
	}
	return true
}

func puzzle1(input []string) int {
	count := 0
	for _, str := range input {
		pattern := make([]string, 10)
		output := make([]string, 4)
		n, err := fmt.Sscanf(str, "%v %v %v %v %v %v %v %v %v %v | %v %v %v %v", &pattern[0], &pattern[1], &pattern[2], 
		&pattern[3], &pattern[4], &pattern[5], &pattern[6], &pattern[7], &pattern[8], &pattern[9], 
		&output[0], &output[1], &output[2], &output[3])
		if err != nil {
			panic(err)
		}
		if n != 14 {
			panic(n)
		}
		sort.Slice(pattern, func(i, j int) bool {
			return len(pattern[i]) < len(pattern[j])
		})

		for _, o := range output {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				count++
			}
		}
	}
	return count
}

func puzzle2(input []string) int {
	count := 0
	for _, str := range input {
		pattern := make([]string, 10)
		output := make([]string, 4)
		n, err := fmt.Sscanf(str, "%v %v %v %v %v %v %v %v %v %v | %v %v %v %v", &pattern[0], &pattern[1], &pattern[2], 
		&pattern[3], &pattern[4], &pattern[5], &pattern[6], &pattern[7], &pattern[8], &pattern[9], 
		&output[0], &output[1], &output[2], &output[3])
		if err != nil {
			panic(err)
		}
		if n != 14 {
			panic(n)
		}
		sort.Slice(pattern, func(i, j int) bool {
			return len(pattern[i]) < len(pattern[j])
		})
		for i, p := range pattern {
			pattern[i] = SortString(p)
		}
		for i, o := range output {
			output[i] = SortString(o)
		}

		mapping := make(map[string]int, 0)
		mapping[pattern[0]] = 1
		mapping[pattern[1]] = 7
		mapping[pattern[2]] = 4
		mapping[pattern[9]] = 8

		reverseMapping := make(map[int]int, 0)
		reverseMapping[1] = 0
		reverseMapping[7] = 1
		reverseMapping[4] = 2
		reverseMapping[8] = 9

		signalMapping := make(map[rune]rune, 0)
		reverseSignalMapping := make(map[rune]rune, 0)

		// diff between 1 and 7 is signal a
		{
			a := diffOneRune(pattern[reverseMapping[1]], pattern[reverseMapping[7]])
			signalMapping[a] = 'a'
			reverseSignalMapping['a'] = a
		}

		possibleZero := []int{6, 7, 8}
		for i := 6; i <= 8; i++ {
			// the only 6-segment number that contains everything from 4 is 9
			if contains(pattern[i], pattern[reverseMapping[4]]) {
				mapping[pattern[i]] = 9
				reverseMapping[9] = i

				// diff between 9 and 4 is a and g. We already know a
				fourStr := pattern[reverseMapping[4]]
				fourStr += string(reverseSignalMapping['a'])
				fourStr =  SortString(fourStr)
				g := diffOneRune(fourStr, pattern[reverseMapping[9]])
				signalMapping[g] = 'g'
				reverseSignalMapping['g'] = g

				for j, pz := range possibleZero {
					if i == pz {
						possibleZero = append(possibleZero[:j], possibleZero[j+1:]...)
					}
				}
			}
			// the only 6-segment number that does not contain 1 is 6
			if !contains(pattern[i], pattern[reverseMapping[1]]) {
				mapping[pattern[i]] = 6
				reverseMapping[6] = i
				for j, pz := range possibleZero {
					if i == pz {
						possibleZero = append(possibleZero[:j], possibleZero[j+1:]...)
					}
				}
			}
		}

		if len(possibleZero) != 1 {
			panic("Possiblezero should be unique")
		}
		mapping[pattern[possibleZero[0]]] = 0
		reverseMapping[0] = possibleZero[0]

		{
			d := diffOneRune(pattern[reverseMapping[8]], pattern[reverseMapping[0]])
			signalMapping[d] = 'd'
			reverseSignalMapping['d'] = d
		}
		possibleTwo := []int{3, 4, 5}
		for i := 3; i <= 5; i++ {
			// the only 5-segment number that does contain 1 is 3
			if contains(pattern[i], pattern[reverseMapping[1]]) {
				mapping[pattern[i]] = 3
				reverseMapping[3] = i
				for j, pt := range possibleTwo {
					if i == pt {
						possibleTwo = append(possibleTwo[:j], possibleTwo[j+1:]...)
					}
				}
			}
			// the only 5-segment number that is contained in 6 is 5
			if contains(pattern[reverseMapping[6]], pattern[i]) {
				mapping[pattern[i]] = 5
				reverseMapping[5] = i
				for j, pt := range possibleTwo {
					if i == pt {
						possibleTwo = append(possibleTwo[:j], possibleTwo[j+1:]...)
					}
				}
			}
		}

		if len(possibleTwo) != 1 {
			panic("possibleTwo should be unique")
		}
		mapping[pattern[possibleTwo[0]]] = 2
		reverseMapping[2] = possibleTwo[0]


		val := 0
		for _, o := range output {
			val *= 10
			val += mapping[o]
		}
		count += val
	}
	return count
}
