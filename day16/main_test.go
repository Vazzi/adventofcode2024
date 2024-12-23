package day16

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution01(t *testing.T) {
	expectedSolution := 11048

	mapData := utils.ReadChars(utils.InputDataPath(16, true))

	solution := firstSolution(mapData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}