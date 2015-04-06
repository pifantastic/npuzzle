package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Board struct {
	Dimension int  `json:"-"`
	State     uint `json:"state"`

	previous *Board `json:"-"`
	move     string `json:"-"`
	index    int    `json:"-"`
}

func NewBoard() *Board {
	board := Board{}
	return &board
}

func NewBoardFromMatrix(matrix [][]int) *Board {
	board := Board{Dimension: len(matrix)}

	for i := 0; i < board.Dimension; i++ {
		for j := 0; j < board.Dimension; j++ {
			index := i*board.Dimension + j
			board.State |= (uint(matrix[i][j]) << (uint(index) << uint(2)))
		}
	}

	return &board
}

func NewBoardFromArray(array []int) *Board {
	board := Board{Dimension: int(math.Sqrt(float64(len(array))))}

	for i := 0; i < board.Len(); i++ {
		board.State |= (uint(array[i]) << (uint(i) << uint(2)))
	}

	return &board
}

func NewRandomBoard(dimension int) *Board {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := dimension * dimension

	// Generate an initial permutation.
	board := NewBoardFromArray(r.Perm(length))

	// Continually generate new permutations until the puzzle is solvable.
	for !board.Solvable() {
		board = NewBoardFromArray(r.Perm(length))
	}

	return board
}

func (board *Board) MarshalJSON() ([]byte, error) {
	state := make([]int, 0)

	for i := 0; i < board.Len(); i++ {
		state = append(state, board.GetTileAt(i))
	}

	return json.Marshal(struct {
		Board
		State []int `json:"state"`
	}{
		Board: Board(*board),
		State: state,
	})
}

func (board *Board) UnmarshalJSON(data []byte) error {
	json.Unmarshal(data, struct {
		Board
		State []int `json:"state"`
	}{})

	return nil
}

func (board *Board) Len() int {
	return board.Dimension * board.Dimension
}

func (board *Board) Solvable() bool {
	inversions := 0

	for i := 0; i < board.Len(); i++ {
		iTile := board.GetTileAt(i)
		if iTile != BLANK {
			for j := i + 1; j < board.Len(); j++ {
				jTile := board.GetTileAt(j)
				if jTile != BLANK && jTile < iTile {
					inversions++
				}
			}
		} else {
			if (board.Dimension & 0x1) == 0 {
				inversions += (1 + i/board.Dimension)
			}
		}
	}
	return (inversions & 0x1) != 1
}

func (board *Board) Solved() bool {
	if board.Dimension == 3 {
		return board.State == EIGHT_PIECE_GOAL
	} else if board.Dimension == 4 {
		return board.State == FIFTEEN_PIECE_GOAL
	}
	return false
}

func (board *Board) String() string {
	buffer := bytes.NewBufferString("\n")

	for i := 0; i < board.Dimension; i++ {
		for j := 0; j < board.Dimension; j++ {
			index := i*board.Dimension + j
			fmt.Fprintf(buffer, "%d ", board.GetTileAt(index))
		}
		fmt.Fprint(buffer, "\n")
	}

	return buffer.String()
}

func (board *Board) Clone() *Board {
	clone := NewBoard()
	*clone = *board
	clone.previous = board
	return clone
}

func (board *Board) Cost() int {
	cost := 0

	for i := 0; i < board.Len(); i++ {
		if board.GetTileAt(i) == BLANK {
			continue
		}

		row := i / board.Dimension
		col := i % board.Dimension

		goalRow := board.GetTileAt(i) / board.Dimension
		goalCol := board.GetTileAt(i) % board.Dimension

		distance := math.Abs(float64(row-goalRow)) + math.Abs(float64(col-goalCol))

		cost += int(distance)
	}

	return cost
}

func (board *Board) PositionOfBlank() int {
	for i := 0; i < board.Len(); i++ {
		if board.GetTileAt(i) == BLANK {
			return i
		}
	}
	return -1
}

func (board *Board) GetTileAt(index int) int {
	return int((board.State >> (uint(index) << uint(2))) & 0xF)
}

func (board *Board) SetTileAt(index int, value int) {
	mask := uint(0xF) << (uint(index) << uint(2))
	board.State &= ^mask
	board.State |= (uint(value) << (uint(index) << uint(2)))
}

func (board *Board) Move(direction string) *Board {
	switch direction {
	case UP:
		return board.Up()
	case DOWN:
		return board.Down()
	case LEFT:
		return board.Left()
	case RIGHT:
		return board.Right()
	}
	return nil
}

func (board *Board) Left() *Board {
	blankPosition := board.PositionOfBlank()
	if blankPosition%board.Dimension == 0 {
		return nil
	}

	clone := board.Clone()
	clone.move = LEFT
	tile := clone.GetTileAt(blankPosition - 1)
	clone.SetTileAt(blankPosition-1, BLANK)
	clone.SetTileAt(blankPosition, tile)

	return clone
}

func (board *Board) Right() *Board {
	blankPosition := board.PositionOfBlank()
	if (blankPosition+1)%board.Dimension == 0 {
		return nil
	}

	clone := board.Clone()
	clone.move = RIGHT
	tile := clone.GetTileAt(blankPosition + 1)
	clone.SetTileAt(blankPosition+1, BLANK)
	clone.SetTileAt(blankPosition, tile)

	return clone
}

func (board *Board) Up() *Board {
	blankPosition := board.PositionOfBlank()
	if blankPosition < board.Dimension {
		return nil
	}

	clone := board.Clone()
	clone.move = UP
	tile := clone.GetTileAt(blankPosition - board.Dimension)
	clone.SetTileAt(blankPosition-board.Dimension, BLANK)
	clone.SetTileAt(blankPosition, tile)

	return clone
}

func (board *Board) Down() *Board {
	blankPosition := board.PositionOfBlank()
	if blankPosition >= (board.Len() - board.Dimension) {
		return nil
	}

	clone := board.Clone()
	clone.move = DOWN
	tile := clone.GetTileAt(blankPosition + board.Dimension)
	clone.SetTileAt(blankPosition+board.Dimension, BLANK)
	clone.SetTileAt(blankPosition, tile)

	return clone
}
