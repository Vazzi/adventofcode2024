package day16

import (
	"container/heap"
	"math"
)

func firstSolution(grid [][]string) int {
	start, end := findStartEnd(&grid)
	score := findPath(&grid, start, end)
	return score
}

func findPath(grid *[][]string, start, end point) int {
	nodes := &priorityQueue{}
	heap.Init(nodes)

	visited := make(map[point]int)

	startNode := &node{point: start, score: 0}

	heap.Push(nodes, startNode)

	for nodes.Len() > 0 {
		currentNode := heap.Pop(nodes).(*node)
		if currentNode.point == end {
			return currentNode.score
		} else if visitedScore(&visited, currentNode.point) > currentNode.score {
			visited[currentNode.point] = currentNode.score

			nextStep := step(currentNode)
			if (*grid)[nextStep.y][nextStep.x] != tileWall { // Step ahead
				neighbourNode := &node{point: nextStep, score: currentNode.score + costStep, parent: currentNode}
				heap.Push(nodes, neighbourNode)
			}

			leftPoint := point{currentNode.x, currentNode.y, turnLeft(currentNode.dir)}
			turnScore := currentNode.score + costTurn
			// Turn left
			heap.Push(nodes, &node{point: leftPoint, score: turnScore})
			rightPoint := point{currentNode.x, currentNode.y, turnRight(currentNode.dir)}
			// Turn right
			heap.Push(nodes, &node{point: rightPoint, score: turnScore})
		}
	}

	return 0
}

func visitedScore(data *map[point]int, key point) int {
	score, ok := (*data)[key]
	if ok {
		return score
	}
	return math.MaxInt
}

func turnLeft(dir direction) direction {
	if dir == up {
		return left
	}
	return direction(int(dir) - 1)
}

func turnRight(dir direction) direction {
	if dir == left {
		return up
	}
	return direction(int(dir) + 1)
}

func step(curr *node) point {
	switch curr.dir {
	case up:
		return point{curr.x, curr.y - 1, curr.dir}
	case down:
		return point{curr.x, curr.y + 1, curr.dir}
	case left:
		return point{curr.x - 1, curr.y, curr.dir}
	case right:
		return point{curr.x + 1, curr.y, curr.dir}
	}
	return curr.point
}

func findStartEnd(mapData *[][]string) (start, end point) {
	for y := range *mapData {
		for x := range (*mapData)[y] {
			if (*mapData)[y][x] == tileStart {
				start = point{x, y, right}
			}
			if (*mapData)[y][x] == tileEnd {
				end = point{x, y, right}
			}
		}
	}
	return start, end
}
