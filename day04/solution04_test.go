package day04

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	testInput := utils.ReadChars(utils.InputDataPath(4, true))
	solution := firstSolution(testInput)
	if solution != 18 {
		t.Fatalf("Expected solution to be 18, got %d", solution)
	}
}

func TestSecondSolution(t *testing.T) {
	testInput := utils.ReadChars(utils.InputDataPath(4, true))
	solution := secondSolution(testInput)
	if solution != 9 {
		t.Fatalf("Expected solution to be 9, got %d", solution)
	}
}
