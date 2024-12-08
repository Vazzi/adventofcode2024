package day08

type antennasMap map[string][]point

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

type point struct {
	x, y int
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

func (point point) makeVectorWithOrigin(origin point) vector {
	return newVector(point.x-origin.x, point.y-origin.y)
}

func (v *vector) mullBy(num int) {
	v.x = v.x * num
	v.y = v.y * num
}

func (v *vector) add(v2 vector) {
	v.x = v.x + v2.x
	v.y = v.y + v2.y
}

func (v vector) makePointFromVectorFromOrigin(origin point) point {
	return point{origin.x + v.x, origin.y + v.y}
}
