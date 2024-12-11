package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputDataPath(day int, test bool) string {
	fileName := fmt.Sprintf("day%02d.txt", day)

	if test {
		return "../inputData/test/" + fileName
	}

	return "./inputData/" + fileName
}

func ReadInts(fileName string, sep string) [][]int {
	lines := ReadLines(fileName)
	data := make([][]int, 0)

	for i := range lines {
		values := strings.Split(lines[i], sep)
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

func ReadLines(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}

func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadWords(fileName string, delimiter string) [][]string {
	lines := ReadLines(fileName)
	var data [][]string

	for _, line := range lines {
		line := strings.Split(line, delimiter)
		data = append(data, line)
	}
	return data
}

func ReadChars(fileName string) [][]string {
	lines := ReadLines(fileName)

	data := make([][]string, 0)

	for _, line := range lines {
		runes := []rune(line)
		chars := make([]string, len(runes))
		for i := range runes {
			chars[i] = string(runes[i])
		}
		data = append(data, chars)
	}

	return data
}
