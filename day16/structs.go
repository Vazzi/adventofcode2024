package day16

type point struct {
	x, y int
}

const (
	tileWall  string = "#"
	tileEmpty        = "."
	tileStart        = "S"
	tileEnd          = "E"
)

const (
	costStep int = 1
	costTurn     = 1000
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(a, b int) bool {
	return pq[a].totalScore < pq[b].totalScore
}

func (pq priorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
	pq[a].index, pq[b].index = a, b
}

func (pq *priorityQueue) Push(x interface{}) {
	newNode := x.(*node)
	newNode.index = len(*pq)
	*pq = append(*pq, newNode)
}

func (pq *priorityQueue) Pop() interface{} {
	oldQueue := *pq
	queueLen := len(oldQueue)

	// Get last node
	lastNode := oldQueue[queueLen-1]
	lastNode.index = -1 // For safety

	// Remove last node
	*pq = oldQueue[0 : queueLen-1]

	// Return last node
	return lastNode
}

type node struct {
	point
	score          int
	heuristicScore int
	totalScore     int
	parent         *node
	index          int
	dir            direction
}
