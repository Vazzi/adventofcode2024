package day15

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 15: Warehouse Woes")

	mapPath := "./inputData/day15-map.txt"
	movesPath := "./inputData/day15-moves.txt"
	fmt.Println("Result for the first solution is: ", firstSolution(mapPath, movesPath))
}
