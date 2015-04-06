package main

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue, 0)

	board1 := NewBoardFromMatrix([][]int{
		[]int{1, 2, 0},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})

	board2 := NewBoardFromMatrix([][]int{
		[]int{1, 2, 5},
		[]int{3, 4, 8},
		[]int{6, 7, 0},
	})

	board3 := NewBoardFromMatrix([][]int{
		[]int{1, 2, 5},
		[]int{3, 4, 0},
		[]int{6, 7, 8},
	})

	heap.Push(&pq, board1)
	heap.Push(&pq, board2)
	heap.Push(&pq, board3)

	heap.Init(&pq)

	board4 := NewBoardFromMatrix([][]int{
		[]int{1, 0, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	})
	heap.Push(&pq, board4)

	expected := []*Board{board4, board1, board3, board2}

	n := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Board)

		if item.State != expected[n].State {
			t.Error("Unexpected item:", item.State)
		}
		n++
	}
}
