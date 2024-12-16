package day15

import "adventofcode2024/utils"

func secondSolution(mapPath, movesPath string) int {
	inputData := utils.ReadChars(mapPath)
	movesData := readMovesFromFile(movesPath)
	// Double the map
	doubledData := doubleWidthOfMap(&inputData)
	playground := newMapData(doubledData)
	utils.PrintMap(&playground.data)

	for _, move := range movesData {
		playground.moveRobot2(move)
	}
	utils.PrintMap(&playground.data)

	return playground.countBoxesCoord2()
}

const (
	charBoxLeft  string = "["
	charBoxRight        = "]"
)

func (p *mapData) canMoveObject(dir string, startPos position) bool {
	// Moving left or right
	if dir == ">" || dir == "<" {
		nextMove := startPos.move(dir)
		nextMove = nextMove.move(dir) // jump twice
		nextObject := p.data[nextMove.y][nextMove.x]
		if nextObject == charWall {
			return false
		} else if nextObject == charBoxLeft || nextObject == charBoxRight {
			return p.canMoveObject(dir, nextMove)
		} else if nextObject == charEmpty {
			return true
		}
	}

	// Moving up or down
	leftSide, rightSide := p.computeBoxSides(startPos)
	leftMove := leftSide.move(dir)
	rightMove := rightSide.move(dir)

	nextLeftObject := p.data[leftMove.y][leftMove.x]
	nextRightObject := p.data[rightMove.y][rightMove.x]
	// Wall
	if nextLeftObject == charWall || nextRightObject == charWall {
		return false
	}
	// One box
	if nextLeftObject == charBoxLeft && nextRightObject == charBoxRight {
		return p.canMoveObject(dir, leftMove)
	}

	// Box on left side
	if nextLeftObject == charBoxRight {
		if !(p.canMoveObject(dir, leftMove)) {
			return false
		}
	}
	// Box on right side
	if nextRightObject == charBoxLeft {
		if !(p.canMoveObject(dir, rightMove)) {
			return false
		}
	}
	return true
}

func (p *mapData) computeBoxSides(startPos position) (position, position) {
	selfObject := p.data[startPos.y][startPos.x]
	leftSide := position{startPos.x, startPos.y}
	rightSide := position{startPos.x, startPos.y}
	// Construct whole object
	if selfObject == charBoxLeft {
		rightSide.x = startPos.x + 1
	} else {
		leftSide.x = startPos.x - 1
	}
	return leftSide, rightSide
}

func (p *mapData) moveObject(dir string, startPos position) {
	// Moving left or right
	if dir == ">" || dir == "<" {
		stepOne := startPos.move(dir)
		stepTwo := stepOne.move(dir)
		nextObject := p.data[stepTwo.y][stepTwo.x]
		if nextObject == charBoxLeft || nextObject == charBoxRight {
			p.moveObject(dir, stepTwo)
		}

		p.data[stepTwo.y][stepTwo.x] = p.data[stepOne.y][stepOne.x]
		p.data[stepOne.y][stepOne.x] = p.data[startPos.y][startPos.x]
		p.data[startPos.y][startPos.x] = charEmpty
		return
	}

	// Moving up or down
	leftSide, rightSide := p.computeBoxSides(startPos)
	leftMove := leftSide.move(dir)
	rightMove := rightSide.move(dir)

	nextLeftObject := p.data[leftMove.y][leftMove.x]
	nextRightObject := p.data[rightMove.y][rightMove.x]
	// One box
	if nextLeftObject == charBoxLeft && nextRightObject == charBoxRight {
		p.moveObject(dir, leftMove)
	}
	// Left box
	if nextLeftObject == charBoxRight {
		p.moveObject(dir, leftMove) // box on left side
	}
	// Right box
	if nextRightObject == charBoxLeft {
		p.moveObject(dir, rightMove) // box on right side
	}
	p.data[leftMove.y][leftMove.x] = p.data[leftSide.y][leftSide.x]
	p.data[rightMove.y][rightMove.x] = p.data[rightSide.y][rightSide.x]
	p.data[leftSide.y][leftSide.x] = charEmpty
	p.data[rightSide.y][rightSide.x] = charEmpty
}

func (p *mapData) countBoxesCoord2() int {
	score := 0
	for y := range p.data {
		for x := range p.data[y] {
			if p.data[y][x] == charBoxLeft {
				score += y*100 + x
			}
		}
	}
	return score
}

func (p *mapData) moveRobot2(move string) bool {
	nextMove := p.robot.move(move)
	nextObject := p.data[nextMove.y][nextMove.x]

	// Wall do not move
	if nextObject == charWall {
		return false
	}
	// Object try to move the object
	if nextObject == charBoxLeft || nextObject == charBoxRight {
		if p.canMoveObject(move, nextMove) {
			p.moveObject(move, nextMove)
		} else {
			return false
		}
	}

	p.data[nextMove.y][nextMove.x] = charRobot
	p.data[p.robot.y][p.robot.x] = charEmpty
	p.robot.x = nextMove.x
	p.robot.y = nextMove.y
	return true
}

func doubleWidthOfMap(data *[][]string) [][]string {
	newData := make([][]string, len(*data))
	for y := range *data {
		newData[y] = make([]string, len((*data)[y])*2)
		for x := range (*data)[y] {
			if (*data)[y][x] == charBox {
				newData[y][x*2] = charBoxLeft
				newData[y][x*2+1] = charBoxRight
				continue
			}
			if (*data)[y][x] == charRobot {
				newData[y][x*2] = charRobot
				newData[y][x*2+1] = charEmpty
				continue
			}
			newData[y][x*2] = (*data)[y][x]
			newData[y][x*2+1] = (*data)[y][x]
		}
	}
	return newData
}
