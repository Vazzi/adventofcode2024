package day10

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	expectedSolution := 36
	inputData := utils.ReadInts(utils.InputDataPath(10, true), "")

	solution := firstSolution(inputData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}

func TestSecondSolution(t *testing.T) {
	expectedSolution := 81
	inputData := utils.ReadInts(utils.InputDataPath(10, true), "")

	solution := secondSolution(inputData)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
