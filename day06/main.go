package day06

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 6: Guard Gallivant")

	mapData := utils.ReadChars(utils.InputDataPath(6, false))
	fmt.Println("Result for the first solution is: ", firstSolution(mapData))

	mapData = utils.ReadChars(utils.InputDataPath(6, false))
	fmt.Println("Result for the second solution is: ", secondSolution(mapData))
}

func outOfBounds(x, y, width, height int) bool {
	return x < 0 || x >= width || y < 0 || y >= height
}
