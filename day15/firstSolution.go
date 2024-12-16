package day15

import (
	"adventofcode2024/utils"
)

const (
	charRobot string = "@"
	charWall         = "#"
	charBox          = "O"
	charEmpty        = "."
)

type position struct {
	x, y int
}

type mapData struct {
	data  [][]string
	robot position
}

func newMapData(data [][]string) *mapData {
	robotPos, robotFound := getRobotPosition(&data)
	if !robotFound {
		panic("Robot not found")
	}
	return &mapData{
		data:  data,
		robot: robotPos,
	}
}

func firstSolution(mapPath, movesPath string) int {
	inputData := utils.ReadChars(mapPath)
	movesData := readMovesFromFile(movesPath)

	playground := newMapData(inputData)

	for _, move := range movesData {
		playground.moveRobot(move)
	}

	return playground.countBoxesCoordinates()
}

func (p *position) move(dir string) position {
	switch dir {
	case "^":
		return position{p.x, p.y - 1}
	case "v":
		return position{p.x, p.y + 1}
	case "<":
		return position{p.x - 1, p.y}
	case ">":
		return position{p.x + 1, p.y}
	}
	return position{p.x, p.y}
}

func (p *mapData) tryToMoveObjects(dir string, startPos position) bool {
	nextMove := startPos.move(dir)
	if p.data[nextMove.y][nextMove.x] == charWall {
		return false
	}
	if p.data[nextMove.y][nextMove.x] == charBox {
		if !p.tryToMoveObjects(dir, nextMove) {
			return false
		}
	}
	p.data[nextMove.y][nextMove.x] = charBox
	p.data[startPos.y][startPos.x] = charEmpty
	return true
}

func (p *mapData) countBoxesCoordinates() int {
	score := 0
	for y := range p.data {
		for x := range p.data[y] {
			if p.data[y][x] == charBox {
				score += y*100 + x
			}
		}
	}
	return score
}

func (p *mapData) moveRobot(move string) bool {
	nextMove := p.robot.move(move)

	// Wall do not move
	if p.data[nextMove.y][nextMove.x] == charWall {
		return false
	}
	// Object try to move the object
	if p.data[nextMove.y][nextMove.x] == charBox {
		if !p.tryToMoveObjects(move, nextMove) {
			return false
		}
	}

	p.data[nextMove.y][nextMove.x] = charRobot
	p.data[p.robot.y][p.robot.x] = charEmpty
	p.robot.x = nextMove.x
	p.robot.y = nextMove.y
	return true
}

func getRobotPosition(data *[][]string) (position, bool) {
	for y := range *data {
		for x := range (*data)[y] {
			if (*data)[y][x] == charRobot {
				return position{x, y}, true
			}
		}
	}
	return position{-1, -1}, false
}

func readMovesFromFile(fileName string) []string {
	lines := utils.ReadLines(fileName)
	var data []string

	for line := range lines {
		for _, char := range lines[line] {
			data = append(data, string(char))
		}
	}
	return data
}
