package day12

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 12: Garden Groups")

	data := utils.ReadChars(utils.InputDataPath(12, false))
	fmt.Println("Result for the first solution is: ", firstSolution(data))

	fmt.Println("Result for the second solution is: ", secondSolution(data))

}
