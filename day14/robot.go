package day14

import (
	"regexp"
	"strconv"
)

type position struct {
	x, y int
}

type velocity struct {
	x, y int
}

type robot struct {
	p position
	v velocity
}

func (r *robot) move(spaceWidth, spaceHeight int) {
	r.p.x = teleportIfNeeded(r.p.x+r.v.x, spaceWidth)
	r.p.y = teleportIfNeeded(r.p.y+r.v.y, spaceHeight)
}

func teleportIfNeeded(pos, size int) int {
	if pos < 0 {
		return size + pos
	}
	return pos % size
}

func newRobot(input string) *robot {
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	matches := re.FindStringSubmatch(input)
	pX, _ := strconv.Atoi(matches[1])
	pY, _ := strconv.Atoi(matches[2])
	vX, _ := strconv.Atoi(matches[3])
	vY, _ := strconv.Atoi(matches[4])
	return &robot{position{pX, pY}, velocity{vX, vY}}
}

func (r *robot) inQuadrant(spaceWidth, spaceHeight int) (int, bool) {
	quadrantId := -1

	if r.p.x == spaceWidth/2 || r.p.y == spaceHeight/2 {
		return quadrantId, false
	}

	if r.p.x < spaceWidth/2 {
		quadrantId = 0
	} else if r.p.x > spaceWidth/2 {
		quadrantId = 1
	}

	if r.p.y > spaceHeight/2 {
		quadrantId += 2
	}

	return quadrantId, quadrantId != -1
}
