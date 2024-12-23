package day17

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution01(t *testing.T) {
	expectedSolution := "4,6,3,5,6,3,5,2,1,0"

	solution := firstSolution(utils.InputDataPath(17, true))
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %s, got %s", expectedSolution, solution)
	}
}
