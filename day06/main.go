package day06

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	mapData := utils.ReadChars("./day06/input.txt")
	fmt.Println("Result for the first solution is: ", firstSolution(mapData))

	mapData = utils.ReadChars("./day06/input.txt")
	fmt.Println("Result for the second solution is: ", secondSolution(mapData))
}

func outOfBounds(x, y, width, height int) bool {
	return x < 0 || x >= width || y < 0 || y >= height
}
