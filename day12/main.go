package day12

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 12: Garden Groups")

	data := utils.ReadChars(utils.InputDataPath(12, false))
	fmt.Println("Result for the first solution is: ", firstSolution(data))

	// Tooooooooooooooo slow need optimisation
	//secondData := utils.ReadInts(utils.InputDataPath(11, false), " ")
	//fmt.Println("Result for the second solution is: ", solution(secondData[0], 75))

}
