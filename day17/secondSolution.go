package day17

import (
	"strings"
)

func secondSolution(fileName string) int {
	_, initInst := readProgram(fileName)
	instructions := readInstructions(initInst)

	i := 0
	for {
		testProg := Program{rA: i}
		if testProg.computeEqualOutput(instructions, initInst) {
			return i
		}
		i += 1
	}
}

func (p *Program) computeEqualOutput(instructions []Instruction, expectedOutput string) bool {
	output := ""
	expectedOutput += ","
	for p.iPointer < len(instructions) {
		inst := instructions[p.iPointer]
		tmpOut := p.process(inst)
		if tmpOut != "" {
			output += tmpOut + ","
		}

		if !strings.HasPrefix(expectedOutput, output) {
			return false
		}

	}
	return expectedOutput == output

}
