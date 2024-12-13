package day13

import (
	"adventofcode2024/utils"
	"regexp"
	"strconv"
)

const (
	aCost int = 3
	bCost     = 1
)

func firstSolution(fileName string) int {
	tasks := readData(fileName)
	result := 0

	for t := range tasks {
		result += solveTask(tasks[t])
	}
	return result
}

type coord struct {
	x, y int
}

type task struct {
	a, b  coord
	price coord
}

func readData(fileName string) []task {
	lines := utils.ReadLines(fileName)
	tasks := []task{}

	for i := 0; i < len(lines); i += 4 {
		a := getNumbers(lines[i])
		b := getNumbers(lines[i+1])
		price := getNumbers(lines[i+2])
		tasks = append(tasks, task{a, b, price})
	}
	return tasks
}

func getNumbers(line string) coord {
	re := regexp.MustCompile(`[XY][+=]([0-9]+)`)
	matches := re.FindAllStringSubmatch(line, -1)

	x, _ := strconv.Atoi(matches[0][1])
	y, _ := strconv.Atoi(matches[1][1])

	return coord{x, y}
}

func solveTask(t task) int {
	countA := t.price.x / t.a.x
	countB := 0

	for {
		x := t.a.x*countA + t.b.x*countB
		y := t.a.y*countA + t.b.y*countB

		if x == t.price.x && y == t.price.y {
			if countA < 100 && countB < 100 {
				break
			}
		}
		if x > t.price.x || y > t.price.y {
			countA--
			continue
		}

		if countA < 0 {
			return 0
		}

		countB++
	}

	return countA*aCost + countB*bCost

}
