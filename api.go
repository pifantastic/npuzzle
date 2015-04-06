package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

type SolutionRequest struct {
	Board []int `json:"board"`
}

func Puzzles(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	size, err := strconv.ParseInt(request.URL.Query().Get("size"), 10, 0)

	if err != nil {
		size = 3
	}

	puzzle := NewPuzzle(int(size))
	puzzleJson, _ := json.Marshal(puzzle)
	response.Write(puzzleJson)
}

func Solutions(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(request.Body)

	var solutionRequest SolutionRequest
	err := decoder.Decode(&solutionRequest)

	if err != nil {
		http.Error(response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	board := NewBoardFromArray(solutionRequest.Board)
	puzzle := NewPuzzle(board.Dimension)
	puzzle.Board = board

	log.Info("Solving:", puzzle)

	solution := puzzle.Solve()
	solutionJson, err := json.Marshal(solution)

	if err != nil {
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response.Write(solutionJson)
}
