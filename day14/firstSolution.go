package day14

import (
	"adventofcode2024/utils"
)

func firstSolution(fileName string, spaceWidth, spaceHeight int) int {
	seconds := 100
	robots := readRobotInputs(fileName)

	for i := 0; i < seconds; i++ {
		for ri, r := range robots {
			r.move(spaceWidth, spaceHeight)
			robots[ri] = r
		}
	}

	return countRobotsInQuadrants(&robots, spaceWidth, spaceHeight)
}

func readRobotInputs(fileName string) []robot {
	lines := utils.ReadLines(fileName)
	var robots []robot

	for _, line := range lines {
		robots = append(robots, *newRobot(line))
	}

	return robots
}

func countRobotsInQuadrants(robots *[]robot, spaceWidth, spaceHeight int) int {
	quadrantsCount := make(map[int]int, 4)
	for _, r := range *robots {
		quadrantId, ok := r.inQuadrant(spaceWidth, spaceHeight)
		if ok {
			quadrantsCount[quadrantId]++
		}
	}

	result := 1
	for i := range quadrantsCount {
		if quadrantsCount[i] == 0 {
			continue
		}
		result *= quadrantsCount[i]
	}

	return result
}
