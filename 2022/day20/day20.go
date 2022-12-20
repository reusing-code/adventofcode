package main

import (
	"fmt"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type entry struct {
	val         int
	originalIdx int
}

func swap(list *[]*entry, i, j int) {
	tmp := (*list)[i]
	(*list)[i] = (*list)[j]
	(*list)[j] = tmp
}

func printList(list []*entry) {
	for _, v := range list {
		fmt.Print(v.val, " ")
	}
	fmt.Println("")
}

func puzzle1(input []string) int {
	intVec := gohelpers.ParseIntVec(input)

	list := make([]*entry, len(intVec))
	for i, v := range intVec {
		list[i] = &entry{v, i}
	}

	for i := 0; i < len(list); i++ {
		for j, v := range list {
			if i == v.originalIdx {
				move := v.val % (len(list) - 1)
				if move < 0 {
					move += len(list) - 1
				}
				destination := (move + j) % (len(list) - 1)
				dir := 1
				if destination < j {
					dir = -1
				}
				for k := j; k != destination; k += dir {
					swap(&list, k, k+dir)
				}
				break
			}
		}
	}

	for i, v := range list {
		if v.val == 0 {
			result := list[(i+1000)%len(list)].val
			result += list[(i+2000)%len(list)].val
			result += list[(i+3000)%len(list)].val
			return result
		}
	}

	return 1
}

func puzzle2(input []string) int {
	intVec := gohelpers.ParseIntVec(input)
	for i, v := range intVec {
		intVec[i] = v * 811589153
	}

	list := make([]*entry, len(intVec))
	for i, v := range intVec {
		list[i] = &entry{v, i}
	}

	for rounds := 0; rounds < 10; rounds++ {
		for i := 0; i < len(list); i++ {
			for j, v := range list {
				if i == v.originalIdx {
					move := v.val % (len(list) - 1)
					if move < 0 {
						move += len(list) - 1
					}
					destination := (move + j) % (len(list) - 1)
					dir := 1
					if destination < j {
						dir = -1
					}
					for k := j; k != destination; k += dir {
						swap(&list, k, k+dir)
					}
					break
				}
			}
		}
	}

	for i, v := range list {
		if v.val == 0 {
			result := list[(i+1000)%len(list)].val
			result += list[(i+2000)%len(list)].val
			result += list[(i+3000)%len(list)].val
			return result
		}
	}

	return 1
}
