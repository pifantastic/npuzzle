package main

import (
	"bytes"
	"container/heap"
	"fmt"
)

const (
	UP    = "Up"
	DOWN  = "Down"
	LEFT  = "Left"
	RIGHT = "Right"

	BLANK = 8

	EIGHT_PIECE_GOAL   uint = 36344967696
	FIFTEEN_PIECE_GOAL uint = 18364758544493064720
)

var MOVES = []string{UP, DOWN, LEFT, RIGHT}

type Puzzle struct {
	Dimension int    `json:"dimension"`
	Board     *Board `json:"board"`
}

func NewPuzzle(dimension int) *Puzzle {
	puzzle := Puzzle{
		Dimension: dimension,
	}
	puzzle.Board = NewRandomBoard(dimension)
	return &puzzle
}

func (puzzle *Puzzle) Solve() []string {
	solved := false

	openMap := make(map[uint]*Board)
	closedMap := make(map[uint]*Board)
	open := make(PriorityQueue, 0)

	heap.Init(&open)
	heap.Push(&open, puzzle.Board)

	var current *Board

	for open.Len() > 0 {
		current = heap.Pop(&open).(*Board)

		if current.Solved() {
			solved = true
			break
		}

		closedMap[current.State] = current

		for _, move := range MOVES {
			next := current.Move(move)

			if next == nil {
				continue
			}

			closedItem, closedItemExists := closedMap[next.State]
			openItem, openItemExists := openMap[next.State]

			if !closedItemExists || next.Cost() < closedItem.Cost() {
				heap.Push(&open, next)
				openMap[next.State] = next
			} else if openItemExists && next.Cost() < openItem.Cost() {
				open.update(openItem)
				openMap[next.State] = next
			}
		}
	}

	if solved {
		solution := []string{current.move}

		for current.previous != nil && current.previous.move != "" {
			solution = append([]string{current.previous.move}, solution...)
			current = current.previous
		}

		return solution
	}

	return nil
}

func (puzzle *Puzzle) String() string {
	buffer := bytes.NewBufferString("\n")

	fmt.Fprintf(buffer, "Dimension: %d\n", puzzle.Dimension)
	fmt.Fprintf(buffer, "Board:%s", puzzle.Board)

	return buffer.String()
}
