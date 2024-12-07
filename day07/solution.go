package day07

import (
	"fmt"
	"strconv"
)

func solution(data []string, useConcat bool) int {
	finalResult := 0
	for _, line := range data {
		result, numbers := parseLine(line)
		if hasSolution(result, numbers, useConcat) {
			finalResult += result
		}
	}
	return finalResult
}

func hasSolution(result int, numbers []int, useConcat bool) bool {
	variations := generateVariations(len(numbers)-1, useConcat)

	for _, variation := range variations {
		localResult := numbers[0]
		for i := 0; i < len(numbers)-1; i++ {
			switch variation[i] {
			case Mul:
				localResult = localResult * numbers[i+1]
			case Add:
				localResult = localResult + numbers[i+1]
			case Concat:
				localResult, _ = concatenate(localResult, numbers[i+1])
			}
		}
		if localResult == result {
			return true
		}
	}
	return false
}

func concatenate(a, b int) (int, error) {
	return strconv.Atoi(fmt.Sprintf("%d%d", a, b))
}

const (
	Mul    int = 0
	Add        = 1
	Concat     = 2
)

func generateVariations(length int, useConcat bool) [][]int {
	operators := []int{Mul, Add}
	if useConcat {
		operators = append(operators, Concat)
	}

	var results [][]int
	var compute func([]int)

	compute = func(current []int) {
		if len(current) == length {
			results = append(results, append([]int{}, current...))
			return
		}

		for _, val := range operators {
			compute(append(current, val))
		}
	}

	compute([]int{})
	return results
}
