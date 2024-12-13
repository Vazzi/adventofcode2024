package day11

import (
	"math"
)

func solution(data []int, count int) int {
	state := map[int]int{}
	for i := range data {
		state[data[i]] += 1
	}

	for c := count; c > 0; c-- {
		nextState := map[int]int{}
		for num := range state {
			if num == 0 {
				nextState[1] += state[num]
				continue
			}

			digitCount := int(math.Floor(math.Log10(float64(num)))) + 1

			if digitCount%2 == 0 {
				divisor := int(math.Pow10(digitCount / 2))
				left := int(math.Floor(float64(num / divisor)))
				right := num % divisor
				nextState[left] += state[num]
				nextState[right] += state[num]
			} else {
				nextState[num*2024] += state[num]
			}
		}
		state = nextState
	}

	result := 0
	for _, count := range state {
		result += count
	}

	return result
}
