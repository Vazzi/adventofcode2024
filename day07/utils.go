package day07

import (
	"strconv"
	"strings"
)

func parseLine(line string) (int, []int) {
	lineSplits := strings.Split(line, ":")

	result, _ := strconv.Atoi(lineSplits[0])
	numberString := lineSplits[1][1:]

	numbers := make([]int, 0)
	numbersInStrings := strings.Split(numberString, " ")

	for _, str := range numbersInStrings {
		number, _ := strconv.Atoi(str)
		numbers = append(numbers, number)
	}

	return result, numbers
}
