package day03

import (
	"adventofcode2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Main() {
	utils.PrintDayHeadline("Day 3: Mull It Over")
	fmt.Println("Result for the first solution is: ", firstSolution(input01))
	fmt.Println("Result for the second solution is: ", secondSolution(input01))
}

func secondSolution(input string) int {
	mulReg, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	doReg, _ := regexp.Compile(`do\(\)`)
	dontReg, _ := regexp.Compile(`don't\(\)`)
	doIndex, dontIndex, expIndex := 0, 0, 0

	// Get indexes for all the expressions
	doIndexes := parseIndexes(doReg.FindAllStringSubmatchIndex(input, -1))
	dontIndexes := parseIndexes(dontReg.FindAllStringSubmatchIndex(input, -1))
	expIndexes := parseIndexes(mulReg.FindAllStringSubmatchIndex(input, -1))

	// Get all mul the expressions
	expressions := mulReg.FindAllString(input, -1)

	solution := 0
	canDo := true

	for i := range len(input) {
		if i == doIndexes[doIndex] {
			if doIndex < len(doIndexes)-1 {
				doIndex++
			}
			canDo = true
		}
		if i == dontIndexes[dontIndex] {
			if dontIndex < len(dontIndexes)-1 {
				dontIndex++
			}
			canDo = false
		}
		if expIndex < len(expIndexes) && i == expIndexes[expIndex] {
			if canDo {
				solution += computeMulFromExpression(expressions[expIndex])
			}
			expIndex++
		}
	}

	return solution
}

func firstSolution(input string) int {
	reg, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	match := reg.FindAllString(input, -1)

	solution := 0
	for _, expression := range match {
		solution += computeMulFromExpression(expression)
	}

	return solution
}

func computeMulFromExpression(expression string) int {
	regForNums, _ := regexp.Compile(`\d+`)
	nums := regForNums.FindAllString(expression, -1)
	return computeMul(nums)
}

func computeMul(nums []string) int {
	firstNum, _ := strconv.Atoi(nums[0])
	secondNum, _ := strconv.Atoi(nums[1])

	return firstNum * secondNum
}

func parseIndexes(input [][]int) []int {
	indexes := make([]int, len(input))
	for i, values := range input {
		indexes[i] = values[0]
	}
	return indexes
}
