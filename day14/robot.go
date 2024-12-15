package day14

import (
	"regexp"
	"strconv"
)

type robot struct {
	px, py int
	vx, vy int
}

func (r *robot) move(spaceWidth, spaceHeight int) {
	r.px = teleportIfNeeded(r.px+r.vx, spaceWidth)
	r.py = teleportIfNeeded(r.py+r.vy, spaceHeight)
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
	return &robot{pX, pY, vX, vY}
}

func (r *robot) inQuadrant(spaceWidth, spaceHeight int) (int, bool) {
	quadrantId := -1

	if r.px == spaceWidth/2 || r.py == spaceHeight/2 {
		return quadrantId, false
	}

	if r.px < spaceWidth/2 {
		quadrantId = 0
	} else if r.px > spaceWidth/2 {
		quadrantId = 1
	}

	if r.py > spaceHeight/2 {
		quadrantId += 2
	}

	return quadrantId, quadrantId != -1
}
