package day16

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 16: Reindeer Maze")

	fmt.Println("Result for the first solution is: ", firstSolution(utils.ReadChars(utils.InputDataPath(16, false))))

}
