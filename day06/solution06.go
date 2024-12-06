package day06

import (
	"adventofcode2024/utils"
	"fmt"
)

func Main() {
	mapData := utils.ReadChars("./day06/input.txt")
	fmt.Println("Result for the first solution is: ", firstSolution(mapData))
}

const (
	Up    int = 0
	Right     = 1
	Down      = 2
	Left      = 3
)

func firstSolution(mapData [][]string) int {
	guardPosX, guardPosY := findGuard(mapData)
	guardDir := Up
	steps := 0

	for {
		nextX, nextY := moveGuard(guardDir, guardPosX, guardPosY)

		if nextX < 0 || nextX >= len(mapData[0]) || nextY < 0 || nextY >= len(mapData) {
			return steps
		}

		if mapData[nextY][nextX] == "#" {
			guardDir = turnGuardRight(guardDir)
			continue
		}
		guardPosX, guardPosY = nextX, nextY
		if mapData[guardPosY][guardPosX] != "X" {
			mapData[guardPosY][guardPosX] = "X"
			steps++
		}
	}

	return steps
}

func moveGuard(dir, x, y int) (int, int) {
	switch dir {
	case Up:
		return x, y - 1
	case Right:
		return x + 1, y
	case Down:
		return x, y + 1
	case Left:
		return x - 1, y
	}
	return x, y
}

func turnGuardRight(guardDir int) int {
	return (guardDir + 1) % 4
}

func findGuard(mapData [][]string) (int, int) {
	for y := range mapData {
		for x := range mapData[y] {
			if mapData[y][x] == "^" {
				return x, y
			}
		}
	}
	panic("Guard not found")
}
