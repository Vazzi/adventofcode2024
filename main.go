package main

import (
	"adventofcode2024/day01"
	"adventofcode2024/day02"
	"adventofcode2024/day03"
	"adventofcode2024/day04"
	"adventofcode2024/day05"
	"adventofcode2024/day06"
	"adventofcode2024/day07"
	"fmt"
)

func main() {
	currentDay := 6
	printHeader(currentDay)
	switch currentDay {
	case 1:
		day01.Main()
	case 2:
		day02.Main()
	case 3:
		day03.Main()
	case 4:
		day04.Main()
	case 5:
		day05.Main()
	case 6:
		day06.Main()
	case 7:
		day07.Main()
	}
}

func printHeader(day int) {
	fmt.Println("****************************************")
	fmt.Println("**********Advent of Code 2024.**********")
	fmt.Println("****************************************")
	fmt.Println("Current day: ", day)
	fmt.Println()
}
