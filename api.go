package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	PAGE_SIZE = 50
)

type SolutionRequest struct {
	Board []int `json:"board"`
}

// Returns a new, randomized puzzle.
func Puzzles(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// The "dimension" is the length of the side of the desired puzzle.
	dimension, err := strconv.ParseInt(request.URL.Query().Get("dimension"), 10, 0)

	if err != nil {
		dimension = 3
	}

	puzzle := NewPuzzle(int(dimension))
	puzzleJson, _ := json.Marshal(puzzle)
	response.Write(puzzleJson)
}

// Solve a given puzzle.
func Solutions(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Request body should be a JSON serialized `SolutionRequest`.
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

// Get the specified page of the leaderboards.
func Leaderboards(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	page, err := strconv.ParseInt(request.URL.Query().Get("page"), 10, 0)

	if err != nil {
		page = 1
	}

	entries := make([]LeaderboardEntry, 0)

	for x := 0; x < PAGE_SIZE; x++ {
		entries = append(entries, LeaderboardEntry{
			Id:    int(page-1)*PAGE_SIZE + x,
			Name:  randomdata.FullName(randomdata.RandomGender),
			Moves: int(page-1)*PAGE_SIZE + x,
			Hash:  fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(x)))),
		})
	}

	leaderboard := Leaderboard{
		Entries: entries,
	}

	leaderboardJson, _ := json.Marshal(leaderboard)
	response.Write(leaderboardJson)
}
