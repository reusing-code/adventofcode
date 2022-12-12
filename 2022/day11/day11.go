package main

import (
	"fmt"
	"strings"

	"github.com/reusing-code/adventofcode/gohelpers"
)

type (
	operation func(int) int
)

type monkey struct {
	items       []int
	op          operation
	testVal     int
	trueTarget  int
	falseTarget int
	inspects    int
}

func parseMonkey(in []string) *monkey {
	itemsString := strings.Split(strings.Split(in[1], ":")[1], ",")
	items := make([]int, len(itemsString))
	for i := range itemsString {
		fmt.Sscanf(itemsString[i], "%d", &items[i])
	}
	var op operation
	if in[2] == "  Operation: new = old * old" {
		op = func(old int) int {
			return old * old
		}
	} else if strings.HasPrefix(in[2], "  Operation: new = old * ") {
		val := 0
		fmt.Sscanf(in[2], "  Operation: new = old * %d", &val)
		op = func(old int) int {
			return old * val
		}
	} else if strings.HasPrefix(in[2], "  Operation: new = old + ") {
		val := 0
		fmt.Sscanf(in[2], "  Operation: new = old + %d", &val)
		op = func(old int) int {
			return old + val
		}
	} else {
		panic("")
	}
	testVal := 0
	fmt.Sscanf(in[3], "  Test: divisible by %d", &testVal)
	trueTo := 0
	falseTo := 0
	fmt.Sscanf(in[4], "    If true: throw to monkey %d", &trueTo)
	fmt.Sscanf(in[5], "    If false: throw to monkey %d", &falseTo)
	return &monkey{items, op, testVal, trueTo, falseTo, 0}
}

func doTurn(monkey *monkey, monkeys *[]*monkey, worryDivider int, minMultiple int) {
	for _, val := range monkey.items {
		monkey.inspects++
		val = monkey.op(val) / worryDivider
		if minMultiple > 0 {
			val = val % minMultiple
		}
		target := 0
		if val%monkey.testVal == 0 {
			target = monkey.trueTarget
		} else {
			target = monkey.falseTarget
		}
		(*monkeys)[target].items = append((*monkeys)[target].items, val)
	}
	monkey.items = make([]int, 0)
}

func puzzle1(input []string) int {
	monkeyLines := gohelpers.SplitByEmptyLine(input)
	monkeys := make([]*monkey, 0, len(monkeyLines))
	for _, m := range monkeyLines {
		monkeys = append(monkeys, parseMonkey(m))
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			doTurn(monkey, &monkeys, 3, -1)
		}
	}
	max1 := 0
	max2 := 0
	for _, m := range monkeys {
		if m.inspects > max1 {
			max2 = max1
			max1 = m.inspects
		} else if m.inspects > max2 {
			max2 = m.inspects
		}
	}
	return max1 * max2
}

func puzzle2(input []string) int {
	monkeyLines := gohelpers.SplitByEmptyLine(input)
	monkeys := make([]*monkey, 0, len(monkeyLines))
	minMultiple := 1
	for _, m := range monkeyLines {
		monkeys = append(monkeys, parseMonkey(m))
		minMultiple *= monkeys[len(monkeys)-1].testVal
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			doTurn(monkey, &monkeys, 1, minMultiple)
		}
	}
	max1 := 0
	max2 := 0
	for _, m := range monkeys {
		if m.inspects > max1 {
			max2 = max1
			max1 = m.inspects
		} else if m.inspects > max2 {
			max2 = m.inspects
		}
	}
	return max1 * max2
}
