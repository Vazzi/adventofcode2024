package day15

import (
	"testing"
)

func TestFirstSolution01(t *testing.T) {
	expectedSolution := 2028

	solution := firstSolution("../inputData/test/day15-map-01.txt", "../inputData/test/day15-moves-01.txt")
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}

func TestFirstSolution02(t *testing.T) {
	expectedSolution := 10092

	solution := firstSolution("../inputData/test/day15-map-02.txt", "../inputData/test/day15-moves-02.txt")
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}

func TestSecondSolution02(t *testing.T) {
	expectedSolution := 9021

	solution := secondSolution("../inputData/test/day15-map-02.txt", "../inputData/test/day15-moves-02.txt")
	if solution != expectedSolution {
		t.Fatalf("Expected solution to be %d, got %d", expectedSolution, solution)
	}
}
