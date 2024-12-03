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
	data := readDataFromFile("./day02/input02.txt")
	numberOfSafeReports := computeNumberOfSafeReports(data, false)

	fmt.Println("Number of safe reports is  ", numberOfSafeReports)

	numberOfSafeReportsWithToleration := computeNumberOfSafeReports(data, true)

	fmt.Println("Number of safe reports is with toleration of one error is ", numberOfSafeReportsWithToleration)
}

func checkReport(report []int) bool {
	direction := report[0] - report[1]

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]

		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if (diff > 0 && direction < 0) || (diff < 0 && direction > 0) {
			return false
		}
	}
	return true
}

func computeNumberOfSafeReports(data [][]int, tolerateSingleError bool) int {
	score := 0
	for _, report := range data {
		if checkReport(report) {
			score++
		} else if tolerateSingleError {
			for i := 0; i < len(report); i++ {
				modifiedArray := make([]int, 0, len(report)-1)
				modifiedArray = append(modifiedArray, report[:i]...)
				modifiedArray = append(modifiedArray, report[i+1:]...)
				if checkReport(modifiedArray) {
					score++
					break
				}
			}
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
