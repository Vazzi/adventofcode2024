package day05

import (
	"slices"
	"strconv"
)

func firstSolution(rules map[string][]string, data [][]string) int {
	score := 0

	for _, order := range data {
		if checkOrder(order, rules) {
			score += getMiddleNumber(order)
		}
	}

	return score
}

func checkOrder(data []string, rules map[string][]string) bool {
	for numberIndex := len(data) - 1; numberIndex >= 0; numberIndex-- {
		number := data[numberIndex]
		for i := numberIndex; i >= 0; i-- {
			if slices.Contains(rules[number], data[i]) {
				return false
			}
		}
	}

	return true
}

func getMiddleNumber(data []string) int {
	result, _ := strconv.Atoi(data[len(data)/2])
	return result
}
