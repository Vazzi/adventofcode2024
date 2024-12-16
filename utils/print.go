package utils

import "fmt"

func PrintDayHeadline(headline string) {
	fmt.Println("***** ", headline, " *****")
	fmt.Println()
}

func PrintMap(data *[][]string) {
	for y := range *data {
		for x := range (*data)[y] {
			fmt.Print((*data)[y][x])
		}
		fmt.Println()
	}
}
