package day14

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 14: Restroom Redoubt")

	fileName := utils.InputDataPath(14, false)

	fmt.Println("Result for the first solution is: ", firstSolution(fileName, 101, 103))

	//fmt.Println("Result for the second solution is: ", secondSolution(utils.InputDataPath(14, false)))
}
