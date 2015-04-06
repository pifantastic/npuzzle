package main

import "testing"

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	if board.State != 0 {
		t.Error("New Board should have state 0")
	}
}

func TestNewBoardFromMatrix(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	if board.State != EIGHT_PIECE_GOAL {
		t.Error("Expected %d, got ", EIGHT_PIECE_GOAL, board.State)
	}

	board2 := NewBoardFromMatrix([][]int{
		[]int{0, 1, 2, 3},
		[]int{4, 5, 6, 7},
		[]int{8, 9, 10, 11},
		[]int{12, 13, 14, 15},
	})

	if board2.State != FIFTEEN_PIECE_GOAL {
		t.Error("Expected ", FIFTEEN_PIECE_GOAL, "got ", board2.State)
	}
}

func TestClone(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	board2 := board.Clone()

	board.SetTileAt(0, 8)
	if board2.GetTileAt(0) != 0 {
		t.Error("Expected 0, got", board2.GetTileAt(0))
	}
}

func TestGetTileAt(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			index := i*3 + j
			if board.GetTileAt(index) != index {
				t.Error("Expected tile at", index, "to equal", index, "got", board.GetTileAt(index))
			}
		}
	}
}

func TestSetTileAt(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			index := i*3 + j
			board.SetTileAt(index, 8-index)
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			index := i*3 + j
			if board.GetTileAt(index) != 8-index {
				t.Error("Expected tile at", index, "to equal", 8-index, "got", board.GetTileAt(index))
			}
		}
	}
}

func TestCost(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	if board.Cost() != 0 {
		t.Error("Expected 0, got", board.Cost())
	}
}

func TestLeft(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{1, 0, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	newBoard := board.Left()
	if newBoard.GetTileAt(0) != BLANK {
		t.Error("Expected", BLANK, "got", newBoard.GetTileAt(0))
	}

	newBoard = newBoard.Left()
	if newBoard != nil {
		t.Error("Expected", nil, "got", newBoard)
	}
}

func TestRight(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{1, 0, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	newBoard := board.Right()
	if newBoard.GetTileAt(2) != BLANK {
		t.Error("Expected", BLANK, "got", newBoard.GetTileAt(2))
	}

	newBoard = newBoard.Right()
	if newBoard != nil {
		t.Error("Expected", nil, "got", newBoard)
	}
}

func TestUp(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{3, 1, 2},
		[]int{0, 4, 5},
		[]int{6, 7, 8},
	})

	newBoard := board.Up()
	if newBoard.GetTileAt(0) != BLANK {
		t.Error("Expected", BLANK, "got", newBoard.GetTileAt(0))
	}

	newBoard = newBoard.Up()
	if newBoard != nil {
		t.Error("Expected", nil, "got", newBoard)
	}
}

func TestDown(t *testing.T) {
	board := NewBoardFromMatrix([][]int{
		[]int{3, 1, 2},
		[]int{0, 4, 5},
		[]int{6, 7, 8},
	})

	newBoard := board.Down()
	if newBoard.GetTileAt(6) != BLANK {
		t.Error("Expected", BLANK, "got", newBoard.GetTileAt(6))
	}

	newBoard = newBoard.Down()
	if newBoard != nil {
		t.Error("Expected", nil, "got", newBoard)
	}
}
