package day01

import (
	"fmt"
	"math"
	"sort"
)

func Main() {
	distance := computeDistance(firstList, secondList)
	fmt.Println("Distance between the two lists is: ", distance)
}

func computeDistance(firstList, secondList []int) int {
	sort.Ints(firstList)
	sort.Ints(secondList)

	totalDistance := 0

	for i, firstValue := range firstList {
		totalDistance += int(math.Abs(float64(firstValue - secondList[i])))
	}

	return totalDistance
}
