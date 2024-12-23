package day17

import (
	"adventofcode2024/utils"
	"strconv"
	"strings"
)

func firstSolution(fileName string) string {
	output := ""
	prog, instructions := readProgram(fileName)

	for prog.iPointer < len(instructions) {
		inst := instructions[prog.iPointer]
		tmpOut := prog.process(inst)
		if tmpOut != "" {
			output += tmpOut + ","
		}
	}

	if len(output) > 0 { // Remove last comma
		output = output[:len(output)-1]
	}
	return output
}

func readProgram(fileName string) (*Program, []Instruction) {
	inputLines := utils.ReadLines(fileName)

	regA, _ := strconv.Atoi(processInputLine(inputLines[0])) // Register A
	regB, _ := strconv.Atoi(processInputLine(inputLines[1])) // Register B
	regC, _ := strconv.Atoi(processInputLine(inputLines[2])) // Register C

	stringInstrs := processInputLine(inputLines[4]) // Instructions
	var instructions []Instruction
	for i := 0; i < len(stringInstrs); i += 4 {
		opcode := int(stringInstrs[i] - '0')
		operand := int(stringInstrs[i+2] - '0')

		instructions = append(instructions, Instruction{opcode, operand})
	}

	newProgram := &Program{rA: regA, rB: regB, rC: regC}
	return newProgram, instructions
}

func processInputLine(line string) string {
	splitLine := strings.Split(line, ": ")
	return splitLine[1]
}
