package day08

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 8: Resonant Collinearity")

	data := utils.ReadChars(utils.InputDataPath(8, false))
	fmt.Println("Result for the first solution is: ", firstSolution(data))

	// fmt.Println("Result for the second solution is: ", secondSolution(data, false))

}
