package day10

func secondSolution(data [][]int) int {
	inputMap := newMapData(data)
	startingPoints := findStartingPoints(inputMap)
	result := 0

	for _, sp := range startingPoints {
		result += len(findAllEnds(sp, inputMap))
	}

	return result
}
