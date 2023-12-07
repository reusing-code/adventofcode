package main

import (
	"cmp"
	"fmt"
	"slices"
)

var cardMapping = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type hand [5]int

func parseHand(in string) hand {
	var result hand
	for i, c := range in {
		result[i] = cardMapping[c]
	}
	return result
}

func handType(a hand) int {
	counts := make(map[int]int)
	for _, v := range a {
		counts[v] += 1
	}
	maxCount := 0
	secondMaxCount := 0
	for _, v := range counts {
		if v > maxCount {
			secondMaxCount = maxCount
			maxCount = v
			continue
		}
		if v > secondMaxCount {
			secondMaxCount = v
		}
	}
	if maxCount == 5 {
		return 6
	}
	if maxCount == 4 {
		return 5
	}
	if maxCount == 3 {
		if secondMaxCount == 2 {
			return 4
		}
		return 3
	}
	if maxCount == 2 {
		if secondMaxCount == 2 {
			return 2
		}
		return 1
	}
	return 0
}

type handBid struct {
	h   hand
	bid int
}

func compare(a, b handBid) int {
	typeA := handType(a.h)
	typeB := handType(b.h)
	cmpType := cmp.Compare(typeA, typeB)
	if cmpType != 0 {
		return cmpType
	}
	for i, va := range a.h {
		c := cmp.Compare(va, b.h[i])
		if c != 0 {
			return c
		}
	}
	panic("same hands in input")
}

func puzzle1(input []string) int {
	hands := make([]handBid, 0)
	for _, line := range input {
		var handStr string
		var bid int
		fmt.Sscanf(line, "%s %d", &handStr, &bid)
		hands = append(hands, handBid{parseHand(handStr), bid})
	}

	slices.SortFunc(hands, compare)
	result := 0
	for i, v := range hands {
		result += (i + 1) * v.bid
	}
	return result
}

var cardMapping2 = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func parseHand2(in string) hand {
	var result hand
	for i, c := range in {
		result[i] = cardMapping2[c]
	}
	return result
}

func handType2(a hand) int {
	counts := make(map[int]int)
	for _, v := range a {
		counts[v] += 1
	}
	maxCount := 0
	secondMaxCount := 0
	jokers := 0
	for k, v := range counts {
		if k == 1 {
			jokers = v
			continue
		}
		if v > maxCount {
			secondMaxCount = maxCount
			maxCount = v
			continue
		}
		if v > secondMaxCount {
			secondMaxCount = v
		}
	}
	maxCount += jokers
	if maxCount == 5 {
		return 6
	}
	if maxCount == 4 {
		return 5
	}
	if maxCount == 3 {
		if secondMaxCount == 2 {
			return 4
		}
		return 3
	}
	if maxCount == 2 {
		if secondMaxCount == 2 {
			return 2
		}
		return 1
	}
	return 0
}

func compare2(a, b handBid) int {
	typeA := handType2(a.h)
	typeB := handType2(b.h)
	cmpType := cmp.Compare(typeA, typeB)
	if cmpType != 0 {
		return cmpType
	}
	for i, va := range a.h {
		c := cmp.Compare(va, b.h[i])
		if c != 0 {
			return c
		}
	}
	panic("same hands in input")
}

func puzzle2(input []string) int {
	hands := make([]handBid, 0)
	for _, line := range input {
		var handStr string
		var bid int
		fmt.Sscanf(line, "%s %d", &handStr, &bid)
		hands = append(hands, handBid{parseHand2(handStr), bid})
	}

	slices.SortFunc(hands, compare2)
	result := 0
	for i, v := range hands {
		result += (i + 1) * v.bid
	}
	return result
}
