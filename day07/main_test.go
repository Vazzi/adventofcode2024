package day07

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	expectedSolution := 3749
	inputData := utils.ReadLines(utils.InputDataPath(7, true))

	solution := firstSolution(inputData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
