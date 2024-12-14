package day13

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 13: Claw Contraption")

	fmt.Println("Result for the first solution is: ", firstSolution(utils.InputDataPath(13, false)))

	fmt.Println("Result for the second solution is: ", secondSolution(utils.InputDataPath(13, false)))

}
