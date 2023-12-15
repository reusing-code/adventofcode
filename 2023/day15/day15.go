package main

import (
	"strconv"
	"strings"
)

func hash(in string) int {
	hash := 0
	for _, c := range in {
		hash += int(c)
		hash *= 17
		hash = hash % 256
	}
	return hash
}

func puzzle1(input []string) int {
	split := strings.Split(input[0], ",")
	result := 0
	for _, v := range split {
		result += hash(v)
	}
	return result
}

type lens struct {
	label string
	focus int
}

type mapItem struct {
	lenses []lens
}

func puzzle2(input []string) int {
	split := strings.Split(input[0], ",")
	var hashMap [256]mapItem
	for i := range hashMap {
		hashMap[i].lenses = make([]lens, 0)
	}
	for _, v := range split {
		if v[len(v)-1] == '-' {
			l := v[:len(v)-1]
			h := hash(l)
			for i := range hashMap[h].lenses {
				if hashMap[h].lenses[i].label == l {
					hashMap[h].lenses = append(hashMap[h].lenses[:i], hashMap[h].lenses[i+1:]...)
					break
				}
			}
		} else {
			l := v[:len(v)-2]
			h := hash(l)
			foc, _ := strconv.Atoi(v[len(v)-1:])
			found := false
			for i := range hashMap[h].lenses {
				if hashMap[h].lenses[i].label == l {
					hashMap[h].lenses[i].focus = foc
					found = true
					break
				}
			}
			if !found {
				hashMap[h].lenses = append(hashMap[h].lenses, lens{l, foc})
			}
		}
	}
	result := 0
	for i := 0; i < 256; i++ {
		for j, v := range hashMap[i].lenses {
			result += (i + 1) * (j + 1) * v.focus
		}
	}
	return result
}
