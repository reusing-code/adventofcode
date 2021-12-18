package main

import (
	"fmt"
	"strconv"
)

type SnailFishNumber struct {
	val    int
	left   *SnailFishNumber
	right  *SnailFishNumber
	parent *SnailFishNumber
}

func parseSnalFishNumber(str string) *SnailFishNumber {
	current := &SnailFishNumber{}
	str = str[1 : len(str)-1]
	for _, c := range str {
		if c == '[' {
			next := &SnailFishNumber{}
			next.parent = current
			if current.left == nil {
				current.left = next
			} else {
				current.right = next
			}
			current = next
		} else if c == ']' {
			current = current.parent
		} else if c == ',' {
		} else {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			next := &SnailFishNumber{}
			next.val = n
			next.parent = current
			if current.left == nil {
				current.left = next
			} else {
				current.right = next
			}
		}
	}
	return current
}

func explode(a *SnailFishNumber, depth int) bool {
	if a.left != nil && a.right != nil {
		return explode(a.left, depth+1) || explode(a.right, depth+1)
	} else {
		if depth > 5 {
			addToLeftNeighbor(a.parent, a.parent.left.val)
			addToRightNeighbor(a.parent, a.parent.right.val)
			a.parent.left = nil
			a.parent.right = nil
			a.parent.val = 0
			a.parent = nil
			return true
		} else {
			return false
		}
	}
}

func addToLeftNeighbor(a *SnailFishNumber, val int) {
	if a.parent == nil {
		return
	}
	current := a
	for {
		if current.parent == nil {
			return
		}
		if current.parent.left == current {
			current = current.parent
		} else {
			current = current.parent
			break
		}
	}
	current = current.left
	for {
		if current.right == nil {
			current.val += val
			return
		} else {
			current = current.right
		}
	}
}

func addToRightNeighbor(a *SnailFishNumber, val int) {
	if a.parent == nil {
		return
	}
	current := a
	for {
		if current.parent == nil {
			return
		}
		if current.parent.right == current {
			current = current.parent
		} else {
			current = current.parent
			break
		}
	}
	current = current.right
	for {
		if current.left == nil {
			current.val += val
			return
		} else {
			current = current.left
		}
	}
}

func doSplit(a *SnailFishNumber) bool {
	if a.left == nil && a.right == nil {
		if a.val >= 10 {
			newLeft := int(a.val / 2)
			newRight := a.val - newLeft
			a.val = 0
			a.left = &SnailFishNumber{}
			a.left.val = newLeft
			a.left.parent = a
			a.right = &SnailFishNumber{}
			a.right.val = newRight
			a.right.parent = a
			return true
		} else {
			return false
		}
	} else {
		return doSplit(a.left) || doSplit(a.right)
	}
}

func reduce(a *SnailFishNumber) {
	fullyReduced := false
	for !fullyReduced {
		exploded := false
		split := false

		depth := 1
		exploded = explode(a, depth)

		if !exploded {
			split = doSplit(a)
		}

		if !exploded && !split {
			fullyReduced = true
		}
	}
}

func addNumbers(a, b *SnailFishNumber) *SnailFishNumber {
	result := &SnailFishNumber{}
	result.left = a
	result.right = b
	a.parent = result
	b.parent = result
	reduce(result)
	return result
}

func calcMagnitude(n *SnailFishNumber) int {
	if n == nil {
		return 0
	}
	if n.left != nil && n.right != nil {
		return 3*calcMagnitude(n.left) + 2*calcMagnitude(n.right)
	}
	return n.val
}

func (s *SnailFishNumber) String() string {
	if s != nil {
		if s.left == nil && s.right == nil {
			return fmt.Sprintf("%v", s.val)
		} else {
			return fmt.Sprintf("[%v, %v]", s.left, s.right)
		}
	} else {
		return "nil"
	}
}

func puzzle1(input []string) int {
	numbers := make([]*SnailFishNumber, 0)
	for _, str := range input {
		number := parseSnalFishNumber(str)
		numbers = append(numbers, number)
	}

	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum = addNumbers(sum, numbers[i])
	}
	return calcMagnitude(sum)
}

func puzzle2(input []string) int {
	max := 0
	for i, str := range input {
		for j := i + 1; j < len(input); j++ {
			n1 := parseSnalFishNumber(str)
			n2 := parseSnalFishNumber(input[j])
			sum := calcMagnitude(addNumbers(n1, n2))
			if sum > max {
				max = sum
			}
			n1 = parseSnalFishNumber(str)
			n2 = parseSnalFishNumber(input[j])
			sum = calcMagnitude(addNumbers(n2, n1))
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
