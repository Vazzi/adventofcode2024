package day12

type garden struct {
	plants  [][]string
	checked []bool
	width   int
	height  int
}

func newGarden(data [][]string) *garden {
	width := len(data[0])
	height := len(data)
	checked := make([]bool, width*height)

	return &garden{
		data,
		checked,
		width,
		height,
	}
}

func (data garden) isChecked(x, y int) bool {
	return data.checked[y*data.width+x]
}

func (data *garden) setChecked(x, y int) {
	data.checked[y*data.width+x] = true
}

func (data garden) isSamePlant(x, y int, plant string) bool {
	if x < 0 || y < 0 || x >= data.width || y >= data.height {
		return false
	}
	if data.plants[y][x] == plant {
		return true
	}
	return false
}

func secondSolution(data [][]string) int {
	result := 0
	inputG := newGarden(data)

	for y := range inputG.height {
		for x := range inputG.width {
			if inputG.isChecked(x, y) {
				continue
			}
			area, fence := computeAreaScore2(inputG, x, y, inputG.plants[y][x])
			result += area * fence
		}
	}

	return result
}

func computeAreaScore2(g *garden, x, y int, plant string) (int, int) {
	if !g.isSamePlant(x, y, plant) || g.isChecked(x, y) {
		return 0, 0
	}

	left := g.isSamePlant(x-1, y, plant)
	leftUp := g.isSamePlant(x-1, y-1, plant)
	right := g.isSamePlant(x+1, y, plant)
	leftDown := g.isSamePlant(x-1, y+1, plant)
	rightUp := g.isSamePlant(x+1, y-1, plant)
	rightDown := g.isSamePlant(x+1, y+1, plant)
	up := g.isSamePlant(x, y-1, plant)
	down := g.isSamePlant(x, y+1, plant)

	pillars := 0
	if !left && !up {
		pillars += 1
	}
	if !up && !right {
		pillars += 1
	}
	if !right && !down {
		pillars += 1
	}
	if !down && !left {
		pillars += 1
	}

	if left && up && !leftUp {
		pillars += 1
	}
	if up && right && !rightUp {
		pillars += 1
	}
	if right && down && !rightDown {
		pillars += 1
	}
	if left && down && !leftDown {
		pillars += 1
	}
	g.setChecked(x, y)
	areaSize := 1
	a, p := computeAreaScore2(g, x+1, y, plant)
	areaSize += a
	pillars += p
	a, p = computeAreaScore2(g, x-1, y, plant)
	areaSize += a
	pillars += p
	a, p = computeAreaScore2(g, x, y+1, plant)
	areaSize += a
	pillars += p
	a, p = computeAreaScore2(g, x, y-1, plant)
	areaSize += a
	pillars += p

	return areaSize, pillars

}
