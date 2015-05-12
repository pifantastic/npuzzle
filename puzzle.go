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

	BLANK = 0

	EIGHT_PIECE_GOAL   uint = 36344967696
	FIFTEEN_PIECE_GOAL uint = 18364758544493064720
)

var MOVES = []string{UP, DOWN, LEFT, RIGHT}

// Represents an n-piece puzzle with n=Dimension.
type Puzzle struct {
	Dimension int    `json:"dimension"`
	Board     *Board `json:"board"`
}

// Create a new puzzle with a randomized board.
func NewPuzzle(dimension int) *Puzzle {
	puzzle := Puzzle{
		Dimension: dimension,
	}
	puzzle.Board = NewRandomBoard(dimension)
	return &puzzle
}

// Returns an array of the "moves" required to solve the puzzle. Implements the
// A* search algorithm (https://en.wikipedia.org/wiki/A*_search_algorithm).
func (puzzle *Puzzle) Solve() []string {
	solved := false

	// Board states already evaluated.
	closedMap := make(map[uint]*Board)

	// Current set of board states being evaluated. The map is for fast lookups
	// and the priority queue sorts board states by cost.
	openMap := make(map[uint]*Board)
	open := make(PriorityQueue, 0)

	heap.Init(&open)

	// Start with the puzzle's current solution on the open queue.
	heap.Push(&open, puzzle.Board)

	var current *Board

	for open.Len() > 0 {
		// Grab the best move from the open queue.
		current = heap.Pop(&open).(*Board)

		// Check if we've reached the solution.
		if current.Solved() {
			solved = true
			break
		}

		// Record that we've seen this board state.
		closedMap[current.State] = current

		// Expand the board's state into all possible moves.
		for _, move := range MOVES {
			next := current.Move(move)

			if next == nil {
				continue
			}

			// Check the new state against the closed and open maps.
			closedItem, closedItemExists := closedMap[next.State]
			openItem, openItemExists := openMap[next.State]

			// If this board state has not been visited before, or its cost has
			// decreased, place it on the open map.
			if !closedItemExists || next.cost < closedItem.cost {
				heap.Push(&open, next)
				openMap[next.State] = next
			} else if openItemExists && next.cost < openItem.cost {
				// If this board state is currently in the open map, and its cost has
				// improved, update its cost in the open queue.
				openItem.cost = next.cost
				open.update(openItem)
				openMap[next.State] = next
			}
		}
	}

	if solved {
		solution := []string{current.move}

		// Reconstruct the solution by tracing the final board state back through
		// its preceding states.
		for current.previous != nil && current.previous.move != "" {
			solution = append([]string{current.previous.move}, solution...)
			current = current.previous
		}

		return solution
	}

	return nil
}

// Returns a string representation of the puzzle.
func (puzzle *Puzzle) String() string {
	buffer := bytes.NewBufferString("\n")

	fmt.Fprintf(buffer, "Dimension: %d\n", puzzle.Dimension)
	fmt.Fprintf(buffer, "Board:%s", puzzle.Board)

	return buffer.String()
}
