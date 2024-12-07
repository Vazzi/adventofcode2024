package day05

import (
	"adventofcode2024/utils"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	rules := readRules("../inputData/test/day05-rules.txt")
	data := utils.ReadWords("../inputData/test/day05-data.txt", ",")

	solution := firstSolution(rules, data)
	if solution != 143 {
		t.Fatalf("Expected solution to be 143, got %d", solution)
	}
}

func TestSecondSolution(t *testing.T) {
	rules := readRules("../inputData/test/day05-rules.txt")
	data := utils.ReadWords("../inputData/test/day05-data.txt", ",")

	solution := secondSolution(rules, data)

	if solution != 123 {
		t.Fatalf("Expected solution to be 123, got %d", solution)
	}
}
