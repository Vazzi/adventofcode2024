package utils

import (
	"os"
	"strings"
)

func ReadLines(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}
