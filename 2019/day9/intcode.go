package main

type program struct {
	name   string
	data   []int64
	input  chan int64
	output chan int64
	iptr   int64
	base   int64
}

type instruction func(params int64)

func (p *program) arith(params int64, op func(int64, int64) int64) {
	mode1 := params % 10
	params /= 10
	mode2 := params % 10
	params /= 10
	result := op(p.param(mode1, p.iptr+1), p.param(mode2, p.iptr+2))
	idx := p.paramIdx(params%10, p.iptr+3)
	p.data[idx] = result
	p.iptr += 4
}

func (p *program) add(params int64) {
	p.arith(params, func(a, b int64) int64 {
		return a + b
	})
}

func (p *program) mult(params int64) {
	p.arith(params, func(a, b int64) int64 {
		return a * b
	})
}

func (p *program) lt(params int64) {
	p.arith(params, func(a, b int64) int64 {
		if a < b {
			return 1
		}
		return 0
	})
}

func (p *program) eq(params int64) {
	p.arith(params, func(a, b int64) int64 {
		if a == b {
			return 1
		}
		return 0
	})
}

func (p *program) in(params int64) {

	val := <-p.input
	idx := p.paramIdx(params%10, p.iptr+1)
	p.data[idx] = val
	p.iptr += 2
}

func (p *program) out(params int64) {
	val := p.param(params%10, p.iptr+1)
	p.output <- val
	p.iptr += 2
}

func (p *program) jmp(params int64, cmp func(int64) bool) {
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

func (p *program) jtrue(params int64) {
	p.jmp(params, func(a int64) bool {
		return a != 0
	})
}

func (p *program) jfalse(params int64) {
	p.jmp(params, func(a int64) bool {
		return a == 0
	})
}

func (p *program) adjustBase(params int64) {
	val := p.param(params, p.iptr+1)
	p.base += val
	p.iptr += 2
}

func (p *program) ensureIdx(idx int64) {
	if idx >= int64(len(p.data)) {
		p.data = append(p.data, make([]int64, idx-int64(len(p.data))+1)...)
	}
}

func (p *program) paramIdx(mode int64, idx int64) int64 {
	var finalIdx int64
	switch mode {
	case 0:
		finalIdx = p.data[idx]
	case 1:
		finalIdx = idx
	case 2:
		finalIdx = p.data[idx] + p.base
	default:
		panic("Unkown mode")
	}
	p.ensureIdx(finalIdx)
	return finalIdx
}

func (p *program) param(mode int64, idx int64) int64 {
	finalIdx := p.paramIdx(mode, idx)
	return p.data[finalIdx]
}

func newProgram(name string, data []int64) *program {
	result := &program{name: name, data: data, iptr: 0, input: make(chan int64), output: make(chan int64), base: 0}
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
		case 9:
			inst = p.adjustBase
		}
		inst(params)
	}
}
