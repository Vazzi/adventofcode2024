package day12

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	expectedSolution := 1930
	inputData := utils.ReadChars(utils.InputDataPath(12, true))

	solution := firstSolution(inputData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}

func TestSecondSolution(t *testing.T) {
	expectedSolution := 1206
	inputData := utils.ReadChars(utils.InputDataPath(12, true))

	solution := secondSolution(inputData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
