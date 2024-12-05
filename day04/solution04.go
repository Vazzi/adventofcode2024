package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Main() {
	data := readDataFromFile("./day04/input02.txt")
	fmt.Println("Result for the first solution is: ", firstSolution(data))
	//fmt.Println("Result for the second solution is: ", secondSolution(input01))
}

func firstSolution(data [][]string) int {
	height := len(data)
	width := len(data[0])
	total := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] == "X" {
				total += findWord(data, x, y)
			}
		}
	}

	return total
}

func findWord(data [][]string, x, y int) int {
	total := 0

	// Check right
	if x+3 < len(data[0]) {
		total += isXMAS(data[y][x : x+4])
	}

	// Check left
	if x-3 >= 0 {
		total += isXMAS(data[y][x-3 : x+1])
	}

	// Check down
	if y+3 < len(data) {
		total += isXMAS([]string{data[y][x], data[y+1][x], data[y+2][x], data[y+3][x]})
	}

	// Check up
	if y-3 >= 0 {
		total += isXMAS([]string{data[y-3][x], data[y-2][x], data[y-1][x], data[y][x]})
	}

	// Check diagonal down right
	if x+3 < len(data[0]) && y+3 < len(data) {
		total += isXMAS([]string{data[y][x], data[y+1][x+1], data[y+2][x+2], data[y+3][x+3]})
	}

	// Check diagonal down left
	if x-3 >= 0 && y+3 < len(data) {
		total += isXMAS([]string{data[y][x], data[y+1][x-1], data[y+2][x-2], data[y+3][x-3]})
	}

	// Check diagonal up right
	if x+3 < len(data[0]) && y-3 >= 0 {
		total += isXMAS([]string{data[y][x], data[y-1][x+1], data[y-2][x+2], data[y-3][x+3]})
	}

	// Check diagonal up left
	if x-3 >= 0 && y-3 >= 0 {
		total += isXMAS([]string{data[y][x], data[y-1][x-1], data[y-2][x-2], data[y-3][x-3]})
	}

	return total
}

func isXMAS(data []string) int {
	input := strings.Join(data, "")

	if input == "XMAS" || input == "SAMX" {
		return 1
	}

	return 0
}

func readDataFromFile(fileName string) [][]string {
	file, _ := os.Open(fileName)

	data := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		chars := make([]string, len(runes))
		for i, _ := range runes {
			chars[i] = string(runes[i])
		}
		data = append(data, chars)
	}

	return data
}
