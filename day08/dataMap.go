package day08

type dataMap struct {
	mapData [][]string
	width   int
	height  int
}

func newDataMap(data [][]string) *dataMap {
	return &dataMap{
		mapData: data,
		width:   len(data[0]),
		height:  len(data),
	}
}

func (data *dataMap) placeAntinode(node point) {
	data.mapData[node.y][node.x] = "#"
}

func (data dataMap) isInMap(point point) bool {
	return point.x >= 0 && point.y >= 0 && point.x < data.width && point.y < data.height
}

func (data dataMap) canCountAntinode(node point) bool {
	if data.isInMap(node) {
		if data.mapData[node.y][node.x] != "#" {
			return true
		}
	}
	return false
}
