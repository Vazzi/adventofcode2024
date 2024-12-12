package day11

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	expectedSolution := 55312
	inputData := utils.ReadInts(utils.InputDataPath(11, true), " ")

	solution := firstSolution(inputData[0])
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
