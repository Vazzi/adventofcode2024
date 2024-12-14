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

	var result int
	for t := range tasks {
		result += solveTaskMathematically(tasks[t])
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

	// Substitution does work on solution 1 but not on solution 2 (10000000000000 +)
	//cb := (y - (ay * x / ax)) / (by - (ay * bx / ax))
	//ca := (1 / ax) * (x - bx*(y-((ay/ax)*x))/(by-(ay/ax*bx)))

	// Cramer's rule does work on both
	// God bless the internet :)
	ca := (x*by - y*bx) / (ax*by - ay*bx)
	cb := (ax*y - ay*x) / (ax*by - ay*bx)

	countA, aok := getWholeNumber(ca)
	countB, bok := getWholeNumber(cb)

	if aok && ca >= 0 && bok && cb >= 0 {
		return countA*aCost + countB*bCost
	}

	return 0
}

func getWholeNumber(num float64) (int, bool) {
	tolerance := 1e-9 // tolerance for floating point error
	if math.Abs(num-math.Floor(num)) < tolerance {
		return int(math.Floor(num)), true
	}
	return 0, false
}
