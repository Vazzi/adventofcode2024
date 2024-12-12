package day11

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 11: Plutonian Pebbles")

	data := utils.ReadInts(utils.InputDataPath(11, false), " ")
	fmt.Println("Result for the first solution is: ", firstSolution(data[0]))

	//fmt.Println("Result for the second solution is: ", secondSolution(data))

}
