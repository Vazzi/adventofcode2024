package day03

import "testing"

var testInput = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestComputeAllMul(t *testing.T) {
	solution := firstSolution(testInput)
	if solution != 161 {
		t.Fatalf("Expected solution to be 161, got %d", solution)
	}
}

func TestSecondSolution(t *testing.T) {
	solution := secondSolution(testInput)
	if solution != 48 {
		t.Fatalf("Expected solution to be 48, got %d", solution)
	}
}
