package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program struct {
	data   []int
	input  []int
	output []int
	iptr   int
}

type instruction func(params int)

func (p *program) arith(params int, op func(int, int) int) {
	mode1 := params % 10
	params /= 10
	mode2 := params % 10
	result := op(p.param(mode1, p.iptr+1), p.param(mode2, p.iptr+2))
	p.data[p.data[p.iptr+3]] = result
	p.iptr += 4
}

func (p *program) add(params int) {
	p.arith(params, func(a, b int) int {
		return a + b
	})
}

func (p *program) mult(params int) {
	p.arith(params, func(a, b int) int {
		return a * b
	})
}

func (p *program) in(params int) {
	val := p.input[0]
	p.input = p.input[1:]
	p.data[p.data[p.iptr+1]] = val
	p.iptr += 2
}

func (p *program) out(params int) {
	val := p.param(params%10, p.iptr+1)
	p.output = append(p.output, val)
	p.iptr += 2
}

func (p *program) param(mode int, idx int) int {
	switch mode {
	case 0:
		return p.data[p.data[idx]]
	case 1:
		return p.data[idx]
	default:
		panic("Unkown mode")
	}
}

func newProgram(data []int) *program {
	result := &program{data: data, input: make([]int, 0), output: make([]int, 0), iptr: 0}
	return result
}

func (p *program) execute() {
	for {
		var inst instruction
		opcode := p.data[p.iptr]
		op := opcode % 100
		params := opcode / 100
		switch op {
		case 99:
			return
		case 1:
			inst = p.add
		case 2:
			inst = p.mult
		case 3:
			inst = p.in
		case 4:
			inst = p.out
		}
		inst(params)
	}
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		data := make([]int, len(split))
		for i, v := range split {
			in, _ := strconv.Atoi(v)
			data[i] = in
		}
		prog := newProgram(data)
		prog.input = append(prog.input, 1)
		prog.execute()

		fmt.Printf("Result %v\n", prog.output)
	}
}
