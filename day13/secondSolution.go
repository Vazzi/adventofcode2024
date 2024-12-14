package day13

import (
	"math"
)

func secondSolution(fileName string) int {
	tasks := readData(fileName)

	for i := range tasks {
		tasks[i].price.x += 10000000000000
		tasks[i].price.y += 10000000000000
	}

	var result int = 0

	for t := range tasks {
		result += int(solveTaskMathematically(tasks[t]))
	}
	return result
}

func solveTaskMathematically(t task) int {

	x := float64(t.price.x)
	y := float64(t.price.y)
	ax := float64(t.a.x)
	ay := float64(t.a.y)
	bx := float64(t.b.x)
	by := float64(t.b.y)

	cb := (y - (ay * x / ax)) / (by - (ay * bx / ax))
	ca := (1.0 / ax) * (x - bx*(y-((ay/ax)*x))/(by-(ay/ax*bx)))

	if isWholeNumber(ca) && isWholeNumber(cb) {
		return int(ca)*aCost + int(cb)*bCost
	}
	return 0
}

func isWholeNumber(num float64) bool {
	tolerance := 1e-9 // tolerance for floating point error
	return math.Abs(num-math.Floor(num)) < tolerance
}
