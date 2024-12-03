package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

func Main() {
	fmt.Println("Result of the expression is: ", computeAllMul(input01))
}

func computeAllMul(input string) int {
	reg, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	regForNums, _ := regexp.Compile(`\d+`)

	match := reg.FindAllString(input, -1)
	solution := 0

	for _, expression := range match {
		nums := regForNums.FindAllString(expression, -1)
		firstNum, _ := strconv.Atoi(nums[0])
		secondNum, _ := strconv.Atoi(nums[1])

		solution += firstNum * secondNum
	}

	return solution
}
