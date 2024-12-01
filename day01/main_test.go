package day01

import (
	"testing"
)

func TestComputeDistance(t *testing.T) {
	firstList := []int{3, 4, 2, 1, 3, 3}
	secondList := []int{4, 3, 5, 3, 9, 3}

	distance := computeDistance(firstList, secondList)
	if distance != 11 {
		t.Fatalf("Expected distance to be 11, got %d", distance)
	}
}

func TestSimilarityScore(t *testing.T) {
	firstList := []int{3, 4, 2, 1, 3, 3}
	secondList := []int{4, 3, 5, 3, 9, 3}

	score := similarityScore(firstList, secondList)
	if score != 31 {
		t.Fatalf("Expected similarity score to be 31, got %d", score)
	}
}
