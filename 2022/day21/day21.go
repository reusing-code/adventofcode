package main

import (
	"fmt"
	"strconv"
	"strings"
)

type opType int

const (
	add opType = iota
	sub
	mul
	div
	none
	check
)

type monkey struct {
	value     int
	op        opType
	operand1  string
	operand2  string
	valCached bool
	humn      bool
}

func (m *monkey) getValue(monkeys map[string]*monkey) int {
	if m.op == none || m.valCached {
		return m.value
	}
	val1 := monkeys[m.operand1].getValue(monkeys)
	val2 := monkeys[m.operand2].getValue(monkeys)
	switch m.op {
	case add:
		m.value = val1 + val2
	case sub:
		m.value = val1 - val2
	case mul:
		m.value = val1 * val2
	case div:
		m.value = val1 / val2
	}
	m.valCached = true
	return m.value
}

func (m *monkey) getValue2(monkeys map[string]*monkey) (int, string) {
	if m.humn {
		return 0, "x"
	}
	if m.op == none || m.valCached {
		return m.value, ""
	}
	val1, str1 := monkeys[m.operand1].getValue2(monkeys)
	val2, str2 := monkeys[m.operand2].getValue2(monkeys)
	if str1 == "" && str2 == "" {
		switch m.op {
		case add:
			m.value = val1 + val2
		case sub:
			m.value = val1 - val2
		case mul:
			m.value = val1 * val2
		case div:
			m.value = val1 / val2
		}
		m.valCached = true
		return m.value, ""
	} else {
		opStr := ""
		parenthesis := true
		switch m.op {
		case add:
			opStr = "+"
		case sub:
			opStr = "-"
		case mul:
			opStr = "*"
			parenthesis = false
		case div:
			opStr = "/"
			parenthesis = false
		case check:
			opStr = "="
			parenthesis = false
		}
		op1 := str1
		if op1 == "" {
			op1 = fmt.Sprint(val1)
		}
		op2 := str2
		if op2 == "" {
			op2 = fmt.Sprint(val2)
		}
		if parenthesis {
			return 0, fmt.Sprintf("(%v%v%v)", op1, opStr, op2)
		} else {
			return 0, fmt.Sprintf("%v%v%v", op1, opStr, op2)
		}
	}
}

func parseMonkey(in string) (string, *monkey) {
	m := &monkey{}
	split := strings.Split(in, ": ")
	name := split[0]
	sep := ""
	if strings.Contains(split[1], "+") {
		m.op = add
		sep = " + "
	} else if strings.Contains(split[1], "-") {
		m.op = sub
		sep = " - "
	} else if strings.Contains(split[1], "*") {
		m.op = mul
		sep = " * "
	} else if strings.Contains(split[1], "/") {
		m.op = div
		sep = " / "
	} else {
		m.op = none
		m.value, _ = strconv.Atoi(split[1])
	}
	if m.op != none {
		split2 := strings.Split(split[1], sep)
		m.operand1 = split2[0]
		m.operand2 = split2[1]
	}
	return name, m
}

func parseMonkeys(input []string) map[string]*monkey {
	result := make(map[string]*monkey)
	for _, str := range input {
		name, m := parseMonkey(str)
		result[name] = m
	}
	return result
}

func puzzle1(input []string) int {
	monkeys := parseMonkeys(input)

	return monkeys["root"].getValue(monkeys)
}

func puzzle2(input []string) int {
	monkeys := parseMonkeys(input)
	monkeys["root"].op = check
	monkeys["humn"].humn = true
	_, str := monkeys["root"].getValue2(monkeys)
	fmt.Println(str)
	return 1
}
