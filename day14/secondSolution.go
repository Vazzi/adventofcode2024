package day14

import (
	"os"
)

func secondSolution(fileName string, spaceWidth, spaceHeight int) int {
	robots := readRobotInputs(fileName)

	for i := 1; ; i++ {
		for ri, r := range robots {
			r.move(spaceWidth, spaceHeight)
			robots[ri] = r
		}

		robotsPositions, overlaps := doesAnyRobotOverlaps(&robots)
		// If no robots overlap, return the number of seconds
		if !overlaps {
			// Check if there is a Christmas tree
			printRobotsPositionInMap(robotsPositions, spaceWidth, spaceHeight)
			return i
		}
	}
}

type positionsMap map[position]bool

func doesAnyRobotOverlaps(robots *[]robot) (*positionsMap, bool) {
	positions := make(positionsMap, len(*robots))

	for _, r := range *robots {
		if _, ok := positions[r.p]; ok {
			return &positions, true
		}
		positions[r.p] = true
	}

	return &positions, false
}

func printRobotsPositionInMap(robots *positionsMap, spaceWidth, spaceHeight int) {
	m := make([][]bool, spaceHeight)

	for y := range m {
		m[y] = make([]bool, spaceWidth)
	}

	for p := range *robots {
		m[p.y][p.x] = true
	}

	printMap := func(m [][]bool) string {
		result := ""
		for y := range m {
			for x := range m[y] {
				if m[y][x] {
					result += "#"
				} else {
					result += "."
				}
			}
			result += "\n"
		}
		return result
	}

	err := os.WriteFile("day14/secondSolutionMap.txt", []byte(printMap(m)), 0644)
	if err != nil {
		return
	}

}
