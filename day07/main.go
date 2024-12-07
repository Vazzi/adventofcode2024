package day07

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 7: Bridge Repair")

	data := utils.ReadLines(utils.InputDataPath(7, false))
	fmt.Println("Result for the first solution is: ", solution(data, false))

	fmt.Println("Result for the second solution is: ", solution(data, true))
}
