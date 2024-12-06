package utils

import (
	"bufio"
	"os"
)

func ReadChars(fileName string) [][]string {
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
