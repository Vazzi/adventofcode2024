package day06

func firstSolution(mapData [][]string) int {
	guard := findGuard(mapData)
	width, height := len(mapData[0]), len(mapData)
	steps := 0

	for {
		nextX, nextY := guard.nextStep()

		if outOfBounds(nextX, nextY, width, height) {
			return steps
		}

		if mapData[nextY][nextX] == "#" {
			guard.dir = guard.turnRight()
			continue
		}

		guard.x, guard.y = nextX, nextY
		if mapData[guard.y][guard.x] != "X" {
			mapData[guard.y][guard.x] = "X"
			steps++
		}
	}
}
