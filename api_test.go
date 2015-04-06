package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPuzzles(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/v1/puzzles", nil)
	if err != nil {
		t.Error(err)
	}

	response := httptest.NewRecorder()

	Puzzles(response, request)

	var puzzle Puzzle
	err = json.Unmarshal(response.Body.Bytes(), &puzzle)
	if err != nil {
		t.Error(err)
	}

	if puzzle.Dimension != 3 {
		t.Error("Expected puzzle dimension to be 3, got", puzzle.Dimension)
	}

	if !puzzle.Board.Solvable() {
		t.Error("Expected puzzle to be solvable")
	}
}

func TestSolutions(t *testing.T) {
	body := strings.NewReader("{\"board\":[1,0,2,3,4,5,6,7,8]}")

	request, err := http.NewRequest("POST", "/api/v1/solutions", body)
	if err != nil {
		t.Error(err)
	}

	response := httptest.NewRecorder()

	Solutions(response, request)

	var solution []string
	err = json.Unmarshal(response.Body.Bytes(), &solution)
	if err != nil {
		t.Error(err)
	}

	if len(solution) != 1 {
		t.Error("Expected solution to have 1 move, got", len(solution))
	}

	if solution[0] != LEFT {
		t.Error("Expected solution to contain one LEFT move, saw", solution[0])
	}
}
