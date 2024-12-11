package day10

type point struct {
	x, y int
}

type mapData struct {
	data          [][]int
	width, height int
}

func newMapData(data [][]int) *mapData {
	return &mapData{
		data,
		len(data[0]),
		len(data),
	}
}

func firstSolution(data [][]int) int {
	inputMap := newMapData(data)
	startingPoints := findStartingPoints(inputMap)
	result := 0

	for _, sp := range startingPoints {
		result += findAllUniquePaths(sp, inputMap)
	}

	return result
}

func findAllEnds(startingPoint point, data *mapData) []point {

	var findEnd (func(int, int, int) []point)
	findEnd = func(x, y, val int) []point {
		ends := make([]point, 0)
		if x < 0 || y < 0 || x > data.width-1 || y > data.height-1 {
			return ends
		}
		if data.data[y][x] != val {
			return ends
		}
		if val == 9 {
			return append(ends, point{x, y})
		}

		ends = append(ends, findEnd(x, y+1, val+1)...)
		ends = append(ends, findEnd(x, y-1, val+1)...)
		ends = append(ends, findEnd(x+1, y, val+1)...)
		ends = append(ends, findEnd(x-1, y, val+1)...)

		return ends
	}

	return findEnd(startingPoint.x, startingPoint.y, 0)
}

func findAllUniquePaths(startingPoint point, data *mapData) int {

	ends := findAllEnds(startingPoint, data)
	uniques := make(map[point]bool)
	for _, p := range ends {
		if !uniques[p] {
			uniques[p] = true
		}
	}
	return len(uniques)

}

func findStartingPoints(data *mapData) []point {
	startingPoints := make([]point, 0)

	for y := range data.height {
		for x := range data.width {
			if data.data[y][x] == 0 {
				startingPoints = append(startingPoints, point{x, y})
			}

		}
	}
	return startingPoints
}
