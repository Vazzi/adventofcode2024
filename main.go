package main

import (
	"adventofcode2024/day05"
	"fmt"
)

func main() {
	day05.Main()
	data := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	numberIndex := 7
	i := 4

	data = append(data[:i], append(data[i+1:numberIndex+1], append([]string{data[i]}, data[numberIndex+1:]...)...)...)
	fmt.Println(data)

}
