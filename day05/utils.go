package day05

import (
	"bufio"
	"os"
	"strings"
)

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
