package main

import (
	"adventofcode2024/day01"
	"adventofcode2024/day02"
	"adventofcode2024/day03"
	"adventofcode2024/day04"
	"adventofcode2024/day05"
	"adventofcode2024/day06"
	"adventofcode2024/day07"
	"adventofcode2024/day08"
	"adventofcode2024/day09"
	"adventofcode2024/day10"
	"adventofcode2024/day11"
	"adventofcode2024/day12"
	"adventofcode2024/day13"
	"adventofcode2024/day14"
	"adventofcode2024/day15"
	"adventofcode2024/day16"
	"adventofcode2024/day17"
	"fmt"
)

func main() {
	var currentDay = 17

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
	case 8:
		day08.Main()
	case 9:
		day09.Main()
	case 10:
		day10.Main()
	case 11:
		day11.Main()
	case 12:
		day12.Main()
	case 13:
		day13.Main()
	case 14:
		day14.Main()
	case 15:
		day15.Main()
	case 16:
		day16.Main()
	case 17:
		day17.Main()
	}
}

func printHeader(day int) {
	fmt.Println("****************************************")
	fmt.Println("**********Advent of Code 2024.**********")
	fmt.Println("****************************************")
	fmt.Println("Current day: ", day)
	fmt.Println()
}
