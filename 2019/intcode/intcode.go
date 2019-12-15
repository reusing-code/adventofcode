package intcode

type Program struct {
	name   string
	data   []int64
	input  chan int64
	output chan int64
	iptr   int64
	base   int64
}

type instruction func(params int64)

func (p *Program) arith(params int64, op func(int64, int64) int64) {
	mode1 := params % 10
	params /= 10
	mode2 := params % 10
	params /= 10
	result := op(p.param(mode1, p.iptr+1), p.param(mode2, p.iptr+2))
	idx := p.paramIdx(params%10, p.iptr+3)
	p.data[idx] = result
	p.iptr += 4
}

func (p *Program) add(params int64) {
	p.arith(params, func(a, b int64) int64 {
		return a + b
	})
}

func (p *Program) mult(params int64) {
	p.arith(params, func(a, b int64) int64 {
		return a * b
	})
}

func (p *Program) lt(params int64) {
	p.arith(params, func(a, b int64) int64 {
		if a < b {
			return 1
		}
		return 0
	})
}

func (p *Program) eq(params int64) {
	p.arith(params, func(a, b int64) int64 {
		if a == b {
			return 1
		}
		return 0
	})
}

func (p *Program) in(params int64) {

	val := <-p.input
	idx := p.paramIdx(params%10, p.iptr+1)
	p.data[idx] = val
	p.iptr += 2
}

func (p *Program) out(params int64) {
	val := p.param(params%10, p.iptr+1)
	p.output <- val
	p.iptr += 2
}

func (p *Program) jmp(params int64, cmp func(int64) bool) {
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

func (p *Program) jtrue(params int64) {
	p.jmp(params, func(a int64) bool {
		return a != 0
	})
}

func (p *Program) jfalse(params int64) {
	p.jmp(params, func(a int64) bool {
		return a == 0
	})
}

func (p *Program) adjustBase(params int64) {
	val := p.param(params, p.iptr+1)
	p.base += val
	p.iptr += 2
}

func (p *Program) ensureIdx(idx int64) {
	if idx >= int64(len(p.data)) {
		p.data = append(p.data, make([]int64, idx-int64(len(p.data))+1)...)
	}
}

func (p *Program) paramIdx(mode int64, idx int64) int64 {
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

func (p *Program) param(mode int64, idx int64) int64 {
	finalIdx := p.paramIdx(mode, idx)
	return p.data[finalIdx]
}

func NewProgram(name string, data []int64, input, output chan int64) *Program {
	result := &Program{name: name, data: append(make([]int64, 0, len(data)), data...),
		iptr: 0, input: input, output: output, base: 0}
	return result
}

func (p *Program) Execute() {
	for {
		opcode := p.data[p.iptr]
		op := opcode % 100
		params := opcode / 100
		switch op {
		case 99:
			return
		case 1:
			p.add(params)
		case 2:
			p.mult(params)
		case 3:
			p.in(params)
		case 4:
			p.out(params)
		case 5:
			p.jtrue(params)
		case 6:
			p.jfalse(params)
		case 7:
			p.lt(params)
		case 8:
			p.eq(params)
		case 9:
			p.adjustBase(params)
		}
	}
}
