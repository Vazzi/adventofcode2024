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
	// up, right, down, left
	directions := []point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	nodes := &priorityQueue{}
	heap.Init(nodes)

	visited := make(map[point]bool)

	startNode := &node{point: start, score: 0, heuristicScore: manhattanDistance(start, end), dir: right}
	startNode.totalScore = startNode.score + startNode.heuristicScore

	heap.Push(nodes, startNode)

	for nodes.Len() > 0 {
		currentNode := heap.Pop(nodes).(*node)

		if currentNode.point == end {
			return currentNode.score
		}

		visited[currentNode.point] = true

		for i, dir := range directions {
			neighbour := point{currentNode.x + dir.x, currentNode.y + dir.y}

			if (*grid)[neighbour.y][neighbour.x] == tileWall || visited[neighbour] {
				continue
			}

			score := currentNode.score + costStep
			if currentNode.dir != direction(i) {
				score += costTurn
			}

			heuristicScore := currentNode.heuristicScore + 1
			totalScore := score + heuristicScore

			neighbourNode := &node{point: neighbour, score: score, heuristicScore: heuristicScore, totalScore: totalScore, parent: currentNode, dir: direction(i)}

			heap.Push(nodes, neighbourNode)
		}
	}

	return 0
}

func manhattanDistance(a, b point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func findStartEnd(mapData *[][]string) (start, end point) {
	for y := range *mapData {
		for x := range (*mapData)[y] {
			if (*mapData)[y][x] == tileStart {
				start = point{x, y}
			}
			if (*mapData)[y][x] == tileEnd {
				end = point{x, y}
			}
		}
	}
	return start, end
}
