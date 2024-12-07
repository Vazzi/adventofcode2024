package day07

func firstSolution(data []string) int {
	solution := 0
	for _, line := range data {
		result, numbers := parseLine(line)
		if hasSolution(result, numbers) {
			solution += result
		}
	}
	return solution
}

func hasSolution(result int, numbers []int) bool {
	variations := generateVariations(len(numbers) - 1)

	for _, variation := range variations {
		localResult := numbers[0]
		for i := 0; i < len(numbers)-1; i++ {
			switch variation[i] {
			case Mul:
				localResult = localResult * numbers[i+1]
			case Add:
				localResult = localResult + numbers[i+1]
			}
		}
		if localResult == result {
			return true
		}
	}
	return false
}

const (
	Mul int = 0
	Add     = 1
)

func generateVariations(length int) [][]int {
	operators := []int{Mul, Add}

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
