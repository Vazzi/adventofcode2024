package day11

import "math"

func firstSolution(data []int) int {

	for c := 25; c > 0; c-- {
		for i := 0; i < len(data); i++ {
			num := data[i]
			if num == 0 {
				data[i] = 1
				continue
			}
			digitCount := int(math.Floor(math.Log10(float64(num)))) + 1

			if digitCount%2 == 0 {
				divisor := int(math.Pow10(digitCount / 2))
				left := int(math.Floor(float64(num / divisor)))
				right := num % divisor

				newData := make([]int, len(data)+1)
				for index := 0; index < cap(newData); index++ {
					if index == i {
						newData[i] = left
						newData[i+1] = right
						index++ //jump to next
					} else if index > i {
						newData[index] = data[index-1]
					} else {
						newData[index] = data[index]
					}
				}
				data = newData
				i++ //jump to next
			} else {
				data[i] = data[i] * 2024
			}
		}
	}

	return len(data)
}
