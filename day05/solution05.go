package day05

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	rules := readRules("./day05/rules02.txt")
	data := utils.ReadWords("./day05/data02.txt", ",")

	fmt.Println("Result for the first solution is: ", firstSolution(rules, data))
	fmt.Println("Result for the second solution is: ", secondSolution(rules, data))
}
