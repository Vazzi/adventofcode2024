package day07

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	expectedSolution := 3749
	inputData := utils.ReadLines(utils.InputDataPath(7, true))

	solution := solution(inputData, false)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}

func TestSecondSolution(t *testing.T) {
	expectedSolution := 11387
	inputData := utils.ReadLines(utils.InputDataPath(7, true))

	solution := solution(inputData, true)
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
