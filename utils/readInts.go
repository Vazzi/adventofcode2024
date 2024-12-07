package utils

import (
	"strconv"
	"strings"
)

func ReadInts(fileName string) [][]int {
	lines := ReadLines(fileName)
	data := make([][]int, 0)

	for i := range lines {
		values := strings.Split(lines[i], " ")
		lineData := make([]int, len(values))
		for index, str := range values {
			number, convertErr := strconv.Atoi(str)
			if convertErr != nil {
				panic(convertErr)
			}
			lineData[index] = number
		}
		data = append(data, lineData)
	}

	return data
}
