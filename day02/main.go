package day02

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Main() {
	//TODO: This solution is incorrect it works only if you do not tolerateSingleError - need to solve it

	data := readDataFromFile("./day02/input02.txt")
	numberOfSafeReports := computeNumberOfSafeReports(data, true)

	fmt.Println("Number of safe reports is ", numberOfSafeReports)
}

func computeNumberOfSafeReports(data [][]int, tolerateSingleError bool) int {
	score := 0
	maxHearts := 1
	if tolerateSingleError {
		maxHearts = 2
	}

	for _, report := range data {
		prevNum := report[0]
		direction := 0
		hearts := maxHearts

		for i := 1; i < len(report); i++ {
			number := report[i]
			diff := prevNum - number

			if diff == 0 || math.Abs(float64(diff)) > 3 {
				hearts--
				if hearts == 0 {
					break
				} else {
					continue
				}
			}

			if (diff > 0 && direction < 0) || (diff < 0 && direction > 0) {
				hearts--
				if hearts == 0 {
					break
				} else {
					continue
				}
			}

			prevNum = number
			direction = diff
		}
		if hearts > 0 {
			score++
		}
	}

	return score
}

func readDataFromFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		lineData := make([]int, len(values))
		for index, str := range values {
			number, convertErr := strconv.Atoi(str)
			if convertErr != nil {
				log.Fatal(convertErr)
			}
			lineData[index] = number
		}
		data = append(data, lineData)
	}

	return data
}
