package day12

func firstSolution(garden [][]string) int {
	result := 0
	plantsChecked := make([]bool, len(garden)*len(garden[0]))
	gWidth := len(garden[0])

	for y := range garden {
		for x := range garden[y] {
			if plantsChecked[y*gWidth+x] {
				continue
			}
			area, fence := computeAreaScore(garden, plantsChecked, x, y, garden[y][x])
			result += area * fence
		}
	}

	return result
}

func computeAreaScore(garden [][]string, plantsChecked []bool, x, y int, plant string) (int, int) {
	if x < 0 || y < 0 || x >= len(garden[0]) || y >= len(garden) {
		return 0, 1
	}
	if garden[y][x] != plant {
		return 0, 1
	}
	if plantsChecked[y*len(garden[0])+x] {
		return 0, 0
	}
	plantsChecked[y*len(garden[0])+x] = true

	area, fence := 0, 0
	areaTmp, fenceTmp := 0, 0

	// Check right
	areaTmp, fenceTmp = computeAreaScore(garden, plantsChecked, x+1, y, plant)
	area += areaTmp
	fence += fenceTmp

	// Check left
	areaTmp, fenceTmp = computeAreaScore(garden, plantsChecked, x-1, y, plant)
	area += areaTmp
	fence += fenceTmp

	// Check up
	areaTmp, fenceTmp = computeAreaScore(garden, plantsChecked, x, y+1, plant)
	area += areaTmp
	fence += fenceTmp

	// Check down
	areaTmp, fenceTmp = computeAreaScore(garden, plantsChecked, x, y-1, plant)
	area += areaTmp
	fence += fenceTmp

	// Return result and do not forget to return self
	return area + 1, fence
}
