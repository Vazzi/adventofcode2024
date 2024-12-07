package utils

import (
	"strings"
)

func ReadWords(fileName string, delimiter string) [][]string {
	lines := ReadLines(fileName)
	var data [][]string

	for _, line := range lines {
		line := strings.Split(line, delimiter)
		data = append(data, line)
	}
	return data
}
