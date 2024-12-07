package day05

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	utils.PrintDayHeadline("Day 5: Print Queue")

	rules := readRules("./inputData/day05-rules.txt")
	data := utils.ReadWords("./inputData/day05-data.txt", ",")

	fmt.Println("Result for the first solution is: ", firstSolution(rules, data))
	fmt.Println("Result for the second solution is: ", secondSolution(rules, data))
}
