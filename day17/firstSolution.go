package day17

import (
	"adventofcode2024/utils"
	"strconv"
	"strings"
)

func firstSolution(fileName string) string {
	output := ""
	prog, instStr := readProgram(fileName)
	instructions := readInstructions(instStr)

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

func readProgram(fileName string) (*Program, string) {
	inputLines := utils.ReadLines(fileName)

	regA, _ := strconv.Atoi(processInputLine(inputLines[0])) // Register A
	regB, _ := strconv.Atoi(processInputLine(inputLines[1])) // Register B
	regC, _ := strconv.Atoi(processInputLine(inputLines[2])) // Register C

	stringInstrs := processInputLine(inputLines[4]) // Instructions

	newProgram := &Program{rA: regA, rB: regB, rC: regC}
	return newProgram, stringInstrs
}

func processInputLine(line string) string {
	splitLine := strings.Split(line, ": ")
	return splitLine[1]
}

func readInstructions(input string) []Instruction {
	var instructions []Instruction
	for i := 0; i < len(input); i += 4 {
		opcode := int(input[i] - '0')
		operand := int(input[i+2] - '0')

		instructions = append(instructions, Instruction{opcode, operand})
	}
	return instructions
}
