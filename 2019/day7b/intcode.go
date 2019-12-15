package main

type program struct {
	name       string
	data       []int
	input      chan int
	output     chan int
	iptr       int
	lastOutput int
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

func (p *program) lt(params int) {
	p.arith(params, func(a, b int) int {
		if a < b {
			return 1
		}
		return 0
	})
}

func (p *program) eq(params int) {
	p.arith(params, func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	})
}

func (p *program) in(params int) {
	val := <-p.input
	p.data[p.data[p.iptr+1]] = val
	p.iptr += 2
}

func (p *program) out(params int) {
	val := p.param(params%10, p.iptr+1)
	p.output <- val
	p.lastOutput = val
	p.iptr += 2
}

func (p *program) jmp(params int, cmp func(int) bool) {
	mode1 := params % 10
	params /= 10
	mode2 := params % 10
	val := p.param(mode1, p.iptr+1)
	jmpDest := p.param(mode2, p.iptr+2)
	if cmp(val) {
		p.iptr = jmpDest
	} else {
		p.iptr += 3
	}
}

func (p *program) jtrue(params int) {
	p.jmp(params, func(a int) bool {
		return a != 0
	})
}

func (p *program) jfalse(params int) {
	p.jmp(params, func(a int) bool {
		return a == 0
	})
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

func newProgram(name string, data []int) *program {
	result := &program{name: name, data: data, iptr: 0}
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
		case 5:
			inst = p.jtrue
		case 6:
			inst = p.jfalse
		case 7:
			inst = p.lt
		case 8:
			inst = p.eq
		}
		inst(params)
	}
}
