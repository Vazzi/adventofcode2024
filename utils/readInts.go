package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInts(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
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
				panic(convertErr)
			}
			lineData[index] = number
		}
		data = append(data, lineData)
	}

	return data
}
