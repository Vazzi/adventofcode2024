package day06

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	inputData := utils.ReadChars(utils.InputDataPath(6, true))

	solution := firstSolution(inputData)
	if solution != 41 {
		t.Fatalf("Expected solution to be 41, got %d", solution)
	}
}

func TestSecondSolution(t *testing.T) {
	inputData := utils.ReadChars(utils.InputDataPath(6, true))

	solution := secondSolution(inputData)

	if solution != 6 {
		t.Fatalf("Expected solution to be 6, got %d", solution)
	}
}
