package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Main() {
	rules := readRules("./day05/rules02.txt")
	data := readData("./day05/data02.txt")
	fmt.Println("Result for the first solution is: ", firstSolution(rules, data))
	//fmt.Println("Result for the second solution is: ", secondSolution(rules, data))
}

func firstSolution(rules map[string][]string, data [][]string) int {
	score := 0

	for _, order := range data {
		if checkOrder(order, rules) {
			score += getMiddleNumber(order)
		}
	}

	return score
}

func checkOrder(data []string, rules map[string][]string) bool {
	for numberIndex := len(data) - 1; numberIndex >= 0; numberIndex-- {
		number := data[numberIndex]
		for i := numberIndex; i >= 0; i-- {
			if slices.Contains(rules[number], data[i]) {
				return false
			}
		}
	}

	return true
}

func getMiddleNumber(data []string) int {
	result, _ := strconv.Atoi(data[len(data)/2])
	return result
}

func secondSolution(rules [][]int, data []int) int {
	return 0
}

func readRules(fileName string) map[string][]string {
	file, _ := os.Open(fileName)
	defer file.Close()

	data := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "|")
		if data[values[0]] == nil {
			data[values[0]] = make([]string, 0)
		}
		data[values[0]] = append(data[values[0]], values[1])
	}
	return data
}

func readData(fileName string) [][]string {
	file, _ := os.Open(fileName)
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		data = append(data, line)
	}
	return data

}
