package day02

import (
	"adventofcode2024/utils"
	"fmt"
	"math"
)

func Main() {
	fmt.Println("--- Day 2: Red-Nosed Reports ---")

	data := utils.ReadInts(utils.InputDataPath(2, false), " ")
	numberOfSafeReports := computeNumberOfSafeReports(data, false)

	fmt.Println("Number of safe reports is  ", numberOfSafeReports)

	numberOfSafeReportsWithToleration := computeNumberOfSafeReports(data, true)

	fmt.Println("Number of safe reports is with toleration of one error is ", numberOfSafeReportsWithToleration)
}

func checkReport(report []int) bool {
	direction := report[0] - report[1]

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]

		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if (diff > 0 && direction < 0) || (diff < 0 && direction > 0) {
			return false
		}
	}
	return true
}

func computeNumberOfSafeReports(data [][]int, tolerateSingleError bool) int {
	score := 0
	for _, report := range data {
		if checkReport(report) {
			score++
		} else if tolerateSingleError {
			for i := 0; i < len(report); i++ {
				modifiedArray := make([]int, 0, len(report)-1)
				modifiedArray = append(modifiedArray, report[:i]...)
				modifiedArray = append(modifiedArray, report[i+1:]...)
				if checkReport(modifiedArray) {
					score++
					break
				}
			}
		}
	}
	return score
}
