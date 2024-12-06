package day05

import "slices"

func secondSolution(rules map[string][]string, data [][]string) int {
	score := 0

	for _, order := range data {
		if !checkOrder(order, rules) {
			newOrder := fixOrder(order, rules)
			score += getMiddleNumber(newOrder)
		}
	}

	return score
}

func fixOrder(data []string, rules map[string][]string) []string {
	for numberIndex := len(data) - 1; numberIndex >= 0; numberIndex-- {
		number := data[numberIndex]
		for i := numberIndex; i >= 0; i-- {
			if slices.Contains(rules[number], data[i]) {
				data = append(data[:i], append(data[i+1:numberIndex+1], append([]string{data[i]}, data[numberIndex+1:]...)...)...)
				numberIndex = len(data)
				break
			}
		}
	}

	return data
}
