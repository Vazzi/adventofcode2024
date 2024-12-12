package day11

import (
	"math"
)

func solution(data []int, count int) int {

	for c := count; c > 0; c-- {
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
				data[i] = left
				data = append(data[:i+1], append([]int{right}, data[i+1:]...)...)
				i++ //jump to next
			} else {
				data[i] = data[i] * 2024
			}
		}
	}

	return len(data)
}
