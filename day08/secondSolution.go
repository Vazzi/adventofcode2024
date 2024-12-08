package day08

func secondSolution(input [][]string) int {
	data := newDataMap(input)
	antennas, score := replaceAndCountAntennas(data)

	for antennaType := range antennas {
		score += secondCountAllAntiNodes(antennas[antennaType], data)
	}
	return score
}

func secondCountAllAntiNodes(antennas []point, mapData *dataMap) int {
	antinodesCount := 0

	for oIndex, origin := range antennas {
		for aIndex, antenna := range antennas {
			if oIndex == aIndex {
				continue
			}
			v := antenna.makeVectorWithOrigin(origin)
			origv := newVector(v.x, v.y)
			for {
				v.add(origv)
				antinode := v.makePointFromVectorFromOrigin(origin)
				if !mapData.isInMap(antinode) {
					break
				}
				if mapData.canCountAntinode(antinode) {
					antinodesCount++
					mapData.placeAntinode(antinode)
				}
			}
		}
	}
	return antinodesCount
}

func replaceAndCountAntennas(data *dataMap) (antennasMap, int) {
	count := 0
	antennas := make(antennasMap)
	for y := range data.mapData {
		for x, char := range data.mapData[y] {
			if char != "." {
				count++
				antennas[char] = append(antennas[char], point{x, y})
				data.mapData[y][x] = "#"
			}
		}
	}
	return antennas, count
}
