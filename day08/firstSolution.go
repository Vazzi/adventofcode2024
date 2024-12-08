package day08

type point struct {
	x, y int
}

type antennasMap map[string][]point

func firstSolution(input [][]string) int {
	data := newDataMap(input)
	antennas := findAntennas(data)
	score := 0

	for antennaType := range antennas {
		score += countAllAntiNodes(antennas[antennaType], data)
	}
	return score
}

func countAllAntiNodes(antennas []point, mapData *dataMap) int {
	antinodesCount := 0

	for oIndex, origin := range antennas {
		for aIndex, antenna := range antennas {
			if oIndex == aIndex {
				continue
			}
			antinode := computeAntinode(origin, antenna)
			if mapData.canCountAntinode(antinode) {
				antinodesCount++
				mapData.placeAntinode(antinode)
			}
		}
	}
	return antinodesCount
}

func computeAntinode(origin, antenna point) point {
	vector := antenna.makeVectorWithOrigin(origin)
	vector.mullBy(2)
	antinode := vector.makePointFromVectorFromOrigin(origin)
	return antinode
}

func findAntennas(data *dataMap) antennasMap {
	antennas := make(antennasMap)
	for y := range data.mapData {
		for x, char := range data.mapData[y] {
			if char != "." {
				antennas[char] = append(antennas[char], point{x, y})
			}
		}
	}
	return antennas
}

func (point point) makeVectorWithOrigin(origin point) vector {
	return newVector(point.x-origin.x, point.y-origin.y)
}

type vector struct {
	point
}

func newVector(x, y int) vector {
	object := vector{
		point: point{x, y},
	}
	return object
}

func (v *vector) mullBy(num int) {
	v.x = v.x * num
	v.y = v.y * num
}

func (v vector) makePointFromVectorFromOrigin(origin point) point {
	return point{origin.x + v.x, origin.y + v.y}
}
