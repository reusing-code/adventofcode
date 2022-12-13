package main

import (
	"fmt"
	"sort"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type element struct {
	isInt   bool
	val     int
	list    []*element
	divider bool
}

func printElem(e *element) {
	if e.isInt {
		fmt.Print(e.val)
	} else {
		fmt.Print("[")
		for i, v := range e.list {
			printElem(v)
			if i < len(e.list)-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]")
	}
}

func parseString(in string) (*element, string) {
	var result *element
	if in[0] == '[' {
		result = &element{false, 0, make([]*element, 0), false}
		rest := in[1:]
		for rest[0] != ']' {
			if rest[0] == ',' {
				rest = rest[1:]
			}
			child, newRest := parseString(rest)
			rest = newRest
			result.list = append(result.list, child)
		}
		return result, rest[1:]
	} else {
		val := 0
		fmt.Sscanf(in, "%d", &val)
		result = &element{true, val, nil, false}
		if val >= 10 {
			return result, in[2:]
		} else {
			return result, in[1:]
		}
	}
}

type order int

const (
	rightOrder order = 1
	wrongOrder order = -1
	undecided  order = 0
)

func inOrderInt(a int, b int) order {
	if a < b {
		return rightOrder
	} else if a == b {
		return undecided
	} else {
		return wrongOrder
	}
}

func inOrder(a *element, b *element) order {
	if a.isInt && b.isInt {
		return inOrderInt(a.val, b.val)
	} else if !a.isInt && !b.isInt {
		for i := range a.list {
			if len(b.list) < i+1 {
				return wrongOrder
			}
			res := inOrder(a.list[i], b.list[i])
			if res != undecided {
				return res
			}
		}
		return inOrderInt(len(a.list), len(b.list))
	} else {
		if a.isInt {
			a.isInt = false
			a.list = []*element{{true, a.val, nil, false}}
		}
		if b.isInt {
			b.isInt = false
			b.list = []*element{{true, b.val, nil, false}}
		}
		return inOrder(a, b)
	}
}

func less(a *element, b *element) bool {
	order := inOrder(a, b)
	return order == rightOrder
}

func puzzle1(input []string) int {
	pairs := gohelpers.SplitByEmptyLine(input)
	result := 0
	for i, pair := range pairs {
		a, _ := parseString(pair[0])
		b, _ := parseString(pair[1])
		if inOrder(a, b) == rightOrder {
			result += i + 1
		}
	}
	return result
}

func puzzle2(input []string) int {
	elems := make([]*element, 0, len(input))
	for _, str := range input {
		if str == "" {
			continue
		}
		newElem, _ := parseString(str)
		elems = append(elems, newElem)
	}

	divider1, _ := parseString("[[2]]")
	divider1.divider = true
	elems = append(elems, divider1)
	divider2, _ := parseString("[[6]]")
	divider2.divider = true
	elems = append(elems, divider2)

	sort.Slice(elems, func(i, j int) bool {
		return less(elems[i], elems[j])
	})
	result := 1
	for i, val := range elems {
		if val.divider {
			result *= i + 1
		}
	}
	return result
}
