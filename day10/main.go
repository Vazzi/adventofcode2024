package day10

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 10: Hoof It")

	data := utils.ReadInts(utils.InputDataPath(10, false), "")
	fmt.Println("Result for the first solution is: ", firstSolution(data))

	//fmt.Println("Result for the second solution is: ", secondSolution(data))
}
