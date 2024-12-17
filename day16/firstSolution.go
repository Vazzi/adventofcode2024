package day16

import (
	"container/heap"
	"fmt"
	"math"
)

func firstSolution(grid [][]string) int {
	start, end := findStartEnd(&grid)
	score := findPath(&grid, start, end)
	return score
}

func findPath(grid *[][]string, start, end point) int {
	// up, right, down, left
	directions := []point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	nodes := &priorityQueue{}
	heap.Init(nodes)

	visited := make(map[point]bool)

	startNode := &node{point: start, score: 0, heuristicScore: manhattanDistance(start, end)}
	startNode.totalScore = startNode.score + startNode.heuristicScore

	heap.Push(nodes, startNode)

	for nodes.Len() > 0 {
		currentNode := heap.Pop(nodes).(*node)

		if currentNode.point == end {
			return countScoreOfPath(currentNode, grid)
		}

		visited[currentNode.point] = true

		for _, dir := range directions {
			neighbour := point{currentNode.x + dir.x, currentNode.y + dir.y}

			if (*grid)[neighbour.y][neighbour.x] == tileWall || visited[neighbour] {
				continue
			}

			score := currentNode.score + 1
			heuristicScore := currentNode.heuristicScore + 1
			totalScore := score + heuristicScore

			neighbourNode := &node{point: neighbour, score: score, heuristicScore: heuristicScore, totalScore: totalScore, parent: currentNode}

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

func countScoreOfPath(n *node, grid *[][]string) int {
	path := reconstructPath(n)

	score := 0
	dir := right
	curr := path[0]
	for i := 1; i < len(path); i++ {
		next := path[i]
		if next.x > curr.x {
			if dir != right {
				score += costTurn
				dir = right
			}
		} else if next.x < curr.x {
			if dir != left {
				score += costTurn
				dir = left
			}
		} else if next.y > curr.y {
			if dir != down {
				score += costTurn
				dir = down
			}
		} else if next.y < curr.y {
			if dir != up {
				score += costTurn
				dir = up
			}
		}

		score++
		curr = next
		fmt.Println(curr, score, (*grid)[curr.y][curr.x])
	}
	return score
}

func reconstructPath(n *node) []*point {
	var path []*point
	for current := n; current != nil; current = current.parent {
		path = append([]*point{&current.point}, path...)
	}
	return path
}
