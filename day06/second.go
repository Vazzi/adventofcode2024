package day06

func secondSolution(mapData [][]string) int {
	obstaclePositions := 0
	guard := findGuard(mapData)
	mapData = markVisitsOnMap(mapData, guard)

	for y := range mapData {
		for x := range mapData[y] {
			if mapData[y][x] == "X" && (x != guard.x || y != guard.y) {
				mapData[y][x] = "#" // Place fake obstacle
				if checkCycle(mapData, guard) {
					obstaclePositions++
				}
				mapData[y][x] = "." // Remove fake obstacle
			}
		}
	}
	return obstaclePositions
}

// Mark where guard was with X
func markVisitsOnMap(mapData [][]string, guard GuardState) [][]string {
	width, height := len(mapData[0]), len(mapData)

	for {
		nextX, nextY := guard.nextStep()

		if outOfBounds(nextX, nextY, width, height) {
			return mapData
		}

		if mapData[nextY][nextX] == "#" {
			guard.dir = guard.turnRight()
			continue
		}
		guard.x, guard.y = nextX, nextY
		mapData[guard.y][guard.x] = "X"
		nextX, nextY = guard.nextStep()
	}
}

// Simulate guard and check if there is a cycle
func checkCycle(mapData [][]string, initGuard GuardState) bool {
	guardStates := make([]GuardState, 0)
	width, height := len(mapData[0]), len(mapData)
	guard := GuardState{initGuard.x, initGuard.y, initGuard.dir}

	for {
		nextX, nextY := guard.nextStep()
		if outOfBounds(nextX, nextY, width, height) {
			return false
		}

		if mapData[nextY][nextX] == "#" {
			guard.dir = guard.turnRight()
			continue
		}

		// Check if guard is in cycle
		// This is really slow, but it works :(
		for _, state := range guardStates {
			if state.x == nextX &&
				state.y == nextY &&
				state.dir == guard.dir {
				return true
			}
		}

		guard.x = nextX
		guard.y = nextY
		guardStates = append(guardStates, GuardState{nextX, nextY, guard.dir})
	}
}
