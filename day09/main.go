package day09

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 9: Disk Fragmenter")

	data := utils.ReadFile(utils.InputDataPath(9, false))
	fmt.Println("Result for the first solution is: ", firstSolution(data))

	//fmt.Println("Result for the second solution is: ", secondSolution(secondData))

}
