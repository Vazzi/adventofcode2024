package day09

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	expectedSolution := 1928
	inputData := utils.ReadFile(utils.InputDataPath(9, true))

	solution := firstSolution(inputData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}

func TestSecondSolution(t *testing.T) {
	expectedSolution := 2858
	inputData := utils.ReadFile(utils.InputDataPath(9, true))

	solution := secondSolution(inputData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
