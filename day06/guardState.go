package day06

const (
	Up    int = 0
	Right     = 1
	Down      = 2
	Left      = 3
)

type guardState struct {
	x, y int
	dir  int
}

func (gs *guardState) turnRight() {
	gs.dir = (gs.dir + 1) % 4
}

func (gs guardState) nextStep() (int, int) {
	switch gs.dir {
	case Up:
		return gs.x, gs.y - 1
	case Right:
		return gs.x + 1, gs.y
	case Down:
		return gs.x, gs.y + 1
	case Left:
		return gs.x - 1, gs.y
	}
	return gs.x, gs.y
}

func findGuard(mapData [][]string) guardState {
	for y := range mapData {
		for x := range mapData[y] {
			if mapData[y][x] == "^" {
				return guardState{x, y, Up}
			}
		}
	}
	panic("Guard not found")
}
