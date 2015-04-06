package main

import "testing"

func TestNewPuzzle(t *testing.T) {
	puzzle := NewPuzzle(3)

	if puzzle.Dimension != 3 {
		t.Error("New Puzzle should have dimension 3")
	}
}
