package day01

import (
	"fmt"
	"math"
	"sort"
)

func Main() {
	distance := computeDistance(firstList, secondList)
	fmt.Println("Distance between the two lists is: ", distance)

	score := similarityScore(firstList, secondList)
	fmt.Println("Similarity score between the two lists is: ", score)
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

func similarityScore(firstList, secondList []int) int {
	sort.Ints(secondList)

	var valuesDict = make(map[int]int)
	lastValue := 0
	for _, value := range secondList {
		if lastValue == value {
			valuesDict[value] += 1
		} else {
			valuesDict[value] = 1
			lastValue = value
		}
	}

	score := 0
	for _, firstValue := range firstList {
		score += valuesDict[firstValue] * firstValue
	}

	return score
}
