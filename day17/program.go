package day17

import (
	"math"
	"strconv"
)

type Program struct {
	rA, rB, rC int
	iPointer   int
}

type Instruction struct {
	opcode  int
	operand int
}

func (p *Program) process(inst Instruction) (output string) {
	output = ""
	switch inst.opcode {
	case 0:
		p.adv(inst.operand)
	case 1:
		p.bxl(inst.operand)
	case 2:
		p.bst(inst.operand)
	case 3:
		if p.jnz(inst.operand) {
			return
		}
	case 4:
		p.bxc()
	case 5:
		output = p.out(inst.operand)
	case 6:
		p.bdv(inst.operand)
	case 7:
		p.cdv(inst.operand)
	}
	p.nextInstruction()
	return

}

func (p *Program) nextInstruction() {
	p.iPointer += 1
}

func (p *Program) bxl(operand int) {
	p.rB = p.rB ^ operand
}

func (p *Program) bst(operand int) {
	p.rB = p.getValueFromComboOperand(operand) % 8
}

func (p *Program) jnz(operand int) bool {
	if p.rA == 0 {
		return false
	}
	p.iPointer = operand / 2
	return true
}

func (p *Program) bxc() {
	p.rB = p.rB ^ p.rC
}

func (p *Program) out(operand int) string {
	return strconv.Itoa(p.getValueFromComboOperand(operand) % 8)
}

func (p *Program) adv(operand int) {
	p.rA = adv(p.rA, p.getValueFromComboOperand(operand))
}

func (p *Program) bdv(operand int) {
	p.rB = adv(p.rA, p.getValueFromComboOperand(operand))
}
func (p *Program) cdv(operand int) {
	p.rC = adv(p.rA, p.getValueFromComboOperand(operand))
}

func adv(numerator, denominator int) int {
	return int(math.Floor(float64(numerator) / math.Pow(2.0, float64(denominator))))
}

func (p *Program) getValueFromComboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return p.rA
	case 5:
		return p.rB
	case 6:
		return p.rC
	}
	return 0
}
