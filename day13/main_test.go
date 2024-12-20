package day13

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	expectedSolution := 480

	solution := firstSolution(utils.InputDataPath(13, true))
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}

func TestSecondSolution(t *testing.T) {
	expectedSolution := 875318608908

	solution := secondSolution(utils.InputDataPath(13, true))
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
