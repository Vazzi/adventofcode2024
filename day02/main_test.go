package day02

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	data := utils.ReadInts(utils.InputDataPath(2, true), " ")

	result := computeNumberOfSafeReports(data, false)
	if result != 2 {
		t.Fatalf("Expected result to be 2, got %d", result)
	}
}

func TestSecondSolution(t *testing.T) {
	data := utils.ReadInts(utils.InputDataPath(2, true), " ")

	result := computeNumberOfSafeReports(data, true)
	if result != 4 {
		t.Fatalf("Expected result to be 4, got %d", result)
	}
}
