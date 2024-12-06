package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadWords(fileName string, delimiter string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), delimiter)
		data = append(data, line)
	}
	return data
}
